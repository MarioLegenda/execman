package single

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

type RustSingleFileBuildResult struct {
	ContainerDirectory string
	ExecutionDirectory string
	FileName           string
}

type RustSingleFileBuildParams struct {
	Extension string
	Text      string
	StateDir  string
}

func InitRustParams(ext string, text string, stateDir string) RustSingleFileBuildParams {
	return RustSingleFileBuildParams{
		Extension: ext,
		Text:      text,
		StateDir:  stateDir,
	}
}

func RustSingleFileBuild(params RustSingleFileBuildParams) (RustSingleFileBuildResult, error) {
	dirName := uuid.New().String()
	tempExecutionDir := fmt.Sprintf("%s/%s", params.StateDir, dirName)
	fileName := "main.rs"

	if err := os.MkdirAll(tempExecutionDir, os.ModePerm); err != nil {
		return RustSingleFileBuildResult{}, fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create execution dir: %s", err.Error()))
	}

	if err := writeContent(fileName, tempExecutionDir, params.Text); err != nil {
		return RustSingleFileBuildResult{}, err
	}

	if err := writeContent("output.txt", tempExecutionDir, ""); err != nil {
		return RustSingleFileBuildResult{}, err
	}

	if err := writeContent("Cargo.toml", tempExecutionDir, fmt.Sprintf(`[package]
name = "name"
version = "0.0.1"
authors = ["No name"]

[[bin]]
name = "%s"
path = "main.rs"
`, dirName)); err != nil {
		return RustSingleFileBuildResult{}, err
	}

	return RustSingleFileBuildResult{
		ContainerDirectory: dirName,
		ExecutionDirectory: tempExecutionDir,
		FileName:           fileName,
	}, nil
}
