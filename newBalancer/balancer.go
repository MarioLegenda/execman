package newBalancer

import (
	"fmt"
	"github.com/MarioLegenda/execman/balancer/runners"
	"math"
	"sync"
)

type Job struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	ContainerName string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string

	PackageName string
}

type Balancer struct {
	// containers is a map of containers this balancer
	// balances. Balancer will pick a container with the least
	// amount of jobs in it and add it to the runner (a runner runs the
	// actual docker exec on the already running container).
	containers map[string]int

	// Same as containers, workers also must be picked with the least amount
	// of jobs in them. From this map, a pick algorithm will pick the worker
	// with the least amount of jobs on it.
	workerControllers map[int]int

	workers []chan Job

	// this balancer is also a lock since it needs to lock
	// the access to containers and workerControllers members
	lock sync.Mutex

	done chan struct{}
}

/*
*
In general, the balancer should balance trough workers jobs to containers. For example:

There are 100 workers and 10 containers, a job worker will be picked with the least number of jobs on
it and the container with the least number of jobs on it. Benchmarking should be done but every container
should have at least 20 workers before it.
*/
func New(initialWorkers int, containers []string) *Balancer {
	balancer := &Balancer{
		containers:        make(map[string]int),
		workerControllers: make(map[int]int),
		workers:           make([]chan Job, initialWorkers),
		done:              make(chan struct{}),
	}

	for i := 0; i < initialWorkers; i++ {
		balancer.workers[i] = make(chan Job)
		balancer.workerControllers[i] = 0
	}

	for _, c := range containers {
		balancer.containers[c] = 0
	}

	return balancer
}

func (b *Balancer) AddJob(job Job) {
	workerIdx := pickWorker(b)

	fmt.Println("worker controllers: ", b.workerControllers)

	b.workers[workerIdx] <- job

	b.lock.Lock()
	b.workerControllers[workerIdx]++
	b.lock.Unlock()
}

func (b *Balancer) StartWorkers() {
	for workerIdx, worker := range b.workers {
		go func(workerIdx int, worker chan Job) {
			for {
				containerName := pickContainer(b)

				fmt.Println("container name: ", containerName)

				select {
				case <-b.done:
					return
				case job := <-worker:
					// result of the job run
					_ = runners.Run(runners.Params{
						ExecutionDir: job.ExecutionDir,

						BuilderType:       job.BuilderType,
						ExecutionType:     job.ExecutionType,
						ContainerName:     containerName,
						EmulatorName:      job.EmulatorName,
						EmulatorExtension: job.EmulatorExtension,
						EmulatorText:      job.EmulatorText,

						PackageName: job.PackageName,
					})

					b.lock.Lock()
					b.workerControllers[workerIdx]--
					b.containers[containerName]--
					b.lock.Unlock()
				}
			}
		}(workerIdx, worker)
	}
}

func (b *Balancer) Close() {
	close(b.done)
}

func pickWorker(b *Balancer) int {
	b.lock.Lock()
	defer b.lock.Unlock()

	leastJobs := math.MaxInt
	leastBusyWorkerIdx := -1
	workers := b.workerControllers

	for workerIdx, numOfJobs := range workers {
		if numOfJobs < leastJobs {
			leastJobs = numOfJobs
			leastBusyWorkerIdx = workerIdx
		}
	}

	return leastBusyWorkerIdx
}

func pickContainer(b *Balancer) string {
	b.lock.Lock()
	defer b.lock.Unlock()

	leastBusyContainer := math.MaxInt
	containers := b.containers
	containerName := ""

	for name, numOfJobs := range containers {
		if numOfJobs < leastBusyContainer {
			containerName = name
		}
	}

	return containerName
}
