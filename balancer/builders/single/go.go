package single

import (
	"emulator/pkg/appErrors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type GoSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type GoSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitGoParams(ext string, text string, stateDir string) GoSingleFileBuildParams {
	return GoSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func GoSingleFileBuild(params GoSingleFileBuildParams) (GoSingleFileBuildResult, *appErrors.Error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return GoSingleFileBuildResult{}, appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return GoSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return GoSingleFileBuildResult{}, err
	}

	return GoSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
