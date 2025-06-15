# Table of Contents

- [Introduction](#introduction)
- [Why](#why)
- [How execman works](#how-execman-works)
- [What execman isn't](#what-execman-isnt)
- [How to use it](#how-to-use-it)
    - [Installation](#installation)
    - [Installing docker images](#installing-docker-images)
    - [Running code](#running-code)
- [About workers and containers](#about-workers-and-containers)
- [Container healing](#container-healing)
- [Timeouts](#timeouts)
- [Available programming languages](#available-programming-languages)
- [Tests](#tests)

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
is load balanced by least number of jobs on the created containers and executed with **docker exec**. 
That means that if users send 10 requests in parallel, those requests will be executed in parallel on any number of containers
that you specify to boot up. Load balancing on containers is also balanced by least amount of jobs.

For example, if you have 10 workers and 1 container for the Ruby language, 10 requests will be executed
in parallel on this one container in roughly the time that it takes to actually execute them on your own local
machine without booting up your container since it is already running.

# What execman isn't

It is not a long-running process. It is meant to be created, run some code and closed. 
Use it in a long-running process such as an HTTP server. Also, shutting it down with a SIGTERM
such as with Ctrl+C will not clean up containers that are created so if you start running it
and close it abruptly, those containers will still remain up so keep that in mind. 

It is also **not meant to be used inside a docker container** since every container needs its own
volume in order to execute code with it.

**execman** does not listen to any signals like SIGTERM or SIGKILL or any other to terminate
the process. If you have such requirements, you should add them yourself. 

# How to use it

## Installation

`go get github.com/MarioLegenda/execman`

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
ran by **docker exec**. It should be writeable and readable. It can be anywhere you want, but it is
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

# About workers and containers

Choosing the right number of workers and containers is a delicate process. From my experiments,
it is best to have around 20 workers per 1 container. You can find the simulations that I did in 
the `simulations` directory. Success rate of running large amounts of code with it are very good.
With 10 workers, a single container, you can run 1000 concurrent code runs and they will all succeed
in around 15 seconds on my computer. 

Also, when I was working on my blogging platform project and this package was part of a microservice that
executed code and showed it to the user on the frontend, I could use a machine instance with 1 CPU and 
1GB of memory, and it worked just fine. Granted, I didn't have much traffic but memory was stable since
I didn't have to boot up a container every time someone tries to run some code on it. 

Another thing to note is that it rarely fails no matter how many jobs I send to it. It might take a long
time to send 10_000 code runs (or requests, however you want to call it), they will all be a success.

As I said, if you need this project, clone this repository and experiment in the _simulations_ directory
for best results that match your use case. 

As far as memory is concerned, every container, depending on the image that it was created from,
takes between 30 and 50MB.

# Container healing

If a container goes from a _running_ state to some other states, the system will clean it up (stop it and remove it)
and create another one in its place. For example, let's say that you created 10 containers and one of them fails. The system
will pick that up and replace that with another container without the user even knowing that something bad happened.

# Timeouts

If the code does not finish in 5 seconds, it times out. This is to protect the container from
an infinite loop which would render the container useless since code will just keep coming
but old code would not finish. 

# Available programming languages

In code that I showed you, you might have seen this:

````go
_ = instance.Run(execman.RubyLang, `puts "Hello world"`)
````

The `execman.RubyLang` is the constant for the Ruby language. Every programming language
has a constant associated with it. These are the constants and also the available programming
languages that you can run as of this moment:

````go
const Golang = "go"
const NodeLatestLang = "node_latest"
const PerlLtsLang = "perl"
const NodeEsmLtsLang = "node_latest_esm"
const Python2Lang = "python2"
const Python3Lang = "python3"
const LuaLang = "lua"
const RubyLang = "ruby"
const PHP74Lang = "php74"
const RustLang = "rust"
const HaskellLang = "haskell"
const C = "c"
const CPlusPlusLang = "c++"
const CSharpLang = "c_sharp_mono"
const JuliaLang = "julia"
const JavaLang = "java"
````

# Tests

If you want to run tests on this package, cd into the directory where you cloned the package and
run:

`go test -race`