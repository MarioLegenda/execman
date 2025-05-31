package containerFactory

import (
	"bytes"
	"context"
	"emulator/pkg/appErrors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"sync"
	"time"
)

var services map[string]Container

type Container interface {
	CreateContainers(string, string, int) []*appErrors.Error
	Close()
	Containers(tagName string) []container
}

type service struct {
	containers map[string][]container
	lock       sync.Mutex
}

type message struct {
	messageType string
	data        interface{}
}

type container struct {
	output chan message
	pid    chan int
	dir    string

	Tag  string
	Name string
}

func Init(name string) {
	if services == nil {
		services = make(map[string]Container)
	}

	s := &service{containers: make(map[string][]container)}

	services[name] = s
}

func Service(name string) Container {
	return services[name]
}

func (d *service) Containers(tagName string) []container {
	return d.containers[tagName]
}

func (d *service) CreateContainers(executionDir, tag string, containerNum int) []*appErrors.Error {
	blocks := makeBlocks(containerNum, 5)

	errs := make([]*appErrors.Error, 0)
	for _, block := range blocks {
		wg := sync.WaitGroup{}

		for _ = range block {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				name := uuid.New().String()

				containerDir := fmt.Sprintf("%s/%s", executionDir, name)
				fsErr := os.Mkdir(containerDir, os.ModePerm)

				if fsErr != nil {
					errs = append(errs, appErrors.New(appErrors.ApplicationError, appErrors.ApplicationRuntimeError, fmt.Sprintf("Could not start container: %s", fsErr.Error())))

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
					if !isContainerRunning(name) {
						errs = append(errs, appErrors.New(appErrors.ApplicationError, appErrors.ApplicationRuntimeError, fmt.Sprintf("Container startup timeout: Tag: %s, Name: %s", newContainer.Tag, newContainer.Name)))

						wg.Done()

						return
					}

					close(newContainer.output)
					d.lock.Lock()
					if _, ok := d.containers[tag]; !ok {
						d.containers[tag] = make([]container, 0)
					}

					d.containers[tag] = append(d.containers[tag], newContainer)
					d.lock.Unlock()

					wg.Done()
				case msg := <-newContainer.output:
					if msg.messageType == "error" {
						err := msg.data.(error)
						close(newContainer.output)

						errs = append(errs, appErrors.New(appErrors.ApplicationError, appErrors.ApplicationRuntimeError, fmt.Sprintf("Could not start container; Name: %s, Tag: %s: %s", newContainer.Name, newContainer.Tag, err.Error())))

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

func (d service) Close() {
	contArr := containersToSlice(d.containers)
	blocks := makeBlocks(len(contArr), 10)

	for _, block := range blocks {
		wg := sync.WaitGroup{}

		for _, b := range block {
			wg.Add(1)
			c := contArr[b]

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

				if err == nil {
					// TODO: // send slack error and log
				}

				if err != nil {
					cmd := exec.Command("rm", []string{"-rf", c.dir}...)

					err := cmd.Run()

					if err != nil {
						wg.Done()
						// TODO: send slack error and log
						return
					}
				}

				wg.Done()
			}(c, &wg)
		}

		wg.Wait()
	}

	cmd := exec.Command("docker", []string{"volume", "rm", "$(docker volume ls -q)"}...)
	cmd.Run()
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
			c.Name,
			"--init",
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
