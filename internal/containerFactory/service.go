// if this package has an error, it HAS to panic because it must succeed

// TODO: make this package into a single struct
package containerFactory

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"sync"
)

type container struct {
	pid int
	dir string

	Tag  string
	Name string
}

type RestartedContainer struct {
	Name             string
	OldContainerName string
}

type ContainerFactory struct {
	// holds all locks of this package
	sync.Mutex
	// holds the execution directory in which container
	// volumes reside
	executionDirectory string
	// holds containers under a specific tag name of the language
	// that the user has selected to run
	containers map[string][]container
}

func New(executionDirectory string) *ContainerFactory {
	return &ContainerFactory{
		Mutex:              sync.Mutex{},
		executionDirectory: executionDirectory,
		containers:         make(map[string][]container),
	}
}

func (c *ContainerFactory) CreateContainers(tag string, containerNum int) {
	wg := sync.WaitGroup{}
	wg.Add(containerNum)
	for i := 0; i < containerNum; i++ {
		go func() {
			defer wg.Done()

			newContainer := c.createContainer(tag)

			if !isContainerRunning(newContainer.Name) {
				panic(fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))
			}
		}()
	}

	wg.Wait()
}

func (c *ContainerFactory) Close() {
	c.Lock()
	contArr := containersToSlice(c.containers)
	c.Unlock()

	// stop all docker containers and clean up volume directories
	wg := sync.WaitGroup{}
	wg.Add(len(contArr))
	for _, entry := range contArr {
		go func(c container, wg *sync.WaitGroup) {
			defer wg.Done()
			cleanupContainer(c.Name, c.pid, c.dir)
		}(entry, &wg)
	}

	wg.Wait()

	// remove the entire execution directory since it is recreated on every start of execman
	err := os.RemoveAll(c.executionDirectory)
	if err != nil {
		panic(fmt.Sprintf("Cannot remove execution directory %s. You will have to remove in manutally.", c.executionDirectory))
	}

	c.containers = make(map[string][]container)
}

func (c *ContainerFactory) Containers(tagName string) []container {
	return c.containers[tagName]
}

// This code has to work and must panic if it does not
func (c *ContainerFactory) Watch(tagName string, done chan interface{}) chan RestartedContainer {
	watchCh := make(chan RestartedContainer)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				conts := c.containers[tagName]
				for _, cont := range conts {
					// if the container is not running (not in running state), we must create another one
					if !isContainerRunning(cont.Name) {
						cleanupContainer(cont.Name, cont.pid, cont.dir)
						errs := make([]error, 0)
						newContainer := c.createContainer(cont.Tag)
						if len(errs) != 0 {
							log := ""
							for _, e := range errs {
								log += e.Error()
							}

							panic(log)
						}

						// a balancer is listening on watchCh and updates its containers list
						watchCh <- RestartedContainer{
							Name:             newContainer.Name,
							OldContainerName: cont.Name,
						}

						c.Lock()
						// remove the old container and put the new one in
						newContainers := make([]container, 0)
						for _, a := range conts {
							if a.Name != cont.Name {
								newContainers = append(newContainers, a)
							}
						}

						c.containers[tagName] = newContainers
						c.containers[tagName] = append(c.containers[tagName], newContainer)
						c.Unlock()

						break
					}
				}
			}
		}
	}()

	return watchCh
}

func (c *ContainerFactory) createContainer(tag string) container {
	name := uuid.New().String()

	// containerDir is within executionDir that the user gives
	containerDir := fmt.Sprintf("%s/%s", c.executionDirectory, name)
	fsErr := os.Mkdir(containerDir, os.ModePerm)

	if fsErr != nil {
		panic(fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", fsErr.Error())))
	}

	pid, err := executeContainer(name, tag, c.executionDirectory)
	newContainer := container{
		pid:  pid,
		dir:  containerDir,
		Tag:  tag,
		Name: name,
	}

	if err != nil {
		panic(fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", err.Error())))
	}

	// we update the containers array right away
	// so if something goes wrong, the close mechanism
	// can clenaup the system from the bad container

	// NOTE: A container might be up but in a not "running" state.
	// That is why we need to put it into the containers array.
	// It is also important for the cleanup since volume directories
	// are created and need to be removed in case of any error
	c.Lock()
	if _, ok := c.containers[tag]; !ok {
		c.containers[tag] = make([]container, 0)
	}

	c.containers[tag] = append(c.containers[tag], newContainer)
	c.Unlock()

	return newContainer
}

func cleanupContainer(name string, pid int, dir string) {
	stopDockerContainer(name, pid)

	err := os.RemoveAll(dir)

	if err != nil {
		cmd := exec.Command("rm", []string{"-rf", dir}...)

		err := cmd.Run()

		if err != nil {
			panic(fmt.Sprintf("Filesystem error: Cannot remove directory %s: %v. You will have to remove in manually\n", dir, err))
		}
	}
}

func executeContainer(containerName, containerTag, executionDir string) (int, error) {
	args := []string{
		"run",
		"-d",
		"-t",
		"--network=none",
		"-v",
		fmt.Sprintf("%s:/app:rw", fmt.Sprintf("%s/%s", executionDir, containerName)),
		"--name",
		containerName,
		containerTag,
		"/bin/sh",
	}

	cmd := exec.Command("docker", args...)
	var outb, errb bytes.Buffer

	cmd.Stderr = &errb
	cmd.Stdout = &outb

	startErr := cmd.Run()

	if startErr != nil {
		return 0, startErr
	}

	return cmd.Process.Pid, nil
}
