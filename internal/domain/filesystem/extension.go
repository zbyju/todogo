package filesystem

import (
	"fmt"
	"slices"
	"strings"
)

type Extension [4]byte

var placeholderExtension = [4]byte{'T', 'O', 'D', 'O'}

func NewExtension(extension string) Extension {
	if len(extension) == 0 {
		return placeholderExtension
	}

	if extension[0] == '.' {
		extension = extension[1:]
	}

	if len(extension) > 4 {
		return placeholderExtension
	}

	lowered := strings.ToLower(extension)
	result := fmt.Sprintf("%-4s", lowered)

	return [4]byte{result[0], result[1], result[2], result[3]}
}

func (e Extension) String() string {
	return string(e[:])
}

var knownExtensions = []Extension{
	{'t', 's', ' ', ' '},
	{'p', 'y', ' ', ' '},
	{'g', 'o', ' ', ' '},
	{'m', 'd', ' ', ' '},
}

func (e Extension) IsKnown() bool {
	return slices.Contains(knownExtensions, e)
}
