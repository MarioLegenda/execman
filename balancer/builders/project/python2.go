package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
	"strings"
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
	ExecutingFile      *repository.File
}

func InitPython2Params(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, executingFile *repository.File) Python2ProjectBuildParams {
	return Python2ProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		ExecutingFile:      executingFile,
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

	fileName := params.ExecutingFile.Name

	if params.ExecutingFile.Depth != 1 {
		for path, files := range paths {
			for _, file := range files {
				if file.Uuid == params.ExecutingFile.Uuid {
					s := strings.Split(path, execDirConstant)

					fileName = fmt.Sprintf("/app%s/%s", s[1], params.ExecutingFile.Name)
				}
			}
		}
	}

	return Python2ProjectBuildResult{
		ContainerDirectory: execDirConstant,
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
