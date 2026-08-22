package parser3

import (
	"fmt"
	"reflect"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type Adapter struct {
	Parser *Parser
}

func (a *Adapter) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	nodes, err := a.Parser.Parse(code, o...)
	if err != nil {
		return nil, &AdapterError{
			Msg:   "adapter parse failed",
			Cause: err,
		}
	}
	if len(nodes) != 1 {
		return nil, &AdapterError{
			Msg: fmt.Sprintf("expected exactly 1 root node, got %d", len(nodes)),
		}
	}

	children, ok := nodes[0].Metadata["children"].([]stringParsing.ParsedNode)
	if !ok {
		raw, has := nodes[0].Metadata["children"]
		if !has {
			keys := make([]string, 0, len(nodes[0].Metadata))
			for k := range nodes[0].Metadata {
				keys = append(keys, k)
			}
			return nil, &AdapterError{
				Msg: fmt.Sprintf("root node '%s' has no 'children' metadata (available keys: %v)", nodes[0].Switch, keys),
			}
		}
		rv := reflect.ValueOf(raw)
		if rv.Kind() != reflect.Slice {
			return nil, &AdapterError{
				Msg: fmt.Sprintf("root node '%s' has 'children' metadata of type %T (expected []ParsedNode)", nodes[0].Switch, raw),
			}
		}
		children = make([]stringParsing.ParsedNode, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i).Interface()
			pn, ok := elem.(stringParsing.ParsedNode)
			if !ok {
				return nil, &AdapterError{
					Msg: fmt.Sprintf("root node '%s' has 'children' metadata with wrong element type %T at index %d (expected ParsedNode)", nodes[0].Switch, elem, i),
				}
			}
			children[i] = pn
		}
	}
	return children, nil
}

func (a *Adapter) String() string {
	return "lc/parsing/stringParsing/parser3/Adapter"
}
