package balancer

import (
	"github.com/MarioLegenda/execman/internal/runners"
	"math"
	"sync"
)

type Result struct {
	Result  string
	Success bool
	Error   error
}

type Job struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	ContainerName string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string

	PackageName string

	ResultCh chan Result
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
	sync.Mutex

	done chan interface{}
}

/*
*
In general, the balancer should balance trough workers jobs to containers. For example:

There are 100 workers and 10 containers, a job worker will be picked with the least number of jobs on
it and the container with the least number of jobs on it. Benchmarking should be done but every container
should have at least 20 workers before it.
*/
func New(initialWorkers int, containers []string, containerTag string, done chan interface{}) *Balancer {
	balancer := &Balancer{
		containers:        make(map[string]int),
		workerControllers: make(map[int]int),
		workers:           make([]chan Job, initialWorkers),
		done:              done,
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
	b.Lock()
	workerIdx := pickWorker(b)
	job.ContainerName = pickContainer(b)
	b.Unlock()

	b.workers[workerIdx] <- job
}

func (b *Balancer) StartWorkers() {
	for workerIdx, worker := range b.workers {
		go func(workerIdx int, worker chan Job) {
			for {
				select {
				case <-b.done:
					close(worker)
					return
				case job := <-worker:
					// result of the job run
					res := runners.Run(runners.Params{
						ExecutionDir: job.ExecutionDir,

						BuilderType:       job.BuilderType,
						ExecutionType:     job.ExecutionType,
						ContainerName:     job.ContainerName,
						EmulatorName:      job.EmulatorName,
						EmulatorExtension: job.EmulatorExtension,
						EmulatorText:      job.EmulatorText,

						PackageName: job.PackageName,
					})

					b.Lock()
					b.workerControllers[workerIdx]--
					b.containers[job.ContainerName]--
					b.Unlock()

					job.ResultCh <- Result{
						Result:  res.Result,
						Success: res.Success,
						Error:   res.Error,
					}
				}
			}
		}(workerIdx, worker)
	}
}

// pick a worker with the least amount of jobs on it
// this function is executed within a lock in the AddJob() function
func pickWorker(b *Balancer) int {
	leastJobs := math.MaxInt
	leastBusyWorkerIdx := -1
	workers := b.workerControllers

	for workerIdx, numOfJobs := range workers {
		if numOfJobs < leastJobs {
			leastJobs = numOfJobs
			leastBusyWorkerIdx = workerIdx
		}
	}

	b.workerControllers[leastBusyWorkerIdx]++
	return leastBusyWorkerIdx
}

// pick a container with the least amount of jobs on it
// this function is executed within a lock in the AddJob() function
func pickContainer(b *Balancer) string {
	leastBusy := math.MaxInt
	var selected string

	for name, count := range b.containers {
		if count < leastBusy {
			leastBusy = count
			selected = name
		}
	}

	b.containers[selected]++
	return selected
}
