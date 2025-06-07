package execman

import (
	"errors"
	"fmt"
	"github.com/MarioLegenda/execman/containerFactory"
	"github.com/MarioLegenda/execman/newBalancer"
	"github.com/MarioLegenda/execman/types"
	"log"
	"os"
)

type ContainerBlueprint struct {
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
	balancers    map[string]*newBalancer.Balancer
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

/*
*
Initializes new execman instance
*/
func New(options Options) (Emulator, error) {
	initRequiredDirectories(true, options.ExecutionDirectory)

	e := Emulator{
		executionDir: options.ExecutionDirectory,
		balancers:    make(map[string]*newBalancer.Balancer),
	}

	containerBlueprints := []ContainerBlueprint{
		createBlueprint(NodeLatestLang, string(types.NodeLts.Tag), options.NodeLts.Workers, options.NodeLts.Containers),
		createBlueprint(JuliaLang, string(types.Julia.Tag), options.Julia.Workers, options.Julia.Containers),
		createBlueprint(NodeEsmLtsLang, string(types.NodeEsm.Tag), options.NodeEsm.Workers, options.NodeEsm.Containers),
		createBlueprint(RubyLang, string(types.Ruby.Tag), options.Ruby.Workers, options.Ruby.Containers),
		createBlueprint(RustLang, string(types.Rust.Tag), options.Rust.Workers, options.Rust.Containers),
		createBlueprint(CPlusPlusLang, string(types.CPlus.Tag), options.CPlus.Workers, options.CPlus.Containers),
		createBlueprint(HaskellLang, string(types.Haskell.Tag), options.Haskell.Workers, options.Haskell.Containers),
		createBlueprint(C, string(types.CLang.Tag), options.CLang.Workers, options.CLang.Containers),
		createBlueprint(PerlLtsLang, string(types.PerlLts.Tag), options.Perl.Workers, options.Perl.Containers),
		createBlueprint(CSharpLang, string(types.CSharpMono.Tag), options.CSharp.Workers, options.CSharp.Containers),
		createBlueprint(Python3Lang, string(types.Python3.Tag), options.Python3.Workers, options.Python3.Containers),
		createBlueprint(LuaLang, string(types.Lua.Tag), options.Lua.Workers, options.Lua.Containers),
		createBlueprint(Python2Lang, string(types.Python2.Tag), options.Python2.Workers, options.Python2.Containers),
		createBlueprint(PHPLang, string(types.Php74.Tag), options.Php74.Workers, options.Php74.Containers),
		createBlueprint(Golang, string(types.GoLang.Tag), options.GoLang.Workers, options.GoLang.Containers),
	}

	for _, c := range containerBlueprints {
		// default case, user did not specify this language at all
		if c.WorkerNum == 0 && c.ContainerNum == 0 {
			continue
		}

		fmt.Println("Creating container: ", c.Tag)

		errs := containerFactory.CreateContainers(options.ExecutionDirectory, c.Tag, c.ContainerNum)

		if len(errs) != 0 {
			e.Close()

			log := ""
			for _, err := range errs {
				log += fmt.Sprintf("%s,", err.Error())
			}

			return Emulator{}, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Cannot boot container for tag %s: %s", c.Tag, log))
		}

		containers := containerFactory.Containers(c.Tag)
		containerNames := make([]string, len(containers))
		for i, c := range containers {
			containerNames[i] = c.Name
		}

		balancer := newBalancer.New(c.WorkerNum, containerNames)
		balancer.StartWorkers()
		e.balancers[c.LangName] = balancer
	}

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

	resultCh := make(chan newBalancer.Result)
	em.balancers[string(lang.Name)].AddJob(newBalancer.Job{
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
	containerFactory.Close()

	for _, e := range em.balancers {
		e.Close()
	}
	//FinalCleanup(true)
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

func createBlueprint(name, tag string, workers, containers int) ContainerBlueprint {
	return ContainerBlueprint{
		LangName:     name,
		WorkerNum:    workers,
		ContainerNum: containers,
		Tag:          tag,
	}
}
