package linked

import (
	"emulator/pkg/repository"
	"fmt"
)

type fileTraverse struct {
	files    []*repository.File
	rootPath string
}

func initFileTraverse(files []*repository.File, rootPath string) fileTraverse {
	return fileTraverse{files: files, rootPath: rootPath}
}

func (ft fileTraverse) createPaths() map[string][]*repository.File {
	paths := make(map[string][]*repository.File)
	mappedSystem := make(map[string]*repository.File)
	var rootDirectory *repository.File

	for _, file := range ft.files {
		if file.IsRoot {
			rootDirectory = file

			continue
		}

		mappedSystem[file.Uuid] = file
	}

	for _, u := range rootDirectory.Children {
		f := mappedSystem[u]

		if !f.IsMain && *f.Parent == rootDirectory.Uuid && f.IsFile {
			path := fmt.Sprintf("%s", ft.rootPath)

			if paths[path] == nil {
				paths[path] = make([]*repository.File, 0)
			}

			paths[path] = append(paths[path], f)
		}
	}

	for _, u := range rootDirectory.Children {
		f := mappedSystem[u]

		if !f.IsFile {
			ft.recursivelyCreatePaths(mappedSystem, f, paths, ft.rootPath)
		}
	}

	return paths
}

func (ft fileTraverse) recursivelyCreatePaths(mappedPaths map[string]*repository.File, directory *repository.File, paths map[string][]*repository.File, rootPath string) map[string][]*repository.File {
	path := fmt.Sprintf("%s/%s", rootPath, directory.Name)

	if len(directory.Children) != 0 {
		for _, u := range directory.Children {
			f := mappedPaths[u]

			if f.IsFile {
				if paths[path] == nil {
					paths[path] = make([]*repository.File, 0)
				}

				paths[path] = append(paths[path], f)
			} else {
				ft.recursivelyCreatePaths(mappedPaths, f, paths, path)
			}
		}
	}

	return paths
}
