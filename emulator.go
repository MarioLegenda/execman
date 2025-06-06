package execman

import (
	"errors"
	"fmt"
	"github.com/MarioLegenda/execman/containerFactory"
	"github.com/MarioLegenda/execman/newBalancer"
	"github.com/MarioLegenda/execman/types"
	"log"
	"os"
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
	execution    *execution
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

	//executioner := initExecutioners(options)

	e := emulator{
		executionDir: options.ExecutionDirectory,
		balancers:    make(map[string]*newBalancer.Balancer),
	}

	containerBlueprints := []ContainerBlueprint{
		createBlueprint(string(types.NodeLts.Name), string(types.NodeLts.Tag), options.NodeLts.Workers, options.NodeLts.Containers),
		createBlueprint(string(types.Julia.Name), string(types.Julia.Tag), options.Julia.Workers, options.Julia.Containers),
		createBlueprint(string(types.NodeEsm.Name), string(types.NodeEsm.Tag), options.NodeEsm.Workers, options.NodeEsm.Containers),
		createBlueprint(string(types.Ruby.Name), string(types.Ruby.Tag), options.Ruby.Workers, options.Ruby.Containers),
		createBlueprint(string(types.Rust.Name), string(types.Rust.Tag), options.Rust.Workers, options.Rust.Containers),
		createBlueprint(string(types.CPlus.Name), string(types.CPlus.Tag), options.CPlus.Workers, options.CPlus.Containers),
		createBlueprint(string(types.Haskell.Name), string(types.Haskell.Tag), options.Haskell.Workers, options.Haskell.Containers),
		createBlueprint(string(types.CLang.Name), string(types.CLang.Tag), options.CLang.Workers, options.CLang.Containers),
		createBlueprint(string(types.PerlLts.Name), string(types.PerlLts.Tag), options.Perl.Workers, options.Perl.Containers),
		createBlueprint(string(types.CSharpMono.Name), string(types.CSharpMono.Tag), options.CSharp.Workers, options.CSharp.Containers),
		createBlueprint(string(types.Python3.Name), string(types.Python3.Tag), options.Python3.Workers, options.Python3.Containers),
		createBlueprint(string(types.Lua.Name), string(types.Lua.Tag), options.Lua.Workers, options.Lua.Containers),
		createBlueprint(string(types.Python2.Name), string(types.Python2.Tag), options.Python2.Workers, options.Python2.Containers),
		createBlueprint(string(types.Php74.Name), string(types.Php74.Tag), options.Php74.Workers, options.Php74.Containers),
		createBlueprint(string(types.GoLang.Name), string(types.GoLang.Tag), options.GoLang.Workers, options.GoLang.Containers),
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

			return emulator{}, fmt.Errorf("%w: %s", ContainerCannotBoot, fmt.Sprintf("Cannot boot container for tag %s: %s", c.Tag, log))
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

func (em emulator) RunJob(language, content string) Result {
	lang, err := selectProgrammingLanguage(language)
	if err != nil {
		return Result{
			Result:  "",
			Success: false,
			Error:   err,
		}
	}

	em.balancers[lang.Language].AddJob(newBalancer.Job{
		ExecutionDir:      em.executionDir,
		BuilderType:       "single_file",
		ExecutionType:     "single_file",
		EmulatorName:      string(lang.Name),
		EmulatorExtension: lang.Extension,
		EmulatorText:      content,
	})

	return Result{}
}

func (em emulator) Close() {
	containerFactory.Close()

	for _, e := range em.balancers {
		e.Close()
	}

	time.Sleep(5 * time.Second)
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

func initExecutioners(options Options) *execution {
	exec, err := Init(options.ExecutionDirectory, []ContainerBlueprint{
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

		// TODO: investigate why is this here?
		time.Sleep(5 * time.Second)

		log.Fatalln("Cannot boot executioners")
	}

	return exec
}

func createBlueprint(name, tag string, workers, containers int) ContainerBlueprint {
	return ContainerBlueprint{
		LangName:     name,
		WorkerNum:    workers,
		ContainerNum: containers,
		Tag:          tag,
	}
}
