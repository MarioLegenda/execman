package runners

import (
	"fmt"
	single2 "github.com/MarioLegenda/execman/balancer/builders/single"
)

type Params struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	ContainerName string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string
	PackageName       string
}

func Run(params Params) Result {
	if params.EmulatorName == string(nodeLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.NodeSingleFileBuild(single2.InitNodeParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return nodeRunner(NodeExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(perlLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.PerlSingleFileBuild(single2.InitPerlParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return perlRunner(PerlExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(luaLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.LuaSingleFileBuild(single2.InitLuaParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return luaRunner(LuaExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(nodeEsm.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.NodeSingleFileBuild(single2.InitNodeParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return nodeRunner(NodeExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(goLang.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.GoSingleFileBuild(single2.InitGoParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return goRunner(GoExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(ruby.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.RubySingleFileBuild(single2.InitRubyParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return rubyRunner(RubyExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(php.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.PhpSingleFileBuild(single2.InitPhpParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return phpRunner(PhpExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(python2.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.PythonSingleFileBuild(single2.InitPythonParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return pythonRunner(PythonExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(python3.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.PythonSingleFileBuild(single2.InitPythonParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return python3Runner(PythonExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(csharpMono.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.CsharpSingleFileBuild(single2.InitCsharpParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return csharpRunner(CsharpExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(haskell.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.HaskellSingleFileBuild(single2.InitHaskellParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return haskellRunner(HaskellExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(cLang.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.CSingleFileBuild(single2.InitCParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return cRunner(CExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(cPlus.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.CPlusSingleFileBuild(single2.InitCPlusParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return cplusRunner(CPlusExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(rust.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.RustSingleFileBuild(single2.InitRustParams(
			params.EmulatorExtension,
			params.EmulatorText,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return rustRunner(RustExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	return Result{}
}
