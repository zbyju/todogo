package filesystem

import (
	"path/filepath"
	"strings"
)

type File struct {
	Path      string
	Name      string
	Extension Extension
	IsIgnored bool
}

func NewFile(fullpath string, extension string, isIgnored bool) File {
	path, name := filepath.Split(fullpath)
	ext := NewExtension(extension)
	return File{
		Path:      path,
		Name:      name,
		Extension: ext,
		IsIgnored: isIgnored,
	}
}

func (file File) String(leftPad string, shouldPrintPath bool) string {
	var sb strings.Builder

	sb.WriteString(leftPad)
	if shouldPrintPath && file.Path != "" {
		sb.WriteString(file.Path)
	}

	if file.IsKnown() {
		sb.WriteString(" ")
	} else {
		sb.WriteString(" ")
	}

	sb.WriteString(file.Name)
	sb.WriteString("\n")

	return sb.String()
}

func (file File) IsKnown() bool {
	return file.Extension.IsKnown()
}
