package ignoring

import "strings"

type Rule struct {
	Path  string
	Value string
}

func (rule Rule) isRuledOutSingle(fullpath string) bool {
	if !strings.HasPrefix(fullpath, rule.Path) {
		return false
	}

	return strings.Contains(fullpath, rule.Value)
}

func IsRuledOut(rules []Rule, fullpath string) bool {
	for _, rule := range rules {
		if rule.isRuledOutSingle(fullpath) {
			return true
		}
	}

	return false
}

func NewRule(path string, line string) *Rule {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}

	rule := Rule{
		Path:  path,
		Value: line,
	}

	return &rule
}
