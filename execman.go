package execman

import (
	"errors"
	"fmt"
	"github.com/MarioLegenda/execman/internal/balancer"
	"github.com/MarioLegenda/execman/internal/containerFactory"
	"github.com/MarioLegenda/execman/internal/types"
	"log"
	"os"
	"sync"
)

type containerBlueprint struct {
	LangName     string
	WorkerNum    int
	ContainerNum int
	Tag          string
	Timeout      int
}

type Result struct {
	Result  string
	Success bool
	Error   error
}

type Option struct {
	Workers    int
	Containers int
	Timeout    int
}

type NodeLts struct {
	Workers    int
	Containers int
	Timeout    int
}

type KotlinLts struct {
	Workers    int
	Containers int
	Timeout    int
}

type Julia struct {
	Workers    int
	Containers int
	Timeout    int
}

type NodeEsm struct {
	Workers    int
	Containers int
	Timeout    int
}

type Ruby struct {
	Workers    int
	Containers int
	Timeout    int
}

type Rust struct {
	Workers    int
	Containers int
	Timeout    int
}

type CPlus struct {
	Workers    int
	Containers int
	Timeout    int
}

type Haskell struct {
	Workers    int
	Containers int
	Timeout    int
}

type CLang struct {
	Workers    int
	Containers int
	Timeout    int
}

type Perl struct {
	Workers    int
	Containers int
	Timeout    int
}

type CSharp struct {
	Workers    int
	Containers int
	Timeout    int
}

type Python3 struct {
	Workers    int
	Containers int
	Timeout    int
}

type Java struct {
	Workers    int
	Containers int
	Timeout    int
}

type Swift struct {
	Workers    int
	Containers int
	Timeout    int
}

type Lua struct {
	Workers    int
	Containers int
	Timeout    int
}

type Kotlin struct {
	Workers    int
	Containers int
	Timeout    int
}

type Python2 struct {
	Workers    int
	Containers int
	Timeout    int
}

type Php74 struct {
	Workers    int
	Containers int
	Timeout    int
}

type Zig struct {
	Workers    int
	Containers int
	Timeout    int
}

type GoLang struct {
	Workers    int
	Containers int
	Timeout    int
}

type Emulator struct {
	executionDir string
	balancers    map[string]*balancer.Balancer
	done         chan interface{}
	cf           *containerFactory.ContainerFactory
}

type Options struct {
	NodeLts
	Julia
	Java
	NodeEsm
	Ruby
	Rust
	CPlus
	Swift
	Haskell
	Zig
	CLang
	Perl
	CSharp
	Kotlin
	Python3
	Lua
	Python2
	Php74
	GoLang

	ExecutionDirectory string
}

func selectProgrammingLanguage(name string) (types.Language, error) {
	if name == "go" {
		return types.GoLang, nil
	} else if name == "node_latest" {
		return types.NodeLts, nil
	} else if name == "node_latest_esm" {
		return types.NodeEsm, nil
	} else if name == "ruby" {
		return types.Ruby, nil
	} else if name == "julia" {
		return types.Julia, nil
	} else if name == "rust" {
		return types.Rust, nil
	} else if name == "c++" {
		return types.CPlus, nil
	} else if name == "haskell" {
		return types.Haskell, nil
	} else if name == "c" {
		return types.CLang, nil
	} else if name == "perl" {
		return types.PerlLts, nil
	} else if name == "c_sharp_mono" {
		return types.CSharpMono, nil
	} else if name == "python3" {
		return types.Python3, nil
	} else if name == "python2" {
		return types.Python2, nil
	} else if name == "lua" {
		return types.Lua, nil
	} else if name == "php74" {
		return types.Php74, nil
	} else if name == "java" {
		return types.JavaLts, nil
	} else if name == "swift" {
		return types.SwiftLts, nil
	} else if name == "kotlin" {
		return types.KotlinLts, nil
	} else if name == "zig" {
		return types.ZigLts, nil
	}

	return types.Language{}, errors.New(fmt.Sprintf("Cannot find language %s", name))
}

