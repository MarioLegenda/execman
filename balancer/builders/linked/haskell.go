package linked

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type HaskellProjectBuildResult struct {
	BinaryFileName      string
	ResolvedFiles       string
	ExecutionDirectory  string
	ContainerDirectory  string
	CompilationFileName string
}

type HaskellProjectBuildParams struct {
	CodeProject        *repository.CodeProject
	Contents           []*repository.FileContent
	ContainerDirectory string
	Text               string
}

func InitHaskellProjectParams(cp *repository.CodeProject, contents []*repository.FileContent, containerDir string, text string) HaskellProjectBuildParams {
	return HaskellProjectBuildParams{
		CodeProject:        cp,
		Contents:           contents,
		ContainerDirectory: containerDir,
		Text:               text,
	}
}

func HaskellProjectBuild(params HaskellProjectBuildParams) (HaskellProjectBuildResult, *appErrors.Error) {
	execDirConstant := uuid.New().String()
	executionDir := fmt.Sprintf("%s/%s", params.ContainerDirectory, execDirConstant)
	ft := initFileTraverse(params.CodeProject.Structure, executionDir)

	paths := ft.createPaths()

	if err := createDir(executionDir); err != nil {
		return HaskellProjectBuildResult{}, err
	}

	if err := createFsSystem(paths, params.Contents); err != nil {
		return HaskellProjectBuildResult{}, nil
	}

	fileName := fmt.Sprintf("%s.%s", execDirConstant, params.CodeProject.Environment.Extension)
	if err := writeContent(fileName, executionDir, params.Text); err != nil {
		return HaskellProjectBuildResult{}, err
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

	return HaskellProjectBuildResult{
		BinaryFileName:      params.CodeProject.Uuid,
		ResolvedFiles:       resolvedFiles,
		ExecutionDirectory:  executionDir,
		ContainerDirectory:  execDirConstant,
		CompilationFileName: fileName,
	}, nil
}
