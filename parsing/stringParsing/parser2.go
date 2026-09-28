package stringParsing

import (
	"strings"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing"
	"github.com/pt-main/lc/v2/public/errors"
)

// Parser2 is a simple command-args line parser.
type Parser2 struct{}

// Parse converts each non-empty line into a ParsedNode.
//
// Err errors.ParsingError:
//   - If no valid lines are found in the input.
//     Meta: EMK(0, "string") - the whole input string.
func (Parser2) Parse(code string, _ ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	result := []ParsedNode{}

	for _, rawLine := range strings.Split(code, "\n") {
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

		result = append(result, ParsedNode{
			Raw:    rawLine,
			Switch: command,
			Metadata: map[string]interface{}{
				"command": command,
				"args":    args,
				"__raw":   rawLine,
			},
		})
	}

	if len(result) == 0 {
		return nil, core.Err(errors.ParsingError, "No valid lines found in input").
			WithMeta(core.EMK(0, "string"), code)
	}
	return addPrevNextNodes(result), nil
}

func (Parser2) String() string {
	return "lc/parsing/stringParsing/Parser2"
}
