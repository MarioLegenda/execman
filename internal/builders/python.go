package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type PythonSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type PythonSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitPythonParams(ext string, text string, stateDir string) RubySingleFileBuildParams {
	return RubySingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func PythonSingleFileBuild(params RubySingleFileBuildParams) (PythonSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return PythonSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return PythonSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return PythonSingleFileBuildResult{}, err
	}

	return PythonSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
