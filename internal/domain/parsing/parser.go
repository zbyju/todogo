package parsing

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zbyju/todogo/internal/domain/contents"
	"github.com/zbyju/todogo/internal/domain/filesystem"
	"github.com/zbyju/todogo/internal/domain/parsing/regex"
)

func ParseFolder(folder filesystem.Folder) error {
	for _, file := range folder.Files {
		err := ParseFile(file)
		if err != nil {
			return err
		}
	}

	for _, folder := range folder.Folders {
		err := ParseFolder(folder)
		if err != nil {
			return err
		}
	}

	return nil
}

func ParseFile(file filesystem.File) error {
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

	scanner := bufio.NewScanner(f)

	state := regex.NewState(file.Extension, false)
	annotations := []contents.Annotation{}
	for scanner.Scan() {
		line := scanner.Text()
		newAnnotations, newState := regex.ParseLine(line, state)
		state = newState
		annotations = append(annotations, newAnnotations...)
	}
	fmt.Println(fullpath, annotations)

	if err := scanner.Err(); err != nil {
		return errors.New("Cannot read file: " + fullpath)
	}

	return nil
}
