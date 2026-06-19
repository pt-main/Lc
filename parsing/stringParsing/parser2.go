package stringParsing

import (
	"errors"
	"strings"
)

type Parser2 struct{}

func (p *Parser2) Parse(code string, i ...interface{}) ([]ParsedNode, error) {
	lines := strings.Split(code, "\n")
	result := []ParsedNode{}

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := ""
		if len(parts) > 1 {
			args = parts[1]
		}

		meta := map[string]interface{}{
			"command": command,
			"args":    args,
			"__raw":   rawLine,
		}

		node := ParsedNode{
			Raw:      rawLine,
			Switch:   command,
			Metadata: meta,
		}
		result = append(result, node)
	}

	if len(result) == 0 {
		return nil, errors.New("No valid lines found")
	}
	return addPrevNextNodes(result), nil
}

func (p *Parser2) String() string {
	return "lc/parsing/stringParsing/Parser2"
}
