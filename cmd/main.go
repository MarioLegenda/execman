package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"github.com/MarioLegenda/execman/types"
	"sync"
)

func main() {
	emulator := execman.New(execman.Options{
		Ruby: execman.Ruby{
			Workers:    30,
			Containers: 10,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			wg.Done()

			res := emulator.RunJob(string(types.Ruby.Name), `puts "Hello world"`)
			fmt.Println(res.Success)
		}()
	}

	wg.Wait()

	emulator.Close()
}
