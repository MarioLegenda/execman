# Introduction

**execman** is a tool to execute code in different programming languages (for now, 14
different programming languages). It came from my abandoned project; a blogging platform
specifically designed for software developers where developers could write blogs with
code snippets that could be executed on the backend and for the result to be shown on
the frontend. The project didn't go anywhere but the microservice that I created to execute
code on the backend is this project. 

# Why?

From what I know, most of the websites that implement some kind of code execution, boot up
a container and execute it inside it. In my case (the blogging platform from above), that proved
as wasteful since booting up a container took time and memory. Roughly, every container (every code
execution) took around 50MB of memory to boot the container into a running state and that is before starting to run
the code. Also, some language like _C_ or _Go_ require you to compile them and then run the code which
takes time and memory. 

# How execman works

**execman** creates containers before you run your code and then runs them concurrently with
**docker exec** command. In front of the containers (you specify the number of containers to run), 
there is a load balancer with a certain number of workers (which you also specify). Running code
is load balanced on the created containers and executed with **docker exec**. That means that if users
send 10 requests in parallel, those requests will be executed in parallel on any number of containers
that you specify to boot up. 

For example, if you have 10 workers and 1 container for the Ruby language, 10 requests will be executed
in parallel on this one container in roughly the time that it takes to actually execute them on your own local
machine without booting up your the container since it is already running.

# What execman isn't

It is not a long-running process. It is meant to be created, run some code and closed. 
Use it in a long-running process such as an HTTP server. Also, shutting it down with a SIGTERM
such as with Ctrl+C will not clean up containers that are created so if you start running it
and close it abruptly, those containers will still remain up so keep that in mind. 

It is also not meant to be used inside a docker container since every container needs its own
volume in order to execute code with it. 

# How to use it

### Installing docker images

First, you need to install the docker images. _cd_ into _dockerImages_.

`cd dockerImages`

And then build them

`bash run_all.sh`

Depending on your network connection and the specs of your computer, this will take a very
long time. On my computer, it takes around 10 minutes and takes up around 3GB on the drive. 

### Running code

First, you have to create a new **execman** instance:

````go
package main

import (
    "execman"
    "fmt"
)

instance, err := execman.New(execman.Options{
    Ruby: execman.Ruby{
        Workers:    10,
        Containers: 1,
    },
    ExecutionDirectory: "/home/user/go/my_package/execution_directory",
})

res := emulator.Run(execman.RubyLang, `puts "Hello world"`)

fmt.Println(res)

instance.Close()
````

`res` is in instance of `execman.Result` and it looks like this:

````go
type Result struct {
    Result  string
    Success bool
    Error   error
}
````

**ExecutionDirectory** is the volume directory where the files to run will be placed and
ran by **docker exec**. It should be writeable. It can be anywhere you want but it is
best to be close to the application that you are working on. 

> [!CAUTION]
> You have to call _instance.Close()_. This will stop all running containers
> and workers that make the system that it is. If you don't do this, these
> containers will still remain running and the volume directories will also
> not get cleaned up. 

This example uses Ruby, but it can be any language you want as long as its supported. 
Currently, **execman** supports languages that you can find in the _dockerImages_ directory.

For example, here is a concurrent example of using it with many languages with
100 concurrent requests to it:

````go
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

````


