package linked

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type Python3ProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type Python3ProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	Text               string
}

func InitPython3Params(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, text string) Python3ProjectBuildParams {
	return Python3ProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		Text:               text,
	}
}

func Python3ProjectBuild(params Python3ProjectBuildParams) (Python3ProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return Python3ProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return Python3ProjectBuildResult{}, nil
	}

	fileName := fmt.Sprintf("%s.%s", execDirConstant, params.CodeProject.Environment.Extension)
	if err := writeContent(fileName, executionDir, params.Text); err != nil {
		return Python3ProjectBuildResult{}, err
	}

	return Python3ProjectBuildResult{
		ContainerDirectory: execDirConstant,
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
