package builders

import (
	"fmt"
	"os"
)

func writeContent(name string, dir string, content string) error {
	handle, cErr := os.Create(fmt.Sprintf("%s/%s", dir, name))
	if cErr != nil {
		return fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot create file: %s", cErr.Error()))
	}

	_, err := handle.WriteString(content)

	if err != nil {
		if err := handle.Close(); err != nil {
			return fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot close a file after trying to write to it: %s", cErr.Error()))
		}

		return fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot write to file: %s", cErr.Error()))
	}

	err = handle.Close()
	if err != nil {
		return fmt.Errorf("%w: %s", FilesystemError, fmt.Sprintf("Cannot close a file: %s", cErr.Error()))
	}

	return nil
}
