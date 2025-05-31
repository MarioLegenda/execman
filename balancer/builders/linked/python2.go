package linked

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type Python2ProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type Python2ProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	Text               string
}

func InitPython2Params(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, text string) Python2ProjectBuildParams {
	return Python2ProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		Text:               text,
	}
}

func Python2ProjectBuild(params Python2ProjectBuildParams) (Python2ProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return Python2ProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return Python2ProjectBuildResult{}, nil
	}

	fileName := fmt.Sprintf("%s.%s", execDirConstant, params.CodeProject.Environment.Extension)
	if err := writeContent(fileName, executionDir, params.Text); err != nil {
		return Python2ProjectBuildResult{}, err
	}

	return Python2ProjectBuildResult{
		ContainerDirectory: execDirConstant,
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
