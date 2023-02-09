package utils

// Node represents a two pointer list node
type Node struct {
	value    string
	next     *Node
	previous *Node
}

// NewNode returns a pointer to a Node
func NewNode(val string) *Node {
	n := Node{
		value:    val,
		next:     nil,
		previous: nil,
	}
	return &n
}
