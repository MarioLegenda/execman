package runners

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type JavaExecParams struct {
	ContainerName      string
	ExecutionDirectory string
	ContainerDirectory string
	ExecutionFile      string
	Timeout            int
}

func javaRunner(params JavaExecParams) Result {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(params.Timeout)*time.Second))
	defer cancel()

	var outb, errb string
	var runResult Result

	tc := make(chan string)
	pidC := make(chan int, 1)

	go func() {
		// TODO: /bin/sh does not want to cd into any directory but ls works fine. /app directory is also there.
		// Everything should work but /bin/sh won't cd into directory. If I enter the container directly,
		// everything works fine and everything is where it should be. Multiple images tried, not just amazon
		cmd := exec.Command("docker", []string{"exec", params.ContainerName, "/bin/sh", "-c", fmt.Sprintf("cd /app/%s && javac %s && java %s", params.ContainerDirectory, params.ExecutionFile, params.ExecutionFile)}...)

		errPipe, err := cmd.StderrPipe()

		if err != nil {
			runResult.Error = ExecutionFailed

			tc <- "error"

			return
		}

		outPipe, err := cmd.StdoutPipe()

		if err != nil {
			runResult.Error = ExecutionFailed

			tc <- "error"

			return
		}

		startErr := cmd.Start()
		pidC <- cmd.Process.Pid

		a, _ := io.ReadAll(errPipe)
		b, _ := io.ReadAll(outPipe)
		errb = string(a)
		outb = string(b)

		if startErr == nil {
			waitErr := cmd.Wait()

			if waitErr != nil {
				runResult.Error = ExecutionFailed

				tc <- "error"

				return
			}
		}

		if startErr != nil {
			runResult.Error = ExecutionFailed

			tc <- "error"

			return
		}

		tc <- "finished"
	}()

	select {
	case res := <-tc:
		if res == "error" {
			out := makeRunDecision(errb, outb, params.ExecutionDirectory)
			if out != "" {
				runResult.Result = out
				runResult.Error = nil
			}

			if errb != "" {
				runResult.Success = false
			} else {
				runResult.Success = true
			}

			destroyContainerProcess(extractUniqueIdentifier(params.ContainerDirectory, false), true)
			destroy(params.ExecutionDirectory)
			return runResult
		}

		out := makeRunDecision(errb, outb, params.ExecutionDirectory)
		runResult.Success = true
		runResult.Result = out
		runResult.Error = nil

		closeExecSession(<-pidC)
		destroy(params.ExecutionDirectory)

		break
	case <-ctx.Done():
		destroyContainerProcess(extractUniqueIdentifier(params.ContainerDirectory, false), true)
		closeExecSession(<-pidC)
		destroy(params.ExecutionDirectory)
		close(pidC)
		return Result{
			Result:  "",
			Success: false,
			Error:   CodeExecutionTimeout,
		}
	}

	runResult.Error = nil

	return runResult
}
