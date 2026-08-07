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
	if err := fn(node); err != nil {
		return err
	}
	for _, child := range GetChildren(node) {
		if err := Walk(&child, fn); err != nil {
			return err
		}
	}
	return nil
}
