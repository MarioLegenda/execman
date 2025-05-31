package single

import (
	"emulator/pkg/appErrors"
	"fmt"
	"os"
)

func writeContent(name string, dir string, content string) *appErrors.Error {
	handle, cErr := os.Create(fmt.Sprintf("%s/%s", dir, name))
	if cErr != nil {
		return appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot create file: %s", cErr.Error()))
	}

	_, err := handle.WriteString(content)

	if err != nil {
		if err := handle.Close(); err != nil {
			return appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot close a file after trying to write to it: %s", err.Error()))
		}

		return appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot write to file: %s", err.Error()))
	}

	err = handle.Close()
	if err != nil {
		return appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot close a file: %s", err.Error()))
	}

	return nil
}
