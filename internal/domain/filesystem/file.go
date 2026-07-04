package filesystem

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/zbyju/todogo/internal/style"
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

func (file File) ColorString(leftPad string, shouldPrintPath bool) string {
	var sb strings.Builder

	sb.WriteString(leftPad)
	if shouldPrintPath && file.Path != "" {
		sb.WriteString(file.Path)
	}

	if file.IsKnown() {
		sb.WriteString(style.Apply(" ", style.Blue, style.Bold))
	} else {
		sb.WriteString(style.Apply(" ", style.Gray))
	}

	sb.WriteString(file.Name)

	return sb.String()
}

func (file File) IsKnown() bool {
	return file.Extension.IsKnown()
}

type OpenedFile struct {
	File   File
	Reader *os.File
}

func NewOpenedFile(file File, reader *os.File) OpenedFile {
	return OpenedFile{
		File:   file,
		Reader: reader,
	}
}
