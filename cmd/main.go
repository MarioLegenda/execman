package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"github.com/MarioLegenda/execman/types"
)

func main() {
	emulator := execman.New(execman.Options{
		CLang: execman.CLang{
			Workers:    10,
			Containers: 1,
		},
		Ruby: execman.Ruby{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	res := emulator.RunJob(string(types.CLang.Name), `
#include <stdio.h>

int main() {
	printf("Hello World");

    return 0;
}`)
	fmt.Println(res)

	res = emulator.RunJob(string(types.Ruby.Name), `puts "Hello world"`)
	fmt.Println(res)
}
