package parser3

import (
	"fmt"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/tap/color"
)

type Adapter struct {
	Parser *Parser
}

func (a *Adapter) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, error) {
	nodes, err := a.Parser.Parse(code, o...)
	if err != nil {
		return nil, fmt.Errorf(
			"%s adapter parse failed: %w",
			color.Set("[?RD]parser3/Adapter[?RT]"),
			err,
		)
	}
	if len(nodes) != 1 {
		return nil, fmt.Errorf(
			"%s expected exactly 1 root node, got %s",
			color.Set("[?RD]parser3/Adapter[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", len(nodes))),
		)
	}
	children, ok := nodes[0].Metadata["children"].([]stringParsing.ParsedNode)
	if !ok {
		return nil, fmt.Errorf(
			"%s root node [?YW]'%s'[?RT] has no 'children' metadata (available keys: %v)",
			color.Set("[?RD]parser3/Adapter[?RT]"),
			nodes[0].Switch,
			func() []string {
				keys := make([]string, 0, len(nodes[0].Metadata))
				for k := range nodes[0].Metadata {
					keys = append(keys, k)
				}
				return keys
			}(),
		)
	}
	return children, nil
}

func (a *Adapter) String() string {
	return "lc/parsing/stringParsing/parser3/Adapter"
}
