package regex

import (
	"strings"

	"github.com/zbyju/todogo/internal/domain/contents"
	"github.com/zbyju/todogo/internal/domain/filesystem"
)

type state struct {
	Extension                filesystem.Extension
	IsCurrentlyInsideComment bool
}

func NewState(extension filesystem.Extension, isCurrentlyInsideComment bool) state {
	return state{
		Extension:                extension,
		IsCurrentlyInsideComment: isCurrentlyInsideComment,
	}
}

type contentLine string

func newLine(line string) contentLine {
	return contentLine(strings.TrimSpace(line))
}

type commentType int

const (
	none commentType = iota
	single
	multi
)

type comment struct {
	Text  string
	Ctype commentType
}

type LinePraser interface {
	IsComment(line contentLine) comment
}

func ParseLine(line string, state state) ([]contents.Annotation, state) {
	annotations := []contents.Annotation{}

	if !state.Extension.IsKnown() {
		return annotations, state
	}

	cl := newLine(line)
	if string(cl) == "" {
		return annotations, state
	}

	if state.Extension.String() == "ts" {
		ts := TypescriptLineParser{}
		comment := ts.IsComment(cl)

		if comment.Ctype == none {
			return annotations, state
		}

		if comment.Ctype == single {
			if after, ok := strings.CutPrefix(comment.Text, "TODO"); ok {
				return append(annotations, contents.Annotation{
					Type: "TODO",
					Text: strings.TrimLeft(after, ": "),
				}), state
			}
		}

		if comment.Ctype == multi {
			if after, ok := strings.CutPrefix(comment.Text, "TODO"); ok {
				return append(annotations, contents.Annotation{
					Type: "TODO",
					Text: strings.TrimLeft(after, ": "),
				}), NewState(state.Extension, true)
			}
		}

	}

	return annotations, state
}
