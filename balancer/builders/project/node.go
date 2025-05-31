package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type NodeProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type NodeProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	ExecutingFile      *repository.File
}

func InitNodeParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, executingFile *repository.File) NodeProjectBuildParams {
	return NodeProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		ExecutingFile:      executingFile,
	}
}

func NodeProjectBuild(params NodeProjectBuildParams) (NodeProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return NodeProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return NodeProjectBuildResult{}, nil
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

	return NodeProjectBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", execDirConstant),
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
