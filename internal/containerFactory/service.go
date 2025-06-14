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

// holds containers under a specific tag name of the language
// that the user has selected to run
var containers = make(map[string][]container)

// holds all locks of this package
var lock sync.Mutex

// holds the execution directory in which container
// volumes reside
var executionDirectory string

// returns all containers by tag name, for example, rust:rust
func Containers(tagName string) []container {
	return containers[tagName]
}

// Runs a goroutine per containerNum and creates a container with the specified tag
func CreateContainers(executionDir, tag string, containerNum int) {
	executionDirectory = executionDir

	wg := sync.WaitGroup{}
	wg.Add(containerNum)
	for i := 0; i < containerNum; i++ {
		go func() {
			defer wg.Done()

			newContainer := createContainer(tag, executionDir)

			if !isContainerRunning(newContainer.Name) {
				panic(fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))
			}
		}()
	}

	wg.Wait()
}

func Close() {
	lock.Lock()
	contArr := containersToSlice(containers)
	lock.Unlock()

	wg := sync.WaitGroup{}
	for _, entry := range contArr {
		wg.Add(1)

		go func(c container, wg *sync.WaitGroup) {
			defer wg.Done()
			cleanupContainer(c.Name, c.pid, c.dir)
		}(entry, &wg)
	}

	wg.Wait()
	containers = make(map[string][]container)
}

// This code has to work and must panic if it does not
func WatchContainers(tagName string, done chan interface{}) chan RestartedContainer {
	watchCh := make(chan RestartedContainer)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				conts := containers[tagName]
				for _, c := range conts {
					// if the container is not running (not in running state), we must create another one
					if !isContainerRunning(c.Name) {
						cleanupContainer(c.Name, c.pid, c.dir)
						errs := make([]error, 0)
						newContainer := createContainer(c.Tag, executionDirectory)
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
							OldContainerName: c.Name,
						}

						lock.Lock()
						// remove the old container and put the new one in
						newContainers := make([]container, 0)
						for _, a := range conts {
							if a.Name != c.Name {
								newContainers = append(newContainers, a)
							}
						}

						containers[tagName] = newContainers
						containers[tagName] = append(containers[tagName], newContainer)
						lock.Unlock()

						break
					}
				}
			}
		}
	}()

	return watchCh
}

func cleanupContainer(name string, pid int, dir string) {
	stopDockerContainer(name, pid)

	err := os.RemoveAll(dir)

	if err != nil {
		cmd := exec.Command("rm", []string{"-rf", dir}...)

		err := cmd.Run()

		if err != nil {
			fmt.Printf("Filesystem error: Cannot remove directory %s: %v. You will have to remove in manually\n", dir, err)
			return
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

func createContainer(tag, executionDir string) container {
	name := uuid.New().String()

	// containerDir is within executionDir that the user gives
	containerDir := fmt.Sprintf("%s/%s", executionDir, name)
	fsErr := os.Mkdir(containerDir, os.ModePerm)

	if fsErr != nil {
		panic(fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", fsErr.Error())))

		return container{}
	}

	pid, err := executeContainer(name, tag, executionDir)
	newContainer := container{
		pid:  pid,
		dir:  containerDir,
		Tag:  tag,
		Name: name,
	}

	// we update the containers array right away
	// so if something goes wrong, the close mechanism
	// can clenaup the system from the bad container

	// NOTE: A container might be up but in a not "running" state.
	// That is why we need to put it into the containers array.
	// It is also important for the cleanup since volume directories
	// are created and need to be removed in case of any error
	lock.Lock()
	if _, ok := containers[tag]; !ok {
		containers[tag] = make([]container, 0)
	}

	containers[tag] = append(containers[tag], newContainer)
	lock.Unlock()

	if err != nil {
		panic(fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", err.Error())))

		return container{}
	}

	return newContainer
}
