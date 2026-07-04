package regex

import (
	"strings"
)

type TypescriptLineParser struct{}

func extractContent(s string) string {
	s = strings.TrimLeft(s, "/*")
	s = strings.TrimSpace(s)

	return s
}

func (t TypescriptLineParser) IsComment(line contentLine) comment {
	s := string(line)
	if strings.HasPrefix(s, "//") {
		return comment{Text: extractContent(s), Ctype: single}
	}

	if strings.HasPrefix(s, "/*") {
		return comment{Text: extractContent(s), Ctype: multi}
	}

	return comment{Text: "", Ctype: none}
}
