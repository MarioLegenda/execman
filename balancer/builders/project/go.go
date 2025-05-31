package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"strings"
)

type GoProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type GoProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	ExecutingFile      *repository.File
	PackageName        string
}

func InitGoParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, executingFile *repository.File, packageName string) GoProjectBuildParams {
	return GoProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		ExecutingFile:      executingFile,
		PackageName:        packageName,
	}
}

func GoProjectBuild(params GoProjectBuildParams) (GoProjectBuildResult, *appErrors.Error) {
	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, params.PackageName)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return GoProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return GoProjectBuildResult{}, nil
	}

	fileName := params.ExecutingFile.Name

	if params.ExecutingFile.Depth != 1 {
		for path, files := range paths {
			for _, file := range files {
				if file.Uuid == params.ExecutingFile.Uuid {
					s := strings.Split(path, params.PackageName)

					fileName = fmt.Sprintf("/app%s/%s", s[1], params.ExecutingFile.Name)
				}
			}
		}
	}

	return GoProjectBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", params.PackageName),
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
