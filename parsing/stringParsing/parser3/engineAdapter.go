package parser3

import (
	"reflect"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

// Adapter wraps a Parser for the engine: it runs the parse and unwraps the
// root node, returning the children the start rule produced.
type Adapter struct {
	Parser *Parser
}

// Parse runs the wrapped parser and returns the start rule's children.
//
// Err parser3/adapter: if the parser is missing, fails, or returns a root
// node whose children metadata is not a node list.
func (a *Adapter) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	if a.Parser == nil {
		return nil, &AdapterError{Msg: "adapter has no parser"}
	}

	nodes, err := a.Parser.Parse(code, o...)
	if err != nil {
		return nil, err
	}
	if len(nodes) != 1 {
		return nil, &AdapterError{
			Msg: "expected exactly 1 root node, got " + itoa(len(nodes)),
		}
	}
	return ChildrenOf(nodes[0])
}

// ChildrenOf extracts the child nodes stored under MetaChildren. It accepts
// any slice of ParsedNode so nodes rebuilt by a user action still work.
func ChildrenOf(node stringParsing.ParsedNode) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	raw, ok := node.Metadata[MetaChildren]
	if !ok {
		keys := make([]string, 0, len(node.Metadata))
		for k := range node.Metadata {
			keys = append(keys, k)
		}
		sortStrings(keys)
		return nil, &AdapterError{
			Msg: "node " + quote(node.Switch) + " has no " + quote(MetaChildren) + " metadata (keys: " + joinStrings(keys) + ")",
		}
	}
	if children, ok := raw.([]stringParsing.ParsedNode); ok {
		return children, nil
	}

	rv := reflect.ValueOf(raw)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, &AdapterError{
			Msg: "node " + quote(node.Switch) + " has " + quote(MetaChildren) + " of type " + typeName(raw) + ", expected []ParsedNode",
		}
	}
	children := make([]stringParsing.ParsedNode, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		elem := rv.Index(i)
		if elem.Kind() == reflect.Ptr && !elem.IsNil() {
			elem = elem.Elem()
		}
		if !elem.CanInterface() {
			return nil, &AdapterError{
				Msg: "node " + quote(node.Switch) + " has an unreadable " + quote(MetaChildren) + " element at index " + itoa(i),
			}
		}
		pn, ok := elem.Interface().(stringParsing.ParsedNode)
		if !ok {
			return nil, &AdapterError{
				Msg: "node " + quote(node.Switch) + " has " + quote(MetaChildren) + " element of type " + typeName(elem.Interface()) + " at index " + itoa(i),
			}
		}
		children[i] = pn
	}
	return children, nil
}

func (a *Adapter) String() string {
	return "lc/parsing/stringParsing/parser3/Adapter"
}
