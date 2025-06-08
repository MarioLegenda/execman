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
}

type Result struct {
	Result  string
	Success bool
	Error   error
}

type Option struct {
	Workers    int
	Containers int
}

type NodeLts struct {
	Workers    int
	Containers int
}

type Julia struct {
	Workers    int
	Containers int
}

type NodeEsm struct {
	Workers    int
	Containers int
}

type Ruby struct {
	Workers    int
	Containers int
}

type Rust struct {
	Workers    int
	Containers int
}

type CPlus struct {
	Workers    int
	Containers int
}

type Haskell struct {
	Workers    int
	Containers int
}

type CLang struct {
	Workers    int
	Containers int
}

type Perl struct {
	Workers    int
	Containers int
}

type CSharp struct {
	Workers    int
	Containers int
}

type Python3 struct {
	Workers    int
	Containers int
}

type Lua struct {
	Workers    int
	Containers int
}

type Python2 struct {
	Workers    int
	Containers int
}

type Php74 struct {
	Workers    int
	Containers int
}

type GoLang struct {
	Workers    int
	Containers int
}

type Emulator struct {
	executionDir string
	balancers    map[string]*balancer.Balancer
}

type Options struct {
	NodeLts
	Julia
	NodeEsm
	Ruby
	Rust
	CPlus
	Haskell
	CLang
	Perl
	CSharp
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
	}

	return types.Language{}, errors.New(fmt.Sprintf("Cannot find language %s", name))
}

func New(options Options) (Emulator, error) {
	initRequiredDirectories(true, options.ExecutionDirectory)

	e := Emulator{
		executionDir: options.ExecutionDirectory,
		balancers:    make(map[string]*balancer.Balancer),
	}

	containerBlueprints := []containerBlueprint{
		createBlueprint(NodeLatestLang, "node:node_latest", options.NodeLts.Workers, options.NodeLts.Containers),
		createBlueprint(JuliaLang, "julia:julia", options.Julia.Workers, options.Julia.Containers),
		createBlueprint(NodeEsmLtsLang, "node:node_latest_esm", options.NodeEsm.Workers, options.NodeEsm.Containers),
		createBlueprint(RubyLang, "ruby:ruby", options.Ruby.Workers, options.Ruby.Containers),
		createBlueprint(RustLang, "rust:rust", options.Rust.Workers, options.Rust.Containers),
		createBlueprint(CPlusPlusLang, "c-plus:c-plus", options.CPlus.Workers, options.CPlus.Containers),
		createBlueprint(HaskellLang, "haskell:haskell", options.Haskell.Workers, options.Haskell.Containers),
		createBlueprint(C, "c:c", options.CLang.Workers, options.CLang.Containers),
		createBlueprint(PerlLtsLang, "perl:perl", options.Perl.Workers, options.Perl.Containers),
		createBlueprint(CSharpLang, "c_sharp_mono:c_sharp_mono", options.CSharp.Workers, options.CSharp.Containers),
		createBlueprint(Python3Lang, "python:python3", options.Python3.Workers, options.Python3.Containers),
		createBlueprint(LuaLang, "lua:lua", options.Lua.Workers, options.Lua.Containers),
		createBlueprint(Python2Lang, "python:python2", options.Python2.Workers, options.Python2.Containers),
		createBlueprint(PHP74Lang, "php:php7.4", options.Php74.Workers, options.Php74.Containers),
		createBlueprint(Golang, "go:go_latest", options.GoLang.Workers, options.GoLang.Containers),
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

			errs := containerFactory.CreateContainers(options.ExecutionDirectory, c.Tag, c.ContainerNum)

			if len(errs) != 0 {
				combinedLogging := ""
				for _, err := range errs {
					combinedLogging += fmt.Sprintf("%s,", err.Error())
				}

				containerErrors = append(errs, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Cannot boot container for tag %s: %s", c.Tag, combinedLogging)))

				return
			}

			containers := containerFactory.Containers(c.Tag)
			containerNames := make([]string, len(containers))
			for i, c := range containers {
				containerNames[i] = c.Name
			}

			b := balancer.New(c.WorkerNum, containerNames)
			b.StartWorkers()
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
	for _, e := range em.balancers {
		e.Close()
	}

	containerFactory.Close()
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

func createBlueprint(name, tag string, workers, containers int) containerBlueprint {
	return containerBlueprint{
		LangName:     name,
		WorkerNum:    workers,
		ContainerNum: containers,
		Tag:          tag,
	}
}
