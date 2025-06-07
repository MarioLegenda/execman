package runners

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"
)

type GoProjectExecParams struct {
	ContainerName      string
	ExecutionDirectory string
	ContainerDirectory string
}

func goProjectRunner(params GoProjectExecParams) Result {
	var outb, errb string
	var runResult Result

	tc := make(chan string)
	pidC := make(chan int, 1)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	go func() {
		cmd := exec.Command("docker", []string{"exec", params.ContainerName, "/bin/bash", "-c", fmt.Sprintf("cd %s && go mod init app/%s >/dev/null 2>&1 && go build && ./%s", params.ContainerDirectory, params.ContainerDirectory, params.ContainerDirectory)}...)
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
				runResult.Error = fmt.Errorf("%w: %s", ExecutionFailed, fmt.Sprintf("Execution failed: %s", waitErr.Error()))

				tc <- "error"

				return
			}
		}

		if startErr != nil {
			runResult.Error = fmt.Errorf("%w: %s", ExecutionFailed, fmt.Sprintf("Execution failed: %s", startErr.Error()))

			tc <- "error"

			return
		}

		tc <- "finished"
	}()

	select {
	case res := <-tc:
		if res == "error" {
			out := makeRunDecision(errb, outb, params.ExecutionDirectory)
			runResult.Success = true
			runResult.Result = out
			runResult.Error = nil

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
