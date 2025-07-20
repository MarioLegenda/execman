package containerFactory

import "errors"

var ContainerCannotBoot = errors.New("container cannot boot up")
var ContainerStartupTimeout = errors.New("container startup timeout")
