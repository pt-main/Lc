package astools

import "github.com/pt-main/lc/v2/parsing/stringParsing"

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
	children := GetChildren(node)
	matches := 0
	for i := range children {
		if children[i].Switch == switchName {
			matches++
		}
	}
	if matches == 0 {
		return nil
	}
	result := make([]stringParsing.ParsedNode, 0, matches)
	for i := range children {
		if children[i].Switch == switchName {
			result = append(result, children[i])
		}
	}
	return result
}

// FindChildrenPointers returns a pointer to every child with the given switch
// name. Unlike FindChildren, which returns copies, the pointers address the
// nodes in the tree, so a write through one reaches the tree itself. The
// pointers share their backing array with the tree, so re-look-up after any
// mutation instead of caching one across it.
func FindChildrenPointers(node *stringParsing.ParsedNode, switchName string) []*stringParsing.ParsedNode {
	children := GetChildren(node)
	matches := 0
	for i := range children {
		if children[i].Switch == switchName {
			matches++
		}
	}
	if matches == 0 {
		return nil
	}
	result := make([]*stringParsing.ParsedNode, 0, matches)
	for i := range children {
		if children[i].Switch == switchName {
			result = append(result, &children[i])
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

// pathArena hands out path slices carved from a few large blocks instead of
// one make per node. Every slice keeps its backing array private by capping it
// to its own length, so a caller that appends to a path cannot reach a
// sibling's storage and a path stays valid after the walk has moved on.
type pathArena struct {
	block []string
	used  int
	size  int
}

// Blocks grow geometrically: a fixed size wastes the unusable tail of every
// chunk, while with a growing one the waste stays proportional to the largest
// block instead of to the whole tree.
const (
	pathArenaFirstBlock = 64
	pathArenaMaxBlock   = 4096
)

func (a *pathArena) next(n int) []string {
	if len(a.block)-a.used < n {
		size := a.size
		switch {
		case size == 0:
			size = pathArenaFirstBlock
		case size < pathArenaMaxBlock:
			size *= 2
		}
		if n > size {
			size = n
		}
		a.size = size
		a.block = make([]string, size)
		a.used = 0
	}
	start := a.used
	a.used += n
	return a.block[start:a.used:a.used]
}

func WalkWithPath(node *stringParsing.ParsedNode, fn func(*stringParsing.ParsedNode, []string) error) error {
	if node == nil {
		return nil
	}

	type frame struct {
		node *stringParsing.ParsedNode
		path []string
	}
	stack := make([]frame, 0, 16)
	var arena pathArena
	stack = append(stack, frame{node, arena.next(1)})
	stack[0].path[0] = getNodeName(node)

	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if err := fn(f.node, f.path); err != nil {
			return err
		}

		children := GetChildren(f.node)
		if len(children) == 0 {
			continue
		}

		// One slice holds the child paths back to back, so a parent with n
		// children costs a single allocation instead of n. Every slot is
		// sized len(f.path)+1 and trimmed to its length, keeping the backing
		// array private to that child.
		stride := len(f.path) + 1
		names := arena.next(len(children) * stride)
		for i := len(children) - 1; i >= 0; i-- {
			child := &children[i]
			childPath := names[i*stride : i*stride+stride : i*stride+stride]
			copy(childPath, f.path)
			childPath[stride-1] = getNodeName(child)
			stack = append(stack, frame{child, childPath})
		}
	}

	return nil
}

func getNodeName(n *stringParsing.ParsedNode) string {
	return n.Switch
}
