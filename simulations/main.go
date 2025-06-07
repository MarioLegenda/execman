package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"log"
	"sync"
	"time"
)

func main() {
	emulator, err := execman.New(execman.Options{
		GoLang: execman.GoLang{
			Workers:    10,
			Containers: 1,
		},
		Ruby: execman.Ruby{
			Workers:    10,
			Containers: 1,
		},
		Rust: execman.Rust{
			Workers:    10,
			Containers: 1,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_ = emulator.Run(execman.RubyLang, `
package main

use "fmt"

func main() {
	fmt.Println("Hello world")
}
`)
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(now))

	emulator.Close()
}
