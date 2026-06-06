package stringParsing

func addPrevNextNodes(nodes []ParsedNode) []ParsedNode {
	for i := 0; i < len(nodes); i++ {
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
