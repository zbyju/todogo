package style

import "strings"

const reset = "\033[0m"

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"
)

const (
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
)

// Apply wraps text in the given codes and appends a single reset.
// Pass any mix of colors and styles, e.g. Apply("hi", Red, Bold).
func Apply(text string, codes ...string) string {
	if len(codes) == 0 {
		return text
	}
	var sb strings.Builder
	for _, c := range codes {
		sb.WriteString(c)
	}
	sb.WriteString(text)
	sb.WriteString(reset)
	return sb.String()
}

// Color wraps text in one color and resets.
func Color(text, color string) string { return Apply(text, color) }

func Boldize(text string) string    { return Apply(text, Bold) }
func Dimize(text string) string     { return Apply(text, Dim) }
func Italicize(text string) string  { return Apply(text, Italic) }
func Underlined(text string) string { return Apply(text, Underline) }
