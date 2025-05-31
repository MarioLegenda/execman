package single

import (
	"emulator/pkg/appErrors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type RubySingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type RubySingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitRubyParams(ext string, text string, stateDir string) RubySingleFileBuildParams {
	return RubySingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func RubySingleFileBuild(params RubySingleFileBuildParams) (RubySingleFileBuildResult, *appErrors.Error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return RubySingleFileBuildResult{}, appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return RubySingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return RubySingleFileBuildResult{}, err
	}

	return RubySingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
