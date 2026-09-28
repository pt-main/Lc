package stringParsing

// addPrevNextNodes links every node to its neighbours through the __prev and
// __next metadata keys, so a node can walk the stream in either direction.
func addPrevNextNodes(nodes []ParsedNode) []ParsedNode {
	for i := range nodes {
		if i > 0 {
			nodes[i].Metadata["__prev"] = &nodes[i-1]
		} else {
			nodes[i].Metadata["__prev"] = nil
		}
		if i < len(nodes)-1 {
			nodes[i].Metadata["__next"] = &nodes[i+1]
		} else {
			nodes[i].Metadata["__next"] = nil
		}
	}
	return nodes
}
