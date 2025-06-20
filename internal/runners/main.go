package runners

import (
	"fmt"
	"github.com/MarioLegenda/execman/internal/builders"
)

type Params struct {
	ExecutionDir string

	BuilderType   string
	ExecutionType string

	Timeout int

	ContainerName string

	EmulatorName      string
	EmulatorExtension string
	EmulatorText      string
	PackageName       string
}

func Run(params Params) Result {
	if params.EmulatorName == string(nodeLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.NodeSingleFileBuild(builders.InitNodeParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(perlLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.PerlSingleFileBuild(builders.InitPerlParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(luaLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.LuaSingleFileBuild(builders.InitLuaParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(nodeEsm.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.NodeSingleFileBuild(builders.InitNodeParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(goLang.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.GoSingleFileBuild(builders.InitGoParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(ruby.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.RubySingleFileBuild(builders.InitRubyParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(php.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.PhpSingleFileBuild(builders.InitPhpParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(python2.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.PythonSingleFileBuild(builders.InitPythonParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(python3.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.PythonSingleFileBuild(builders.InitPythonParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(csharpMono.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.CsharpSingleFileBuild(builders.InitCsharpParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(haskell.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.HaskellSingleFileBuild(builders.InitHaskellParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(cLang.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.CSingleFileBuild(builders.InitCParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(cPlus.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.CPlusSingleFileBuild(builders.InitCPlusParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(rust.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.RustSingleFileBuild(builders.InitRustParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(julia.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.JuliaSingleFileBuild(builders.InitJuliaParams(
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
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(swift.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.SwiftSingleFileBuild(builders.InitSwiftParams(
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

		return swiftRunner(SwiftExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(kotlin.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.KotlinSingleFileBuild(builders.InitKotlinParams(
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

		return kotlinRunner(KotlinExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(java.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.JavaSingleFileBuild(builders.InitJavaParams(
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

		return javaRunner(JavaExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		})
	}

	if params.EmulatorName == string(zigLts.name) && params.BuilderType == "single_file" && params.ExecutionType == "single_file" {
		build, err := builders.ZigSingleFileBuild(builders.InitZigParams(
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

		return zigRunner(ZigExecParams{
			ExecutionDirectory: build.ExecutionDirectory,
			ContainerDirectory: build.ContainerDirectory,
			ExecutionFile:      build.FileName,
			ContainerName:      params.ContainerName,
			Timeout:            params.Timeout,
		})
	}

	return Result{}
}
