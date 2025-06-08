package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type SwiftSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type SwiftSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitSwiftParams(ext string, text string, stateDir string) SwiftSingleFileBuildParams {
	return SwiftSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func SwiftSingleFileBuild(params SwiftSingleFileBuildParams) (SwiftSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return SwiftSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return SwiftSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return SwiftSingleFileBuildResult{}, err
	}

	return SwiftSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
