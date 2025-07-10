package runners

import (
	"fmt"
	"github.com/MarioLegenda/execman/internal/builders"
	"os/exec"
)

type Params struct {
	ExecutionDir string

	Timeout int

	ContainerName string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string
	PackageName       string
}

func Run(params Params) Result {
	build, err := builders.Build(builders.InitBuildParams(
		params.EmulatorExtension,
		params.EmulatorText,
		fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
	))

	if err != nil {
		return Result{
			Result:  "",
			Success: false,
			Error:   err,
		}
	}

	if params.EmulatorName == string(nodeLts.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "node", process}...)
		})
	}

	if params.EmulatorName == string(perlLts.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "perl", "-I", containerDirectory, process}...)
		})
	}

	if params.EmulatorName == string(luaLts.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "lua", process}...)
		})
	}

	if params.EmulatorName == string(nodeEsm.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "node", process}...)
		})
	}

	if params.EmulatorName == string(goLang.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", containerName, "/bin/bash", "-c", fmt.Sprintf("cd %s && go mod init app/%s >/dev/null 2>&1 && go build && ./%s", containerDirectory, containerDirectory, containerDirectory)}...)
		})
	}

	if params.EmulatorName == string(ruby.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "ruby", process}...)
		})
	}

	if params.EmulatorName == string(php.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "php", process}...)
		})
	}

	if params.EmulatorName == string(python2.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("/app/%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "python", process}...)
		})
	}

	if params.EmulatorName == string(python3.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("/app/%s/%s", containerDirectory, executionFile)
			return exec.Command("docker", []string{"exec", containerName, "python3", process}...)
		})
	}

	if params.EmulatorName == string(csharpMono.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", containerName, "/bin/bash", "-c", fmt.Sprintf("cd %s && mcs -out:%s.exe -pkg:dotnet -recurse:'*.cs' && mono %s.exe", containerDirectory, containerDirectory, containerDirectory)}...)
		})
	}

	if params.EmulatorName == string(haskell.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", containerName, "/bin/bash", "-c", fmt.Sprintf("cd %s && ghc %s > output.txt && ./%s > output.txt", containerDirectory, executionFile, executionFile[:len(executionFile)-3])}...)
		})
	}

	if params.EmulatorName == string(cLang.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf(
				"cd %s && gcc -o %s %s > output.txt && ./%s > output.txt",
				containerDirectory,
				containerDirectory,
				executionFile,
				containerDirectory,
			)

			return exec.Command("docker", []string{"exec", params.ContainerName, "/bin/sh", "-c", process}...)
		})
	}

	if params.EmulatorName == string(cPlus.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf(
				"cd %s && g++ -o %s %s > output.txt && ./%s",
				containerDirectory,
				containerDirectory,
				executionFile,
				containerDirectory,
			)
			return exec.Command("docker", []string{"exec", params.ContainerName, "/bin/sh", "-c", process}...)
		})
	}

	if params.EmulatorName == string(rust.name) {
		// since rust build step is different from the rest of them, it is used
		// as a specific build step and is overriden.
		build, err := builders.RustSingleFileBuild(builders.InitRustParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", containerName, "/bin/bash", "-c", fmt.Sprintf("cd %s && cargo run --quiet | tee output.txt", containerDirectory)}...)
		})
	}

	if params.EmulatorName == string(julia.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf("%s/%s", containerDirectory, executionFile)

			return exec.Command("docker", []string{"exec", containerName, "julia", process}...)
		})
	}

	if params.EmulatorName == string(swift.name) {
		return swiftRunner(SwiftExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(kotlin.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", params.ContainerName, "/bin/sh", "-c", fmt.Sprintf("cd %s && kotlinc %s -include-runtime -d %s.jar && java -jar %s.jar", containerDirectory, executionFile, containerDirectory, containerDirectory)}...)
		})
	}

	if params.EmulatorName == string(java.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", containerName, "/bin/sh", "-c", fmt.Sprintf("cd /app/%s && javac %s && java %s", containerDirectory, executionFile, executionFile)}...)
		})
	}

	if params.EmulatorName == string(zigLts.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			return exec.Command("docker", []string{"exec", containerName, "/bin/sh", "-c", fmt.Sprintf("cd %s && zig run %s", containerDirectory, executionFile)}...)
		})
	}

	if params.EmulatorName == string(bash.name) {
		return runner(RunnerParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		}, func(containerName, executionDirectory, executionFile, containerDirectory string) *exec.Cmd {
			process := fmt.Sprintf(
				"cd %s && bash %s",
				containerDirectory,
				executionFile,
			)
			return exec.Command("docker", []string{"exec", containerName, "/bin/sh", "-c", process}...)
		})
	}

	return Result{}
}
