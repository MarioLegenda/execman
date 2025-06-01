package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"github.com/MarioLegenda/execman/types"
)

func main() {
	emulator := execman.New(execman.Options{
		Ruby: execman.Ruby{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})
	defer emulator.Close()

	res := emulator.RunJob(string(types.Ruby.Name), `puts "Hello world"`)
	fmt.Println(res)
}
