package execman

import (
	"emulator/pkg/projectExecution"
	"emulator/pkg/singleFileExecution"
	_var "emulator/var"
	"errors"
	"fmt"
	"github.com/MarioLegenda/ellie/types"
	"os"
	"sync"
	"time"
)

type Result struct {
	Result  string
	Success bool
	Error   error
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

type Emulator interface {
	RunJob(language string, content string) Result
	Close()
}

type emulator struct {
	executionDir string
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
	LogDirectory       string
}

func selectProgrammingLanguage(name string) (types.Language, error) {
	if name == "go" {
		return types.GoLang, nil
	} else if name == "node_latest" {
		return types.NodeLts, nil
	} else if name == "node_esm" {
		return types.NodeEsm, nil
	} else if name == "ruby" {
		return types.Ruby, nil
	} else if name == "julia" {
		return types.Julia, nil
	} else if name == "rust" {
		return types.Rust, nil
	} else if name == "cplus" {
		return types.CPlus, nil
	} else if name == "haskell" {
		return types.Haskell, nil
	} else if name == "c" {
		return types.CLang, nil
	} else if name == "perl" {
		return types.PerlLts, nil
	} else if name == "csharp" {
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
			errorHandler.TerminateWithMessage(fmt.Sprintf("Cannot create %s directory: %s", projectsDir, fsErr.Error()))
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
					errorHandler.TerminateWithMessage(fmt.Sprintf("Cannot create %s directory: %s", dir, fsErr.Error()))
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

func initExecutioners(options Options) {
	err := execution2.Init(options.ExecutionDirectory, _var.PROJECT_EXECUTION, []execution2.ContainerBlueprint{
		createBlueprint("NODE_LTS", string(types.NodeLts.Tag), options.NodeLts.Workers, options.NodeLts.Containers),
		createBlueprint("JULIA", string(types.Julia.Tag), options.Julia.Workers, options.Julia.Containers),
		createBlueprint("NODE_ESM", string(types.NodeEsm.Tag), options.NodeEsm.Workers, options.NodeEsm.Containers),
		createBlueprint("RUBY", string(types.Ruby.Tag), options.Ruby.Workers, options.Ruby.Containers),
		createBlueprint("RUST", string(types.Rust.Tag), options.Rust.Workers, options.Rust.Containers),
		createBlueprint("CPLUS", string(types.CPlus.Tag), options.CPlus.Workers, options.CPlus.Containers),
		createBlueprint("HASKELL", string(types.Haskell.Tag), options.Haskell.Workers, options.Haskell.Containers),
		createBlueprint("C", string(types.CLang.Tag), options.CLang.Workers, options.CLang.Containers),
		createBlueprint("PERL", string(types.PerlLts.Tag), options.Perl.Workers, options.Perl.Containers),
		createBlueprint("C_SHARP", string(types.CSharpMono.Tag), options.CSharp.Workers, options.CSharp.Containers),
		createBlueprint("PYTHON3", string(types.Python3.Tag), options.Python3.Workers, options.Python3.Containers),
		createBlueprint("LUA", string(types.Lua.Tag), options.Lua.Workers, options.Lua.Containers),
		createBlueprint("PYTHON2", string(types.Python2.Tag), options.Python2.Workers, options.Python2.Containers),
		createBlueprint("PHP74", string(types.Php74.Tag), options.Php74.Workers, options.Php74.Containers),
		createBlueprint("GO", string(types.GoLang.Tag), options.GoLang.Workers, options.GoLang.Containers),
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("Cannot boot: %s", err.Error()))

		if !execution2.Service(_var.PROJECT_EXECUTION).Closed() {
			execution2.Service(_var.PROJECT_EXECUTION).Close()
		}

		time.Sleep(5 * time.Second)

		if os.Getenv("APP_ENV") == "prod" {
			execution2.FinalCleanup(true)
		}

		errorHandler.TerminateWithMessage("Cannot boot executioner.")
	}
}

func createBlueprint(name, tag string, workers, containers int) execution2.ContainerBlueprint {
	return execution2.ContainerBlueprint{
		WorkerNum:    workers,
		ContainerNum: containers,
		Tag:          tag,
	}
}

func New(options Options) Emulator {
	initRequiredDirectories(true, options.ExecutionDirectory)
	singleFileExecution.InitService()
	projectExecution.InitService()

	initExecutioners(options)

	return emulator{
		executionDir: options.ExecutionDirectory,
	}
}

func (e emulator) RunJob(language, content string) Result {
	lang, err := selectProgrammingLanguage(language)
	if err != nil {
		return Result{
			Result:  "",
			Success: false,
			Error:   err,
		}
	}

	res := execution2.Service(_var.PROJECT_EXECUTION).RunJob(execution2.Job{
		ExecutionDir:      e.executionDir,
		BuilderType:       "single_file",
		ExecutionType:     "single_file",
		EmulatorName:      string(lang.Name),
		EmulatorTag:       string(lang.Tag),
		EmulatorExtension: string(lang.Extension),
		EmulatorText:      content,
	})

	realResult := Result{
		Result:  res.Result,
		Success: res.Success,
		Error:   nil,
	}

	if res.Error != nil {
		realResult.Error = errors.New(res.Error.Error())
	}

	return realResult
}

func (e emulator) Close() {
	wg := sync.WaitGroup{}
	for _, e := range []string{_var.PROJECT_EXECUTION} {
		wg.Add(1)

		go func(name string, wg *sync.WaitGroup) {
			if !execution2.Service(name).Closed() {
				execution2.Service(name).Close()
			}
			wg.Done()
		}(e, &wg)
	}

	wg.Wait()

	time.Sleep(5 * time.Second)
	execution2.FinalCleanup(true)
}
