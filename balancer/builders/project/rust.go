package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type RustProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type RustProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
}

func InitRustParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string) RustProjectBuildParams {
	return RustProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
	}
}

func RustProjectBuild(params RustProjectBuildParams) (RustProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return RustProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return RustProjectBuildResult{}, nil
	}

	if err := writeContent("Cargo.toml", executionDir, fmt.Sprintf(`[package]
name = "name"
version = "0.0.1"
authors = ["No name"]

[[bin]]
name = "%s"
path = "main.rs"
`, execDirConstant)); err != nil {
		return RustProjectBuildResult{}, err
	}

	return RustProjectBuildResult{
		ContainerDirectory: execDirConstant,
		ExecutionDirectory: executionDir,
	}, nil
}
