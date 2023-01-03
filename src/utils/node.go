package utils

type Node struct {
	value    string
	next     *Node
	previous *Node
}

func NewNode(val string) *Node {
	n := Node{
		value:    val,
		next:     nil,
		previous: nil,
	}
	return &n
}
