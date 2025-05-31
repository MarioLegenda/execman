package single

import (
	"emulator/pkg/appErrors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type JuliaSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type JuliaSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitJuliaParams(ext string, text string, stateDir string) JuliaSingleFileBuildParams {
	return JuliaSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func JuliaSingleFileBuild(params JuliaSingleFileBuildParams) (JuliaSingleFileBuildResult, *appErrors.Error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return JuliaSingleFileBuildResult{}, appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return JuliaSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return JuliaSingleFileBuildResult{}, err
	}

	return JuliaSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
