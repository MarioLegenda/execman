package containerFactory

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"sync"
	"time"
)

type container struct {
	pid int
	dir string

	Tag  string
	Name string
}

type RestartedContainer struct {
	Name string
	Tag  string
}

var containers = make(map[string][]container)
var lock sync.Mutex
var executionDirectory string

func Containers(tagName string) []container {
	return containers[tagName]
}

func CreateContainers(executionDir, tag string, containerNum int) []error {
	blocks := makeBlocks(containerNum, 5)
	executionDirectory = executionDir

	errs := make([]error, 0)
	for _, block := range blocks {
		wg := sync.WaitGroup{}

		for range block {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()

				newContainer := createContainer(tag, executionDir, &errs)
				if len(errs) != 0 {
					return
				}

				select {
				case <-time.After(5 * time.Second):
					if !isContainerRunning(newContainer.Name) {
						errs = append(errs, fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))

						return
					}
				}
			}(&wg)
		}

		wg.Wait()
	}

	return errs
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
					if !isContainerRunning(c.Name) {
						fmt.Println("Container is not running. Creating another one")
						cleanupContainer(c.Name, c.pid, c.dir)
						errs := make([]error, 0)
						newContainer := createContainer(c.Tag, executionDirectory, &errs)
						if len(errs) != 0 {
							return
						}

						watchCh <- RestartedContainer{
							Name: c.Name,
							Tag:  c.Tag,
						}

						select {
						case <-time.After(5 * time.Second):
							if !isContainerRunning(newContainer.Name) {
								errs = append(errs, fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))

								return
							}
						}
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

func createContainer(tag, executionDir string, errs *[]error) container {
	name := uuid.New().String()

	containerDir := fmt.Sprintf("%s/%s", executionDir, name)
	fsErr := os.Mkdir(containerDir, os.ModePerm)

	if fsErr != nil {
		*errs = append(*errs, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", fsErr.Error())))

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
		*errs = append(*errs, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", err.Error())))

		return container{}
	}

	return newContainer
}
