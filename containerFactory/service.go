package containerFactory

import (
	"bytes"
	"context"
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
	pid    chan int
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

				newContainer := container{
					output: make(chan message),
					pid:    make(chan int),
					dir:    containerDir,
					Tag:    tag,
					Name:   name,
				}

				createContainer(newContainer, executionDir)

				select {
				case <-time.After(1 * time.Second):
					// we update the containers array right away
					// so if something goes wrong, the close mechanism
					// can clenaup the system from the bad container
					lock.Lock()
					if _, ok := containers[tag]; !ok {
						containers[tag] = make([]container, 0)
					}

					containers[tag] = append(containers[tag], newContainer)
					lock.Unlock()

					if !isContainerRunning(name) {
						errs = append(errs, fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))

						wg.Done()

						return
					}

					close(newContainer.output)
					lock.Lock()
					if _, ok := containers[tag]; !ok {
						containers[tag] = make([]container, 0)
					}

					containers[tag] = append(containers[tag], newContainer)
					lock.Unlock()

					wg.Done()
				case msg := <-newContainer.output:
					if msg.messageType == "error" {
						err := msg.data.(error)
						close(newContainer.output)

						errs = append(errs, fmt.Errorf("%w: %s", ContainerStartupTimeout, fmt.Sprintf("Could not start container; Name: %s, Tag: %s: %s", newContainer.Name, newContainer.Tag, err.Error())))

						wg.Done()
					}
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
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))
			out := make(chan int)

			go func(pidCh chan int, out chan int) {
				select {
				case <-ctx.Done():
					out <- 0
				case pid := <-c.pid:
					cancel()
					out <- pid
				}
			}(c.pid, out)

			pid := <-out

			stopDockerContainer(c.Name, pid)

			close(c.pid)

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

func createContainer(c container, executionDir string) {
	go func(c container) {
		args := []string{
			"run",
			"-d",
			"-t",
			"--network=none",
			"-v",
			fmt.Sprintf("%s:/app:rw", fmt.Sprintf("%s/%s", executionDir, c.Name)),
			"--name",
			"--init",
			c.Name,
			c.Tag,
			"/bin/sh",
		}

		cmd := exec.Command("docker", args...)
		var outb, errb bytes.Buffer

		cmd.Stderr = &errb
		cmd.Stdout = &outb

		startErr := cmd.Run()
		c.pid <- cmd.Process.Pid

		if startErr != nil {
			c.output <- message{
				messageType: "error",
				data:        startErr,
			}

			return
		}
	}(c)
}
