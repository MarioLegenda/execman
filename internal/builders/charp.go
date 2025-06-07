package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type CsharpSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type CsharpSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitCsharpParams(ext string, text string, stateDir string) CsharpSingleFileBuildParams {
	return CsharpSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func CsharpSingleFileBuild(params CsharpSingleFileBuildParams) (CsharpSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return CsharpSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return CsharpSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return CsharpSingleFileBuildResult{}, err
	}

	return CsharpSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
