package linked

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
)

type RubyProjectBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type RubyProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	Text               string
}

func InitRubyParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, text string) RubyProjectBuildParams {
	return RubyProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		Text:               text,
	}
}

func RubyProjectBuild(params RubyProjectBuildParams) (RubyProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()

	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return RubyProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return RubyProjectBuildResult{}, nil
	}

	fileName := fmt.Sprintf("%s.%s", execDirConstant, params.CodeProject.Environment.Extension)
	if err := writeContent(fileName, executionDir, params.Text); err != nil {
		return RubyProjectBuildResult{}, nil
	}

	return RubyProjectBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", execDirConstant),
		ExecutionDirectory: executionDir,
		FileName:           fileName,
	}, nil
}
