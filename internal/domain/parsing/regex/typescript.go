package regex

import (
	"strings"

	"github.com/zbyju/todogo/internal/domain/contents"
)

type TypescriptLineParser struct{}

func extractContent(s string) string {
	s = strings.TrimLeft(s, "/*")
	s = strings.TrimSpace(s)

	return s
}

func (t TypescriptLineParser) IsComment(line contentLine) contents.Comment {
	s := string(line)
	if strings.HasPrefix(s, "//") {
		return contents.NewComment(extractContent(s))
	}

	if strings.HasPrefix(s, "/*") {
		return contents.NewComment(extractContent(s))
	}

	return contents.NewComment("")
}
