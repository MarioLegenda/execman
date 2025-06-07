package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"sync"
)

func main() {
	emulator, _ := execman.New(execman.Options{
		CPlus: execman.CPlus{
			Workers:    5,
			Containers: 10,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			wg.Done()

			res := emulator.Run(execman.CPlusPlusLang, `
#include <iostream>

int main() {
    std::cout << "Hello world";
    return 0;
}`)

			fmt.Println(res)
		}()
	}

	wg.Wait()

	emulator.Close()
}
