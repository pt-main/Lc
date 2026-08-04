package stringParsing

import (
	"errors"
	"strings"

	"github.com/pt-main/lc/parsing"
)

// # Parser2
//
// Simple parser with grammar like 'cmd args ...', every line converts into
//
//	ParsedNode{Switch: cmd, Metadata{"value": args}},
//
// where cmd is first wotd in line, args is all after cmd.
type Parser2 struct{}

func (p *Parser2) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, error) {
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
