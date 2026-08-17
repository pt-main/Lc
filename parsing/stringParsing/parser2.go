package stringParsing

import (
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public/errors"
)

// Parser2 is a simple command‑args line parser.
type Parser2 struct{}

// Parse converts each non‑empty line into a ParsedNode.
//
// Err errors.ParsingError:
//   - If no valid lines are found in the input.
//     Meta: EMK(0, "string") – the whole input string.
func (p *Parser2) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
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
		return nil, core.Err(errors.ParsingError, "No valid lines found in input").
			WithMeta(core.EMK(0, "string"), code)
	}
	return addPrevNextNodes(result), nil
}

func (p *Parser2) String() string {
	return "lc/parsing/stringParsing/Parser2"
}
