package regex

import (
	"bufio"
	"errors"
	"strings"

	"github.com/zbyju/todogo/internal/domain/contents"
	"github.com/zbyju/todogo/internal/domain/filesystem"
)

type state struct {
	IsCurrentlyInsideComment bool
}

func NewState(isCurrentlyInsideComment bool) state {
	return state{
		IsCurrentlyInsideComment: isCurrentlyInsideComment,
	}
}

type regexParser struct {
}

func NewRegexParser() regexParser {
	return regexParser{}
}

func (regexParser regexParser) ParseFile(file filesystem.OpenedFile) ([]contents.Annotation, error) {
	scanner := bufio.NewScanner(file.Reader)

	state := NewState(false)
	ts := TypescriptLineParser{}
	annotations := []contents.Annotation{}
	for scanner.Scan() {
		line := scanner.Text()
		newAnnotations, newState := parseLine(line, state, ts)
		state = newState
		annotations = append(annotations, newAnnotations...)
	}

	if err := scanner.Err(); err != nil {
		return annotations, errors.New("Cannot read file: " + file.File.Name)
	}
	return annotations, nil
}

type contentLine string

func newLine(line string) contentLine {
	return contentLine(strings.TrimSpace(line))
}

type lineParser interface {
	IsComment(line contentLine) contents.Comment
}

func parseLine(line string, state state, lp lineParser) ([]contents.Annotation, state) {
	annotations := []contents.Annotation{}

	cl := newLine(line)
	if string(cl) == "" {
		return annotations, state
	}

	comment := lp.IsComment(cl)
	if comment.CType == contents.CNone {
		return annotations, state
	}

	annotation := contents.NewAnnotation(comment)

	if annotation.AType == contents.ANone {
		return annotations, state
	}

	annotations = append(annotations, annotation)
	return annotations, state
}
