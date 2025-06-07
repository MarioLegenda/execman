package containerFactory

import "errors"

var ApplicationError = errors.New("an application error occurred")
var ContainerCannotBoot = errors.New("container cannot boot up")
var ContainerStartupTimeout = errors.New("container startup timeout")
var TimeoutError = errors.New("a timeout occurred")
var ExecutionFailed = errors.New("execution failed")
var CodeExecutionTimeout = errors.New("code execution timed out")
var FilesystemError = errors.New("filesystem error")
