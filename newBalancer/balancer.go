package newBalancer

import (
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
}

/**

 */
func (b *Balancer) New(initialWorkers int, containers []string) *Balancer {
	balancer := &Balancer{
		containers:        make(map[string]int),
		workerControllers: make(map[int]int),
		workers:           make([]chan Job, initialWorkers),
	}

	for i := 0; i < initialWorkers; i++ {
		balancer.workers[i] = make(chan Job)
	}

	for _, c := range containers {
		balancer.containers[c] = 0
	}

	return balancer
}

func (b *Balancer) StartWorkers() {
	for _, worker := range b.workers {
		// future job
		_ = <-worker
		containerName := pickContainer(b)

		select {
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
		}
	}
}

func pickWorker(b *Balancer) int {
	b.lock.Lock()
	defer b.lock.Unlock()

	leastBusyWorker := math.MaxInt
	workers := b.workerControllers

	for workerIdx, numOfJobs := range workers {
		if numOfJobs < leastBusyWorker {
			leastBusyWorker = workerIdx
		}
	}

	return leastBusyWorker
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