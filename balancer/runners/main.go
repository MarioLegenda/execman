package runners

import (
	linked2 "emulator/pkg/execution/balancer/builders/linked"
	project2 "emulator/pkg/execution/balancer/builders/project"
	single2 "emulator/pkg/execution/balancer/builders/single"
	"emulator/pkg/repository"
	"fmt"
)

type Params struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	ContainerName string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string

	CodeProject   *repository.CodeProject
	Contents      []*repository.FileContent
	ExecutingFile *repository.File
	PackageName   string
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

	if params.EmulatorName == string(nodeLts.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.NodeProjectBuild(project2.InitNodeParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(luaLts.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.LuaProjectBuild(project2.InitLuaParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(perlLts.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.PerlProjectBuild(project2.InitPerlParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(nodeEsm.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.NodeProjectBuild(project2.InitNodeParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(goLang.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.GoProjectBuild(project2.InitGoParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
			params.PackageName,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return goProjectRunner(GoProjectExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: params.PackageName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(ruby.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.RubyProjectBuild(project2.InitRubyParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(cLang.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.CProjectBuild(project2.InitCParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return cProjectRunner(CProjectExecParams{
			ContainerName:      params.ContainerName,
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ResolvedPaths:      build.ResolvedFiles,
			BinaryFileName:     build.BinaryFileName,
		})
	}

	if params.EmulatorName == string(cPlus.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.CPlusProjectBuild(project2.InitCPlusParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return cPlusProjectRunner(CPlusProjectExecParams{
			ContainerName:       params.ContainerName,
			ExecutionDirectory:  build.ExecutionDirectory,
			ContainerDirectory:  build.ContainerDirectory,
			BinaryFileName:      build.BinaryFileName,
			CompilationFileName: build.CompilationFileName,
		})
	}

	if params.EmulatorName == string(csharpMono.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.CsharpProjectFileBuild(project2.InitCsharpProjectParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(python2.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.Python2ProjectBuild(project2.InitPython2Params(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(python3.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.Python3ProjectBuild(project2.InitPython3Params(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(haskell.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.HaskellProjectBuild(project2.InitHaskellProjectParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return haskellProjectRunner(HaskellExecProjectParams{
			ExecutionDirectory:  build.ExecutionDirectory,
			ContainerDirectory:  build.ContainerDirectory,
			ContainerName:       params.ContainerName,
			CompilationFileName: build.CompilationFileName,
		})
	}

	if params.EmulatorName == string(rust.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.RustProjectBuild(project2.InitRustParams(
			params.CodeProject,
			params.Contents,
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
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(php.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.Php74ProjectBuild(project2.InitPhp74Params(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
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

	if params.EmulatorName == string(ruby.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.RubyProjectBuild(linked2.InitRubyParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(luaLts.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.LuaProjectBuild(linked2.InitLuaParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(perlLts.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.PerlProjectBuild(linked2.InitPerlParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(nodeLts.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.NodeProjectBuild(linked2.InitNodeParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(nodeEsm.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.NodeProjectBuild(linked2.InitNodeParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(goLang.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.GoProjectBuild(linked2.InitGoParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
			params.PackageName,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return goProjectRunner(GoProjectExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(haskell.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.HaskellProjectBuild(linked2.InitHaskellProjectParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return haskellProjectRunner(HaskellExecProjectParams{
			ExecutionDirectory:  build.ExecutionDirectory,
			ContainerDirectory:  build.ContainerDirectory,
			ContainerName:       params.ContainerName,
			CompilationFileName: build.CompilationFileName,
		})
	}

	if params.EmulatorName == string(rust.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.RustProjectBuild(linked2.InitRustParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(cLang.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.CProjectBuild(linked2.InitCParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return cProjectRunner(CProjectExecParams{
			ContainerName:      params.ContainerName,
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ResolvedPaths:      build.ResolvedFiles,
			BinaryFileName:     build.BinaryFileName,
		})
	}

	if params.EmulatorName == string(cPlus.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.CPlusProjectBuild(linked2.InitCPlusParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return cPlusProjectRunner(CPlusProjectExecParams{
			ContainerName:       params.ContainerName,
			ExecutionDirectory:  build.ExecutionDirectory,
			ContainerDirectory:  build.ContainerDirectory,
			BinaryFileName:      build.BinaryFileName,
			CompilationFileName: build.CompilationFileName,
		})
	}

	if params.EmulatorName == string(csharpMono.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.CsharpProjectFileBuild(linked2.InitCsharpProjectParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(php.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.Php74ProjectBuild(linked2.InitPhp74Params(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(python2.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.Python2ProjectBuild(linked2.InitPython2Params(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(python3.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.Python3ProjectBuild(linked2.InitPython3Params(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
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

	if params.EmulatorName == string(julia.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := single2.JuliaSingleFileBuild(single2.InitJuliaParams(
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

		return juliaRunner(JuliaExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(julia.name) && params.BuilderType == "project" && params.ExecutionType == "project" {
		build, err := project2.JuliaProjectBuild(project2.InitJuliaParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.ExecutingFile,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return juliaRunner(JuliaExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	if params.EmulatorName == string(julia.name) && params.BuilderType == "linked" && params.ExecutionType == "linked" {
		build, err := linked2.JuliaProjectBuild(linked2.InitJuliaParams(
			params.CodeProject,
			params.Contents,
			fmt.Sprintf("%s/%s", params.ExecutionDir, params.ContainerName),
			params.EmulatorText,
		))

		if err != nil {
			return Result{
				Result:  "",
				Success: false,
				Error:   err,
			}
		}

		return juliaRunner(JuliaExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
		})
	}

	return Result{}
}
