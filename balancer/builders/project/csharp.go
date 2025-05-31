package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type CsharpProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type CsharpProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	ExecutingFile      *repository.File
}

func InitCsharpProjectParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, executingFile *repository.File) CsharpProjectBuildParams {
	return CsharpProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		ExecutingFile:      executingFile,
	}
}

func CsharpProjectFileBuild(params CsharpProjectBuildParams) (CsharpProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return CsharpProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return CsharpProjectBuildResult{}, nil
	}

	return CsharpProjectBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", execDirConstant),
		ExecutionDirectory: executionDir,
	}, nil
}
