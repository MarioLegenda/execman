package single

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type LuaSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type LuaSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitLuaParams(ext string, text string, stateDir string) LuaSingleFileBuildParams {
	return LuaSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func LuaSingleFileBuild(params LuaSingleFileBuildParams) (LuaSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := fmt.Sprintf("%s.%s", dirName, params.Extension)

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return LuaSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return LuaSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return LuaSingleFileBuildResult{}, err
	}

	return LuaSingleFileBuildResult{
		ContainerDirectory: fmt.Sprintf("/app/%s", dirName),
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
