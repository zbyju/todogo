package contents

import "strings"

type Annotation struct {
	Type string
	Text string

	Context string
}

func (a Annotation) String() string {
	var sb strings.Builder
	sb.WriteString(a.Type)
	if a.Context != "" {
		sb.WriteString("(")
		sb.WriteString(a.Context)
		sb.WriteString(")")
	}
	sb.WriteString(": ")
	sb.WriteString(a.Text)
	return sb.String()
}
