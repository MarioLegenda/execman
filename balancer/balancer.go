package balancer

import (
	"github.com/MarioLegenda/execman/balancer/runners"
	"sync"
)

type Balancer interface {
	StartWorkers()
	AddJob(Job)
	Close()
}

type Job struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string

	PackageName string

	Output chan runners.Result
}

type worker struct {
	input chan Job
	name  string
	index int
}

type balancer struct {
	workers    []worker
	name       string
	controller []int32
	lock       sync.Mutex
	closing    bool
}

func NewBalancer(name string, initialWorkers int) Balancer {
	b := &balancer{
		workers:    make([]worker, 0),
		controller: make([]int32, 0),
		name:       name,
	}

	for i := 0; i < initialWorkers; i++ {
		b.workers = append(b.workers, worker{
			input: make(chan Job, 1),
			name:  name,
			index: i,
		})
		b.controller = append(b.controller, 0)
	}

	return b
}

func (b *balancer) StartWorkers() {
	wg := sync.WaitGroup{}
	for _, w := range b.workers {
		wg.Add(1)
		go func(worker worker, wg *sync.WaitGroup) {
			wg.Done()

			for job := range worker.input {
				if b.closing {
					job.Output <- runners.Result{
						Result:  "",
						Success: false,
						Error:   runners.WorkerShutdownError,
					}

					continue
				}

				res := runners.Run(runners.Params{
					ExecutionDir: job.ExecutionDir,

					BuilderType:       job.BuilderType,
					ExecutionType:     job.ExecutionType,
					ContainerName:     worker.name,
					EmulatorName:      job.EmulatorName,
					EmulatorExtension: job.EmulatorExtension,
					EmulatorText:      job.EmulatorText,

					PackageName: job.PackageName,
				})

				if res.Error != nil {
					b.lock.Lock()
					b.controller[worker.index] = b.controller[worker.index] - 1
					b.lock.Unlock()

					job.Output <- runners.Result{
						Result:  "",
						Success: false,
						Error:   res.Error,
					}

					continue
				}

				b.lock.Lock()
				b.controller[worker.index] = b.controller[worker.index] - 1
				b.lock.Unlock()

				job.Output <- res
			}
		}(w, &wg)
	}

	wg.Wait()
}

func (b *balancer) AddJob(j Job) {
	b.lock.Lock()

	idx := 0
	first := b.controller[0]
	for i, r := range b.controller {
		if r < first {
			idx = i
		}
	}

	b.controller[idx] = b.controller[idx] + 1

	b.lock.Unlock()

	b.workers[idx].input <- j
}

func (b *balancer) Close() {
	b.lock.Lock()
	b.closing = true
	b.lock.Unlock()

	l := len(b.controller)
	for {
		a := 0
		for _, r := range b.controller {
			if r == 0 {
				a++
			}

			if a == l {
				return
			}
		}
	}
}