func New(options Options) (Emulator, error) {
	initRequiredDirectories(true, options.ExecutionDirectory)

	e := Emulator{
		executionDir: options.ExecutionDirectory,
		balancers:    make(map[string]*balancer.Balancer),
		done:         make(chan interface{}),
		cf:           containerFactory.New(options.ExecutionDirectory),
	}

	containerBlueprints := []containerBlueprint{
		createBlueprint(NodeLatestLang, "node:node_latest", options.NodeLts.Workers, options.NodeLts.Containers, options.NodeLts.Timeout),
		createBlueprint(JavaLang, "java:java_latest", options.Java.Workers, options.Java.Containers, options.Java.Timeout),
		createBlueprint(JuliaLang, "julia:julia", options.Julia.Workers, options.Julia.Containers, options.Julia.Timeout),
		createBlueprint(KotlinLang, "kotlin:kotlin", options.Kotlin.Workers, options.Kotlin.Containers, options.Kotlin.Timeout),
		createBlueprint(ZigLang, "zig:zig", options.Zig.Workers, options.Zig.Containers, options.Zig.Timeout),
		// something is wrong with the way build files are built since they can't be deleted by the clenaup process
		// createBlueprint(SwiftLang, "swift:latest", options.Swift.Workers, options.Swift.Containers),
		createBlueprint(NodeEsmLtsLang, "node:node_latest_esm", options.NodeEsm.Workers, options.NodeEsm.Containers, options.NodeEsm.Timeout),
		createBlueprint(RubyLang, "ruby:ruby", options.Ruby.Workers, options.Ruby.Containers, options.Ruby.Timeout),
		createBlueprint(RustLang, "rust:rust", options.Rust.Workers, options.Rust.Containers, options.Rust.Timeout),
		createBlueprint(CPlusPlusLang, "c-plus:c-plus", options.CPlus.Workers, options.CPlus.Containers, options.CPlus.Timeout),
		createBlueprint(HaskellLang, "haskell:haskell", options.Haskell.Workers, options.Haskell.Containers, options.Haskell.Timeout),
		createBlueprint(C, "c:c", options.CLang.Workers, options.CLang.Containers, options.CLang.Timeout),
		createBlueprint(PerlLtsLang, "perl:perl", options.Perl.Workers, options.Perl.Containers, options.Perl.Timeout),
		createBlueprint(CSharpLang, "c_sharp_mono:c_sharp_mono", options.CSharp.Workers, options.CSharp.Containers, options.CSharp.Timeout),
		createBlueprint(Python3Lang, "python:python3", options.Python3.Workers, options.Python3.Containers, options.Python3.Timeout),
		createBlueprint(LuaLang, "lua:lua", options.Lua.Workers, options.Lua.Containers, options.Lua.Timeout),
		createBlueprint(Python2Lang, "python:python2", options.Python2.Workers, options.Python2.Containers, options.Python2.Timeout),
		createBlueprint(PHP74Lang, "php:php7.4", options.Php74.Workers, options.Php74.Containers, options.Php74.Timeout),
		createBlueprint(Golang, "go:go_latest", options.GoLang.Workers, options.GoLang.Containers, options.GoLang.Timeout),
	}

	// perform initial validation
	for _, c := range containerBlueprints {
		// error if some options are specified but the system cannot work with those options
		if c.WorkerNum == 0 && c.ContainerNum != 0 {
			return Emulator{}, fmt.Errorf("%w: %s", InvalidOptions, fmt.Sprintf("%s cannot have no workers", c.LangName))
		}

		if c.WorkerNum != 0 && c.ContainerNum == 0 {
			return Emulator{}, fmt.Errorf("%w: %s", InvalidOptions, fmt.Sprintf("%s cannot have no containers", c.LangName))
		}
	}

	wg := sync.WaitGroup{}
	containerErrors := make([]error, 0)
	for _, c := range containerBlueprints {
		// default case, user did not specify this language at all
		if c.WorkerNum == 0 && c.ContainerNum == 0 {
			continue
		}

		wg.Add(1)
		go func(c containerBlueprint) {
			defer wg.Done()

			fmt.Printf("Creating containters for [%s]\n", c.Tag)

			e.cf.CreateContainers(c.Tag, c.ContainerNum)

			containers := e.cf.Containers(c.Tag)
			containerNames := make([]string, len(containers))
			for i, c := range containers {
				containerNames[i] = c.Name
			}

			watchCh := e.cf.Watch(c.Tag, e.done)

			b := balancer.New(c.WorkerNum, containerNames, e.done, watchCh, c.Timeout)
			b.StartWorkers()
			b.Watch()
			e.balancers[c.LangName] = b
		}(c)
	}

	wg.Wait()

	if len(containerErrors) != 0 {
		fmt.Println("Some containers could not run. Below is are the errors of those containers:")
		for _, e := range containerErrors {
			fmt.Println(e.Error())
		}

		// if there are errors with creating some containers, others might
		// already be created. We call Close() here to stop those containers
		// and stop all balancers
		e.Close()

		return e, ContainerCannotBoot
	}

	fmt.Println("execman is ready to be used!")

	return e, nil
}

func (em Emulator) Run(language, content string) Result {
	lang, err := selectProgrammingLanguage(language)
	if err != nil {
		return Result{
			Result:  "",
			Success: false,
			Error:   err,
		}
	}

	b := em.balancers[string(lang.Name)]

	resultCh := make(chan balancer.Result)
	b.AddJob(balancer.Job{
		ExecutionDir:      em.executionDir,
		BuilderType:       "single_file",
		ExecutionType:     "single_file",
		EmulatorName:      string(lang.Name),
		EmulatorExtension: lang.Extension,
		EmulatorText:      content,
		ResultCh:          resultCh,
	})

	res := <-resultCh

	return Result{
		Result:  res.Result,
		Success: res.Success,
		Error:   res.Error,
	}
}

func (em Emulator) Close() {
	close(em.done)

	em.cf.Close()
}

func initRequiredDirectories(output bool, executionDir string) {
	projectsDir := executionDir
	directoriesExist := true
	if _, err := os.Stat(projectsDir); os.IsNotExist(err) {
		directoriesExist = false

		if output {
			fmt.Println("")
			fmt.Println("Creating required directories...")
		}
		fsErr := os.Mkdir(projectsDir, os.ModePerm)

		if fsErr != nil {
			log.Fatalln(fmt.Sprintf("Cannot create %s directory: %s", projectsDir, fsErr.Error()))
		}
	}

	if !directoriesExist {
		rest := []string{
			projectsDir,
		}

		for _, dir := range rest {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				fsErr := os.Mkdir(dir, os.ModePerm)

				if fsErr != nil {
					log.Fatalln(fmt.Sprintf("Cannot create %s directory: %s", dir, fsErr.Error()))
				}
			}
		}
	} else {
		if output {
			fmt.Println("")
			fmt.Println("Required directories already created! Skipping...")
			fmt.Println("")
		}
	}

	if !directoriesExist {
		if output {
			fmt.Println("Required directories created!")
			fmt.Println("")
		}
	}
}

func createBlueprint(name, tag string, workers, containers, timeout int) containerBlueprint {
	if timeout == 0 {
		timeout = 5
	}

	return containerBlueprint{
		LangName:     name,
		WorkerNum:    workers,
		ContainerNum: containers,
		Tag:          tag,
		Timeout:      timeout,
	}
}
