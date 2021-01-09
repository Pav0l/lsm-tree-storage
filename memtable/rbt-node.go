package memtable

import "unsafe"

type color bool

// Red and Black are a colors of RBT
const (
	Red   color = false
	Black color = true
)

// EmptyNodeSize is the default size of Node in bytes
const (
	EmptyNodeSize int = int(unsafe.Sizeof(Node{}))
)

// Node serves as basic node in Red-black tree
type Node struct {
	Color  color
	key    string
	value  string
	parent *Node
	left   *Node
	right  *Node
}

// EstimateSize returns calculated size of Node in bytes
func (n *Node) EstimateSize() int {
	return EmptyNodeSize + len(n.key) + len(n.value)
}

// GetParent returns the Parent node
func GetParent(n *Node) *Node {
	return n.parent
}

// GetSibling returns the Sibling node
func GetSibling(n *Node) *Node {
	parent := GetParent(n)

	if parent == nil {
		// no parent == no sibling
		return nil
	}

	if parent.left == n {
		return parent.right
	}
	return parent.left
}

// GetGrandParent returns the GrandParent node
func GetGrandParent(n *Node) *Node {
	parent := GetParent(n)
	if parent == nil {
		return nil
	}
	return GetParent(parent)
}

// GetUncle returns the Uncle node
func GetUncle(n *Node) *Node {
	parent := GetParent(n)
	if parent == nil {
		return nil
	}

	return GetSibling(parent)
}
