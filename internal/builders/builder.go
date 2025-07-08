package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type BuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type BuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitBuildParams(ext string, text string, stateDir string) BuildParams {
	return BuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func Build(params BuildParams) (BuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return BuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return BuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return BuildResult{}, err
	}

	return BuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
