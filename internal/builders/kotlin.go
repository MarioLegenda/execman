package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type KotlinSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type KotlinSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitKotlinParams(ext string, text string, stateDir string) KotlinSingleFileBuildParams {
	return KotlinSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func KotlinSingleFileBuild(params KotlinSingleFileBuildParams) (KotlinSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return KotlinSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return KotlinSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return KotlinSingleFileBuildResult{}, err
	}

	return KotlinSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
