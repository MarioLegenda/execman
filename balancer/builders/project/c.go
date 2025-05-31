package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type CProjectBuildResult struct {
	BinaryFileName     string
	ResolvedFiles      string
	ExecutionDirectory string
	ContainerDirectory string
}

type CProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
}

func InitCParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string) CProjectBuildParams {
	return CProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
	}
}

func CProjectBuild(params CProjectBuildParams) (CProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()
	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return CProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return CProjectBuildResult{}, nil
	}

	resolvedFiles := ""
	for dir, files := range paths {
		s := strings.Split(dir, execDirConstant)
		dockerPath := s[1]

		for _, file := range files {
			if dockerPath == "" {
				resolvedFiles += fmt.Sprintf("%s ", file.Name)
			} else {
				resolvedFiles += fmt.Sprintf("%s/%s ", dockerPath, file.Name)
			}
		}
	}

	return CProjectBuildResult{
		BinaryFileName:     params.CodeProject.Uuid,
		ResolvedFiles:      resolvedFiles,
		ExecutionDirectory: executionDir,
		ContainerDirectory: execDirConstant,
	}, nil
}
