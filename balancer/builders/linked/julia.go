package linked

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type JuliaProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type JuliaProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	Text               string
}

func InitJuliaParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, text string) JuliaProjectBuildParams {
	return JuliaProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		Text:               text,
	}
}

func JuliaProjectBuild(params JuliaProjectBuildParams) (JuliaProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return JuliaProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return JuliaProjectBuildResult{}, nil
	}

	fileName := fmt.Sprintf("%s.%s", execDirConstant, params.CodeProject.Environment.Extension)
	if err := writeContent(fileName, executionDir, params.Text); err != nil {
		return JuliaProjectBuildResult{}, nil
	}

	return JuliaProjectBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", execDirConstant),
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
