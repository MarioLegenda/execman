package main

import (
	"github.com/MarioLegenda/ellie"
)

func main() {
	emulator := execman.New(execman.Options{
		Ruby: execman.Ruby{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/emulator/simulator/execution_directory",
		LogDirectory:       "/home/mario/go/emulator/simulator/var",
	})
}
