package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"sync"
	"time"
)

func main() {
	emulator, _ := execman.New(execman.Options{
		CPlus: execman.CPlus{
			Workers:    1000,
			Containers: 10,
		},
		Haskell: execman.Haskell{
			Workers:    1000,
			Containers: 10,
		},
		GoLang: execman.GoLang{
			Workers:    1000,
			Containers: 10,
		},
		Perl: execman.Perl{
			Workers:    1000,
			Containers: 10,
		},
		Lua: execman.Lua{
			Workers:    1000,
			Containers: 10,
		},
		CSharp: execman.CSharp{
			Workers:    1000,
			Containers: 10,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	now := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_ = emulator.Run(execman.CPlusPlusLang, `
#include <iostream>

int main() {
    std::cout << "Hello world";
    return 0;
}`)
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(now))

	emulator.Close()
}
