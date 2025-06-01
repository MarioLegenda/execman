package execman

import "errors"

var ContainerCannotBoot = errors.New("container cannot boot up")
var TimeoutError = errors.New("a timeout occurred")
