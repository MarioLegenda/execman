package single

import (
	"emulator/pkg/appErrors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type PhpSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type PhpSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitPhpParams(ext string, text string, stateDir string) PhpSingleFileBuildParams {
	return PhpSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func PhpSingleFileBuild(params PhpSingleFileBuildParams) (PhpSingleFileBuildResult, *appErrors.Error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return PhpSingleFileBuildResult{}, appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return PhpSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return PhpSingleFileBuildResult{}, err
	}

	return PhpSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
