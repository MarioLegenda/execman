package containerFactory

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

func isContainerRunning(name string) bool {
	cmd := exec.Command("docker", []string{
		"container",
		"inspect",
		"-f",
		"'{{.State.Status}}'",
		name,
	}...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	o := strings.Trim(string(out), " ")

	match, _ := regexp.MatchString("running", o)

	return match
}

func stopDockerContainer(containerName string, pid int) {
	if pid == -1 {
		fmt.Println("Something went wrong. PID should not be -1")
		log.Fatalln("Cannot kill a container with PID -1")
	}

	var stopCmd *exec.Cmd

	stopCmd = exec.Command("docker", []string{"container", "stop", containerName}...)
	stopErr := stopCmd.Run()

	if stopErr == nil {
		var rmCmd *exec.Cmd

		rmCmd = exec.Command("docker", []string{"remove", containerName}...)
		rmErr := rmCmd.Run()

		if rmErr != nil {

			killErr := syscall.Kill(pid, 9)

			if killErr != nil {
				panic(fmt.Sprintf("Could not stop container %s with PID %d: %s", containerName, pid, killErr.Error()))
			}
		}
	}

	if stopErr != nil {
		panic(fmt.Sprintf("Could not stop container %s with PID %d: %s", containerName, pid, stopErr.Error()))
	}
}

func makeBlocks(num int, delimiter int) [][]int {
	portions := num / delimiter
	leftover := num % delimiter
	if leftover != 0 {
		portions++
	}

	blocks := make([][]int, 0)
	current := 0
	for i := 0; i < portions; i++ {
		b := make([]int, 0)
		d := delimiter

		if i == portions-1 && leftover != 0 {
			d = leftover
		}

		for a := 0; a < d; a++ {
			b = append(b, current)
			current++
		}

		blocks = append(blocks, b)
	}

	return blocks
}

func containersToSlice(containers map[string][]container) []container {
	s := make([]container, 0)
	for _, v := range containers {
		s = append(s, v...)
	}

	return s
}
