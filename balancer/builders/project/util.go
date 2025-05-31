package project

import (
	"emulator/pkg/appErrors"
	"emulator/pkg/repository"
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

func createDir(path string) *appErrors.Error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cErr := os.MkdirAll(path, os.ModePerm)
		if cErr != nil {
			return appErrors.New(appErrors.ApplicationError, appErrors.FilesystemError, fmt.Sprintf("Cannot create directory: %s", cErr.Error()))
		}
	}

	return nil
}

func createFsSystem(paths map[string][]*repository.File, contents []*repository.FileContent) *appErrors.Error {
	contentsMap := make(map[string]*repository.FileContent)

	for _, c := range contents {
		contentsMap[c.Uuid] = c
	}

	for path, files := range paths {
		if err := createDir(path); err != nil {
			return err
		}

		if files != nil && len(files) != 0 {
			for _, f := range files {
				if content, ok := contentsMap[f.Uuid]; ok {
					if err := writeContent(f.Name, path, content.Content); err != nil {
						return err
					}

					continue
				}

				content := &repository.FileContent{
					CodeProjectUuid: "",
					Uuid:            "",
					Content:         "",
				}

				if err := writeContent(f.Name, path, content.Content); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
