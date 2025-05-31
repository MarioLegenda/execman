package linked

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type CPlusProjectBuildResult struct {
	BinaryFileName      string
	ExecutionDirectory  string
	ContainerDirectory  string
	CompilationFileName string
}

type CPlusProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	Text               string
}

func InitCPlusParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, text string) CPlusProjectBuildParams {
	return CPlusProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		Text:               text,
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

	fileName := fmt.Sprintf("%s.%s", execDirConstant, params.CodeProject.Environment.Extension)
	if err := writeContent(fileName, executionDir, params.Text); err != nil {
		return CPlusProjectBuildResult{}, err
	}

	return CPlusProjectBuildResult{
		BinaryFileName:      params.CodeProject.Uuid,
		ExecutionDirectory:  executionDir,
		CompilationFileName: fileName,
		ContainerDirectory:  execDirConstant,
	}, nil
}
