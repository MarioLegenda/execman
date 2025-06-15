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
			Workers:    10,
			Containers: 10,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()
	wg := sync.WaitGroup{}
	failed := 0
	lock := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			res := instance.Run(execman.RubyLang, `puts "Hello world"`)
			if !res.Success {
				lock.Lock()
				failed++
				lock.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Elapsed time: ", time.Since(now))
	fmt.Println("Number of failed jobs: ", failed)

	instance.Close()
}

func tickerImplementation() {
	instance, err := execman.New(execman.Options{
		Ruby: execman.Ruby{
			Workers:    100,
			Containers: 10,
		},
		NodeLts: execman.NodeLts{
			Workers:    100,
			Containers: 10,
		},
		ExecutionDirectory: "/home/mario/go/execman/execution_directory",
	})

	if err != nil {
		log.Fatalln(err)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	elapsedTicker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ticker.C:

			rubyRes := instance.Run(execman.RubyLang, `puts "Hello world"`)
			nodeRes := instance.Run(execman.NodeLatestLang, `console.log("Hello world")`)
			fmt.Println(fmt.Sprintf("Ruby success: %v; Node success: %v", rubyRes.Success, nodeRes.Success))

		case <-elapsedTicker.C:
			instance.Close()
			return
		}
	}
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
