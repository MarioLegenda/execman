package execman

import (
	"fmt"
	"github.com/MarioLegenda/execman/balancer"
	"github.com/MarioLegenda/execman/balancer/runners"
	"github.com/MarioLegenda/execman/containerFactory"
	"github.com/MarioLegenda/execman/sdk"
	"sync"
)

type Job struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	EmulatorName      string
	EmulatorExtension string
	EmulatorTag       string
	EmulatorText      string
	PackageName       string
}

type Execution interface {
	Close()
	Closed() bool
	RunJob(j Job) runners.Result
}

type execution struct {
	controller map[string][]int32
	balancers  map[string][]balancer.Balancer
	lock       sync.Mutex
	close      bool
	name       string
}

type ContainerBlueprint struct {
	WorkerNum    int
	ContainerNum int
	Tag          string
}

func Init(executionDir string, blueprints []ContainerBlueprint) (*execution, error) {
	blueprints = sdk.Filter(blueprints, func(idx int, value ContainerBlueprint) bool {
		return value.ContainerNum != 0
	})

	s := &execution{
		balancers:  make(map[string][]balancer.Balancer),
		controller: make(map[string][]int32),
	}

	err := s.init(executionDir, blueprints)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (e *execution) Closed() bool {
	return e.close
}

func (e *execution) RunJob(j Job) runners.Result {
	e.lock.Lock()

	balancers := e.balancers[j.EmulatorTag]
	controller := e.controller[j.EmulatorTag]

	if e.close {
		e.lock.Unlock()

		return runners.Result{
			Result:  "",
			Success: false,
			Error:   fmt.Errorf("%w", TimeoutError),
		}
	}

	idx := 0
	first := controller[0]
	for i, r := range controller {
		if r < first {
			idx = i
		}
	}

	e.controller[j.EmulatorTag][idx] = e.controller[j.EmulatorTag][idx] + 1

	b := balancers[idx]

	e.lock.Unlock()

	output := make(chan runners.Result)

	b.AddJob(balancer.Job{
		ExecutionDir:      j.ExecutionDir,
		BuilderType:       j.BuilderType,
		ExecutionType:     j.ExecutionType,
		EmulatorName:      j.EmulatorName,
		EmulatorExtension: j.EmulatorExtension,
		EmulatorText:      j.EmulatorText,

		PackageName: j.PackageName,

		Output: output,
	})

	out := <-output
	close(output)

	e.lock.Lock()
	e.controller[j.EmulatorTag][idx] = e.controller[j.EmulatorTag][idx] - 1
	e.lock.Unlock()

	return out
}

func (e *execution) Close() {
	e.lock.Lock()
	e.close = true
	e.lock.Unlock()

	for _, balancers := range e.balancers {
		for _, b := range balancers {
			b.Close()
		}
	}

	containerFactory.Close()
}

func (e *execution) init(executionDir string, blueprints []ContainerBlueprint) error {
	workers := make(map[string]int)
	for _, blueprint := range blueprints {
		fmt.Println("Creating container: ", blueprint.Tag)

		errs := containerFactory.CreateContainers(executionDir, blueprint.Tag, blueprint.ContainerNum)

		if len(errs) != 0 {
			e.Close()

			log := ""
			for _, err := range errs {
				log += fmt.Sprintf("%s,", err.Error())
			}

			return fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Cannot boot container for tag %s: %s", blueprint.Tag, log))
		}

		workers[blueprint.Tag] = blueprint.WorkerNum

		containers := containerFactory.Containers(blueprint.Tag)

		for _, c := range containers {
			e.createBalancer(c.Name, c.Tag, blueprint.WorkerNum)
		}
	}

	return nil
}

func (e *execution) createBalancer(containerName, tag string, workerNum int) {
	b := balancer.NewBalancer(containerName, workerNum)
	b.StartWorkers()

	_, ok := e.balancers[tag]

	if ok {
		e.balancers[tag] = append(e.balancers[tag], b)
	} else {
		e.balancers[tag] = make([]balancer.Balancer, 0)
		e.balancers[tag] = append(e.balancers[tag], b)
	}

	_, ok = e.controller[tag]

	if ok {
		e.controller[tag] = append(e.controller[tag], 0)
	} else {
		e.controller[tag] = make([]int32, 0)
		e.controller[tag] = append(e.controller[tag], 0)
	}
}
