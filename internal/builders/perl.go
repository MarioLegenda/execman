package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type PerlSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type PerlSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitPerlParams(ext string, text string, stateDir string) PerlSingleFileBuildParams {
	return PerlSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func PerlSingleFileBuild(params PerlSingleFileBuildParams) (PerlSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return PerlSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return PerlSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return PerlSingleFileBuildResult{}, err
	}

	return PerlSingleFileBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", dirName),
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
