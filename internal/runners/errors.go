package runners

import "errors"

var WorkerShutdownError = errors.New("worker is shutting down")
var ExecutionFailed = errors.New("execution failed")
var CodeExecutionTimeout = errors.New("code execution timed out")
var FilesystemError = errors.New("filesystem error")
