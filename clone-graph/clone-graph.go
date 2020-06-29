package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	mapComplete := make(map[*Node]*Node)
	return clone(node, mapComplete)
}

func clone(node *Node, m map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := m[node]; ok {
		return v
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	m[node] = newNode

	for k, _ := range node.Neighbors {
		newNode.Neighbors[k] = clone(node.Neighbors[k], m)
	}

	return newNode
}
