package main

import (
	"fmt"
	"github.com/MarioLegenda/execman"
	"log"
	"sync"
	"time"
)

func singleIterations() {
	instance, err := execman.New(execman.Options{
		Ruby: execman.Ruby{
			Workers:    125,
			Containers: 10,
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
			res := instance.Run(execman.RubyLang, `puts "Hello world"`)

			fmt.Println(res.Success)
		}()
	}

	wg.Wait()

	fmt.Println(time.Since(now))

	instance.Close()
}

func averageTime() {
	instance, err := execman.New(execman.Options{
		Ruby: execman.Ruby{
			Workers:    125,
			Containers: 10,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	if err != nil {
		log.Fatalln(err)
	}

	averages := make([]time.Duration, 0)
	lock := sync.Mutex{}
	for a := 0; a < 10; a++ {
		fmt.Println("Iteration: ", a)

		now := time.Now()
		wg := sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				_ = instance.Run(execman.RubyLang, `puts "Hello world"`)
				lock.Lock()
				averages = append(averages, time.Since(now))
				lock.Unlock()
			}()
		}

		wg.Wait()
	}

	var total time.Duration
	for _, d := range averages {
		total += d
	}
	average := total / time.Duration(len(averages))

	fmt.Printf("Average duration: %s\n", average)

	instance.Close()
}
