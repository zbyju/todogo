package filesystem

import (
	"path/filepath"
	"sort"
	"strings"
)

type Folder struct {
	Path      string
	Name      string
	Folders   []Folder
	Files     []File
	IsIgnored bool
}

const FolderDelimiter = "/"

func NewFolder(fullpath string, folders []Folder, files []File, isIgnored bool) Folder {
	path, name := filepath.Split(fullpath)

	sort.Slice(folders, func(i, j int) bool {
		return folders[i].Name < folders[j].Name
	})

	// Sort files by Name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return Folder{
		Path:      path,
		Name:      name,
		Folders:   folders,
		Files:     files,
		IsIgnored: isIgnored,
	}
}

func (folder Folder) IsEmpty() bool {
	return len(folder.Folders) == 0 && len(folder.Files) == 0
}

func (folder Folder) String(leftPad string, leftPadStr string, shouldPrintPath bool) string {
	var sb strings.Builder

	sb.WriteString(leftPad)
	if shouldPrintPath && folder.Path != "" {
		sb.WriteString(folder.Path)
	}

	if !folder.IsEmpty() {
		sb.WriteString(" ")
	} else {
		sb.WriteString(" ")
	}

	sb.WriteString(folder.Name)
	sb.WriteString(FolderDelimiter)
	sb.WriteString("\n")

	newLeftPad := leftPad + leftPadStr
	for _, subfolder := range folder.Folders {
		sb.WriteString(subfolder.String(newLeftPad, leftPadStr, shouldPrintPath))
	}

	for _, file := range folder.Files {
		sb.WriteString(file.String(newLeftPad, false))
	}

	return sb.String()
}
