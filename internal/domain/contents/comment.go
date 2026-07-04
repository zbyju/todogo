package contents

import "strings"

type CommentType int

const (
	CNone CommentType = iota
	CSingle
	CMulti
)

type Comment struct {
	CType CommentType
	Text  []string
}

/* Creates a comment based on a single string representing the whole content of the comment.
 *
 * Example (C-like language, single):
 *   - Original file contains: "// TODO(CODE-1234): this is a comment"
 *   - Expected input to this function: "TODO(CODE-1234): this is a comment"
 *   - Output of this function: Comment{CType: Single, Text: ["this is a comment"], Context: "CODE-1234"}
 *
 * Example (C-like language, multi):
 *   - Original file contains: "/* TODO(CODE-1234): this is a comment\n* Another line * /"
 *   - Expected input to this function: "TODO(CODE-1234): this is a comment\nAnother line"
 *   - Output of this function: Comment{CType: Multi, Text: ["this is a comment", "Another line"], Context: "CODE-1234"}
 *
 * Example (Python-like, single):
 *   - Original file contains: "# TODO(CODE-1234): this is a comment"
 *   - Expected input to this function: "TODO(CODE-1234): this is a comment"
 *   - Output of this function: Comment{CType: Multi, Text: ["this is a comment"], Context: "CODE-1234"}
 */
func NewComment(comment string) Comment {
	lines := splitAndRemoveLastLine(comment)

	if len(lines) == 0 || lines[0] == "" {
		return Comment{
			CType: CNone,
		}
	}

	commentType := CSingle

	if len(lines) > 1 {
		commentType = CMulti
	}

	return Comment{
		CType: commentType,
		Text:  lines,
	}
}

func (c Comment) String() string {
	if c.CType == CNone {
		return "None"
	}

	var sb strings.Builder
	if c.CType == CSingle {
		sb.WriteString("Single: ")
	} else {
		sb.WriteString("Multi: ")
	}
	sb.WriteString(strings.Join(c.Text, "\n"))
	return sb.String()
}
