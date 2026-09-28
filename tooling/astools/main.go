package astools

import "github.com/pt-main/lc/parsing/stringParsing"

// GetChildren returns the child slice stored in the node metadata. The slice
// shares its backing array with the tree, so &children[i] is a real pointer
// into it. Appending to the returned slice can reallocate that array, after
// which previously returned pointers no longer refer to the tree; re-look-up
// instead of caching a pointer across a mutation.
func GetChildren(node *stringParsing.ParsedNode) []stringParsing.ParsedNode {
	if node == nil {
		return nil
	}
	if children, ok := node.Metadata["children"].([]stringParsing.ParsedNode); ok {
		return children
	}
	return nil
}

// FindChild returns the first child with the given switch name, or nil.
func FindChild(node *stringParsing.ParsedNode, switchName string) *stringParsing.ParsedNode {
	children := GetChildren(node)
	for i := range children {
		if children[i].Switch == switchName {
			return &children[i]
		}
	}
	return nil
}

// FindChildIndex returns the index of the first child with the given switch
// name, or -1 when there is none.
func FindChildIndex(node *stringParsing.ParsedNode, switchName string) int {
	children := GetChildren(node)
	for i := range children {
		if children[i].Switch == switchName {
			return i
		}
	}
	return -1
}

// FindChildren returns every child with the given switch name.
func FindChildren(node *stringParsing.ParsedNode, switchName string) []stringParsing.ParsedNode {
	var result []stringParsing.ParsedNode
	children := GetChildren(node)
	for i := range children {
		if children[i].Switch == switchName {
			result = append(result, children[i])
		}
	}
	return result
}

func GetTokenValue(node *stringParsing.ParsedNode) string {
	if node == nil {
		return ""
	}
	if node.Raw != "" {
		return node.Raw
	}
	// The lexer stores the matched text under "__value".
	if val, ok := node.Metadata["__value"].(string); ok {
		return val
	}
	return ""
}

func GetChildAt(node *stringParsing.ParsedNode, index int) *stringParsing.ParsedNode {
	children := GetChildren(node)
	if index >= 0 && index < len(children) {
		return &children[index]
	}
	return nil
}

func Walk(node *stringParsing.ParsedNode, fn func(*stringParsing.ParsedNode) error) error {
	if node == nil {
		return nil
	}

	stack := []*stringParsing.ParsedNode{node}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if err := fn(cur); err != nil {
			return err
		}

		children := GetChildren(cur)
		for i := len(children) - 1; i >= 0; i-- {
			stack = append(stack, &children[i])
		}
	}
	return nil
}

func WalkWithPath(node *stringParsing.ParsedNode, fn func(*stringParsing.ParsedNode, []string) error) error {
	if node == nil {
		return nil
	}

	type frame struct {
		node *stringParsing.ParsedNode
		path []string
	}
	stack := []frame{{node, []string{getNodeName(node)}}}

	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if err := fn(f.node, f.path); err != nil {
			return err
		}

		children := GetChildren(f.node)

		for i := len(children) - 1; i >= 0; i-- {
			child := &children[i]
			// A fresh slice per child: appending to f.path directly would let
			// siblings share one backing array and overwrite each other.
			childPath := make([]string, len(f.path), len(f.path)+1)
			copy(childPath, f.path)
			childPath = append(childPath, getNodeName(child))
			stack = append(stack, frame{child, childPath})
		}
	}
	return nil
}

func getNodeName(n *stringParsing.ParsedNode) string {
	return n.Switch
}
