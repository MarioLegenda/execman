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

type Container interface {
	CreateContainers(string, string, int) []error
	Close()
	Containers(tagName string) []container
}

var containers = make(map[string][]container)
var lock sync.Mutex

type message struct {
	messageType string
	data        interface{}
}

type container struct {
	/**
	outuput channel is used only to signal if a container could
	boot up or not and on creation time. After that, it is closed and
	not used anymore. see CreateContainers function
	*/
	output chan message
	pid    int
	dir    string

	Tag  string
	Name string
}

func Containers(tagName string) []container {
	return containers[tagName]
}

func CreateContainers(executionDir, tag string, containerNum int) []error {
	blocks := makeBlocks(containerNum, 5)

	errs := make([]error, 0)
	for _, block := range blocks {
		wg := sync.WaitGroup{}

		for range block {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				name := uuid.New().String()

				containerDir := fmt.Sprintf("%s/%s", executionDir, name)
				fsErr := os.Mkdir(containerDir, os.ModePerm)

				if fsErr != nil {
					errs = append(errs, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", fsErr.Error())))

					wg.Done()

					return
				}

				pid, err := createContainer(name, tag, executionDir)
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
					errs = append(errs, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Could not start container: %s", err.Error())))

					wg.Done()

					return
				}

				select {
				case <-time.After(1 * time.Second):
					if !isContainerRunning(name) {
						errs = append(errs, fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))

						wg.Done()

						return
					}

					wg.Done()
				}
			}(&wg)
		}

		wg.Wait()
	}

	time.Sleep(2 * time.Second)

	return errs
}

func Close() {
	contArr := containersToSlice(containers)

	wg := sync.WaitGroup{}
	for _, entry := range contArr {
		wg.Add(1)

		go func(c container, wg *sync.WaitGroup) {
			stopDockerContainer(c.Name, c.pid)

			err := os.RemoveAll(c.dir)

			if err != nil {
				cmd := exec.Command("rm", []string{"-rf", c.dir}...)

				err := cmd.Run()

				if err != nil {
					fmt.Printf("Filesystem error: Cannot remove directory %s: %v. You will have to remove in manually\n", c.dir, err)
					wg.Done()
					// TODO: send slack error and log
					return
				}
			}

			wg.Done()
		}(entry, &wg)
	}

	wg.Wait()
	containers = make(map[string][]container)
}

func createContainer(containerName, containerTag, executionDir string) (int, error) {
	args := []string{
		"run",
		"-d",
		"-t",
		"--network=none",
		"-v",
		fmt.Sprintf("%s:/app:rw", fmt.Sprintf("%s/%s", executionDir, containerName)),
		"--name",
		"--init",
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
