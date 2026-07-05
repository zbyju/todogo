package ignoring

import (
	"bufio"

	"github.com/zbyju/todogo/internal/domain/filesystem"
)

type IgnoreParser interface {
	ParseIgnore(of *filesystem.OpenedFile) ([]Rule, error)
}

func NewIgnoreParser() IgnoreParser {
	return SimpleIgnoreParser{}
}

type SimpleIgnoreParser struct{}

func (ip SimpleIgnoreParser) ParseIgnore(file *filesystem.OpenedFile) ([]Rule, error) {
	scanner := bufio.NewScanner(file.Reader)

	rules := []Rule{}
	for scanner.Scan() {
		line := scanner.Text()
		if rule := NewRule(file.File.Path, line); rule != nil {
			rules = append(rules, *rule)
		}
	}

	if err := scanner.Err(); err != nil {
		return rules, err
	}

	return rules, nil
}
