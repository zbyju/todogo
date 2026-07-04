package contents

import (
	"strings"
	"unicode"

	"github.com/zbyju/todogo/internal/style"
)

type AnnotationType int

const (
	ANone AnnotationType = iota
	ASingle
	AMulti
)

type Annotation struct {
	AType AnnotationType
	AText string
	Text  []string

	Context string
}

func processFirstLine(firstline string) Annotation {
	processedAtext := false
	processingContext := false

	var atext strings.Builder
	hasAText := false
	hasContext := false
	var context strings.Builder
	var rest strings.Builder
	for _, c := range firstline {
		// 1. Try to process the annotation type
		if !processedAtext {
			if unicode.IsUpper(c) {
				atext.WriteString(string(c))
			} else if c == ':' {
				hasAText = true
				processedAtext = true
				processingContext = false
				hasContext = false
			} else if c == '(' {
				processingContext = true
			} else {
				if atext.String() == "" {
					return Annotation{AType: ANone}
				}
				return Annotation{AType: ANone, AText: "", Text: []string{atext.String() + context.String() + rest.String()}, Context: ""}
			}
			continue
		}

		// 2. Should we attempt to process the context?
		if processingContext {
			if c == ')' {
				hasContext = true
			} else if hasContext && c == ':' {
				hasAText = true
				processingContext = false
			} else {
				context.WriteString(string(c))
			}
			continue
		}

		// 3. Finish the rest
		rest.WriteString(string(c))
	}

	if !hasAText {
		return Annotation{AType: ANone, AText: "", Text: []string{atext.String() + context.String() + rest.String()}, Context: ""}
	}

	if !hasContext {
		return Annotation{AType: ASingle, AText: atext.String(), Text: []string{context.String() + rest.String()}, Context: ""}
	}

	return Annotation{AType: ASingle, AText: atext.String(), Text: []string{rest.String()}, Context: context.String()}
}

func NewAnnotation(comment Comment) Annotation {
	if comment.CType == CNone || len(comment.Text) == 0 {
		return Annotation{AType: ANone}
	}

	first := processFirstLine(comment.Text[0])
	if first.AType == ANone {
		return Annotation{AType: ANone}
	}

	atype := ASingle
	if len(comment.Text) > 1 {
		atype = AMulti
	}

	text := []string{}
	if len(first.Text) != 0 {
		text = append(text, first.Text[0])
		for _, t := range comment.Text[1:] {
			text = append(text, strings.TrimSpace(t))
		}
	} else {
		for _, t := range comment.Text {
			text = append(text, strings.TrimSpace(t))
		}
	}

	return Annotation{
		AType:   atype,
		AText:   strings.TrimSpace(first.AText),
		Text:    text,
		Context: first.Context,
	}
}

func (a Annotation) String() string {
	var sb strings.Builder
	sb.WriteString(a.AText)
	if a.Context != "" {
		sb.WriteString("(")
		sb.WriteString(a.Context)
		sb.WriteString(")")
	}
	sb.WriteString(":")
	sb.WriteString(strings.Join(a.Text, "\n"))
	return sb.String()
}

// annotationColor picks a color by annotation type; unknown types are gray.
func annotationColor(atext string) string {
	switch strings.ToUpper(atext) {
	case "TODO":
		return style.Yellow
	case "FIXME", "BUG", "XXX":
		return style.Red
	case "NOTE", "INFO":
		return style.Blue
	default:
		return style.Gray
	}
}

func (a Annotation) ColorString() string {
	var sb strings.Builder
	sb.WriteString(style.Apply(a.AText, annotationColor(a.AText), style.Bold))
	if a.Context != "" {
		sb.WriteString("(")
		sb.WriteString(a.Context)
		sb.WriteString(")")
	}
	sb.WriteString(":")
	sb.WriteString(strings.Join(a.Text, "\n"))
	return sb.String()
}

func splitAndRemoveLastLine(comment string) []string {
	lines := strings.Split(comment, "\n")

	n := len(lines)
	if n <= 1 {
		return lines
	}

	if strings.TrimSpace(lines[n-1]) == "" {
		return lines[:n-1]
	}

	return lines
}
