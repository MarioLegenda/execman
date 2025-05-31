package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type CPlusProjectBuildResult struct {
	BinaryFileName      string
	ResolvedFiles       string
	ExecutionDirectory  string
	ContainerDirectory  string
	CompilationFileName string
}

type CPlusProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
}

func InitCPlusParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string) CPlusProjectBuildParams {
	return CPlusProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
	}
}

func CPlusProjectBuild(params CPlusProjectBuildParams) (CPlusProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()
	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return CPlusProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return CPlusProjectBuildResult{}, nil
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

	return CPlusProjectBuildResult{
		BinaryFileName:      params.CodeProject.Uuid,
		ResolvedFiles:       resolvedFiles,
		ExecutionDirectory:  executionDir,
		ContainerDirectory:  execDirConstant,
		CompilationFileName: "main.cpp",
	}, nil
}
