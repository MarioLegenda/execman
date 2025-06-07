package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type HaskellSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type HaskellSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitHaskellParams(ext string, text string, stateDir string) HaskellSingleFileBuildParams {
	return HaskellSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func HaskellSingleFileBuild(params HaskellSingleFileBuildParams) (HaskellSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return HaskellSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return HaskellSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return HaskellSingleFileBuildResult{}, err
	}

	return HaskellSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
