package builders

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type NodeSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type NodeSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitNodeParams(ext string, text string, stateDir string) NodeSingleFileBuildParams {
	return NodeSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func NodeSingleFileBuild(params NodeSingleFileBuildParams) (NodeSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return NodeSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return NodeSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return NodeSingleFileBuildResult{}, err
	}

	return NodeSingleFileBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", dirName),
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
