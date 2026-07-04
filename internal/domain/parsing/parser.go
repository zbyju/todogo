package parsing

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zbyju/todogo/internal/domain/contents"
	"github.com/zbyju/todogo/internal/domain/filesystem"
	"github.com/zbyju/todogo/internal/domain/parsing/regex"
)

type AnnotationParser interface {
	ParseComment(file filesystem.OpenedFile) []contents.Annotation
}

func ProcessFolder(folder filesystem.Folder) error {
	for _, file := range folder.Files {
		err := ProcessFile(file)
		if err != nil {
			return err
		}
	}

	for _, folder := range folder.Folders {
		err := ProcessFolder(folder)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: hello
func ProcessFile(file filesystem.File) error {
	if !file.IsKnown() {
		return nil
	}

	fullpath := filepath.Join(file.Path, file.Name)
	f, err := os.Open(fullpath)
	if err != nil {
		return errors.New("Cannot open file: " + fullpath)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Cannot close file: ", err)
		}
	}()

	of := filesystem.NewOpenedFile(file, f)

	parser := regex.NewRegexParser()
	annotations, err := parser.ParseFile(of)

	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		return nil
	}

	fmt.Println(file.String("", false))
	for _, annotation := range annotations {
		fmt.Println(" ", annotation.String())
	}

	return nil
}
