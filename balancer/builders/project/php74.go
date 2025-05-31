package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type Php74ProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type Php74ProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	ExecutingFile      *repository.File
}

func InitPhp74Params(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, executingFile *repository.File) Php74ProjectBuildParams {
	return Php74ProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		ExecutingFile:      executingFile,
	}
}

func Php74ProjectBuild(params Php74ProjectBuildParams) (Php74ProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return Php74ProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return Php74ProjectBuildResult{}, nil
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

	return Php74ProjectBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", execDirConstant),
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
