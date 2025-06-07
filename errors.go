package execman

import "errors"

var ContainerCannotBoot = errors.New("container cannot boot up")
var BalancerNotCreated = errors.New("balancer not created")
