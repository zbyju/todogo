package fdiscovery

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zbyju/todogo/internal/domain/filesystem"
)

type FileDiscovery interface {
	Discover(path string) (filesystem.Folder, error)
}

type FileDiscoveryImpl struct {
}

func discoverFolder(path string) (*filesystem.Folder, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.New("Cannot read the directory: " + err.Error())
	}

	files := []filesystem.File{}
	folders := []filesystem.Folder{}
	for _, entry := range entries {
		fullpath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			subfolder, err := discoverFolder(fullpath)

			if err != nil {
				fmt.Println("Cannot discover folder: ", err)
			}

			folders = append(folders, *subfolder)
		} else {
			ext := filepath.Ext(fullpath)
			files = append(files, filesystem.NewFile(fullpath, ext))
		}
	}

	result := filesystem.NewFolder(path, folders, files)
	return &result, nil
}

func (fd FileDiscoveryImpl) Discover(path string) (filesystem.Folder, error) {
	folder, err := discoverFolder(path)

	if err != nil {
		return filesystem.Folder{}, err
	}

	return *folder, nil
}
