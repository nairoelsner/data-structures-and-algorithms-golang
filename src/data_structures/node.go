package datastructures

type Node struct {
	value any
	next  *Node
}

func (n *Node) GetValue() any {
	return n.value
}
