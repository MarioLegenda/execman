package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type CPlusSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type CPlusSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitCPlusParams(ext string, text string, stateDir string) CPlusSingleFileBuildParams {
	return CPlusSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func CPlusSingleFileBuild(params CPlusSingleFileBuildParams) (CPlusSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return CPlusSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return CPlusSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return CPlusSingleFileBuildResult{}, err
	}

	return CPlusSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
