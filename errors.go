package execman

import "errors"

var ContainerCannotBoot = errors.New("container cannot boot up")
var InvalidOptions = errors.New("invalid options")
