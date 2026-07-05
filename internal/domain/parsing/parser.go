package parsing

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zbyju/todogo/internal/domain/contents"
	"github.com/zbyju/todogo/internal/domain/filesystem"
	"github.com/zbyju/todogo/internal/domain/ignoring"
	"github.com/zbyju/todogo/internal/domain/parsing/regex"
)

type AnnotationParser interface {
	ParseComment(file filesystem.OpenedFile) []contents.Annotation
}

func ProcessFolder(folder filesystem.Folder, rules *[]ignoring.Rule) error {
	if ignoring.IsRuledOut(*rules, folder.Fullpath()) {
		return nil
	}

	for _, file := range folder.Files {
		err := ProcessFile(file, rules)
		if err != nil {
			fmt.Println("Cannot process file: ", err)
		}
	}

	for _, folder := range folder.Folders {
		err := ProcessFolder(folder, rules)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessFile(file filesystem.File, rules *[]ignoring.Rule) error {
	if !file.IsRelevant() {
		return nil
	}

	if ignoring.IsRuledOut(*rules, file.Fullpath()) {
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

	if file.IsKnown() {
		parser := regex.NewRegexParser()
		annotations, err := parser.ParseFile(of)

		if err != nil {
			return err
		}

		if len(annotations) == 0 {
			return nil
		}

		fmt.Println(file.ColorString("", false))
		for _, annotation := range annotations {
			fmt.Println("   ", annotation.ColorString())
		}

		return nil
	}

	if file.IsIgnoreFile() {
		parser := ignoring.NewIgnoreParser()
		newRules, err := parser.ParseIgnore(&of)

		if err != nil {
			return err
		}
		*rules = append(*rules, newRules...)
	}

	return nil
}
