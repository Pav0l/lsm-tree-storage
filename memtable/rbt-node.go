package memtable

type color bool

// Red and Black are a colors of RBT
const (
	Red   color = false
	Black color = true
)

// Node serves as basic node in Red-black tree
type Node struct {
	Color  color
	key    string
	parent *Node
	left   *Node
	right  *Node
}

// InvertColor of the Node on which this method is called on
func (n *Node) InvertColor() {
	if n.Color == Black {
		n.Color = Red
	} else {
		n.Color = Black
	}
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
