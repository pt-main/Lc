package astools

import "github.com/pt-main/lc/parsing/stringParsing"

func GetChildren(node *stringParsing.ParsedNode) []stringParsing.ParsedNode {
	if node == nil {
		return nil
	}
	if children, ok := node.Metadata["children"].([]stringParsing.ParsedNode); ok {
		return children
	}
	return nil
}

func FindChild(node *stringParsing.ParsedNode, switchName string) *stringParsing.ParsedNode {
	for _, child := range GetChildren(node) {
		if child.Switch == switchName {
			return &child
		}
	}
	return nil
}

func FindChildIndex(node *stringParsing.ParsedNode, switchName string) int {
	if node == nil {
		return -1
	}
	children := GetChildren(node)
	for i, child := range children {
		if child.Switch == switchName {
			return i
		}
	}
	return -1
}

func FindChildren(node *stringParsing.ParsedNode, switchName string) []stringParsing.ParsedNode {
	var result []stringParsing.ParsedNode
	for _, child := range GetChildren(node) {
		if child.Switch == switchName {
			result = append(result, child)
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
	if val, ok := node.Metadata["value"].(string); ok {
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
			childPath := append(f.path, getNodeName(node))
			stack = append(stack, frame{child, childPath})
		}
	}
	return nil
}

func getNodeName(n *stringParsing.ParsedNode) string {
	if n.Switch != "" {
		return n.Switch
	}
	return ""
}
