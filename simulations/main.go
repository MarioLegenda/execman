package main

import (
	"github.com/MarioLegenda/execman"
	"log"
	"sync"
)

func main() {
	instance, err := execman.New(execman.Options{
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

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_ = instance.Run(execman.RubyLang, `puts "Hello world"`)
			_ = instance.Run(execman.Golang, `
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
`)
			_ = instance.Run(execman.RustLang, `
fn main() {
    println!("Hello world");
}
`)
		}()
	}

	wg.Wait()

	instance.Close()
}
