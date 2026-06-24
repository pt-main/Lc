package parser3

import (
	"fmt"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type Adapter struct {
	Parser *Parser
}

func (a *Adapter) ParseFlat(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, error) {
	nodes, err := a.Parser.Parse(code)
	if err != nil {
		return nil, err
	}
	if len(nodes) != 1 {
		return nil, fmt.Errorf("expected root node")
	}
	children, ok := nodes[0].Metadata["children"].([]stringParsing.ParsedNode)
	if !ok {
		return nil, fmt.Errorf("root has no children")
	}
	return children, nil
}

func (a *Adapter) String() string {
	return "lc/parsing/stringParsing/parser3/Adapter"
}
