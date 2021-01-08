package memtable

// RedBlackTree holds the tree data structure
type RedBlackTree struct {
	Size uint
	Root *Node
}

// Insert creates a node with value and adds it to the tree
func (rbt *RedBlackTree) Insert(key string) *Node {
	// insert new node and color it red
	n := Node{Color: Red, key: key}

	rbt.insertNode(&n, rbt.Root)

	rbt.insertFixTree(&n)

	rbt.Root = &n
	for GetParent(rbt.Root) != nil {
		rbt.Root = GetParent(rbt.Root)
	}

	rbt.Size++

	return rbt.Root
}

// Search finds a key in the tree
func (rbt *RedBlackTree) Search(key string, root *Node) string {
	if root.key == key {
		return root.key
	}

	if root.left != nil && key < root.key {
		return rbt.Search(key, root.left)
	} else if root.right != nil && key > root.key {
		return rbt.Search(key, root.right)
	}

	return ""
}

func (rbt *RedBlackTree) rotateLeft(n *Node) {
	newParent := n.right
	if newParent == nil {
		panic("Node does not have a right child")
	}

	newRight := newParent.left

	nParent := GetParent(n)
	newParent.parent = nParent
	// n might be root, but if not fix n parents children
	if nParent != nil {
		if nParent.left == n {
			nParent.left = newParent
		} else if nParent.right == n {
			nParent.right = newParent
		}
	}

	n.parent = newParent
	newParent.left = n

	n.right = newRight
	if n.right != nil {
		newRight.parent = n
	}
}

func (rbt *RedBlackTree) rotateRight(n *Node) {
	newParent := n.left
	if newParent == nil {
		panic("Node does not have a left child")
	}

	newLeft := newParent.right

	nParent := GetParent(n)
	newParent.parent = nParent
	// n might be root, but if not fix n parents children
	if nParent != nil {
		if nParent.left == n {
			nParent.left = newParent
		} else if nParent.right == n {
			nParent.right = newParent
		}
	}

	n.parent = newParent
	newParent.right = n

	n.left = newLeft
	if n.left != nil {
		newLeft.parent = n
	}
}

func (rbt *RedBlackTree) insertNode(n *Node, root *Node) {
	if root != nil {
		if n.key < root.key {
			if root.left != nil {
				rbt.insertNode(n, root.left)
				return
			}
			root.left = n
		} else {
			// n.key >= root.key
			if root.right != nil {
				rbt.insertNode(n, root.right)
				return
			}
			root.right = n
		}
	}

	n.parent = root
}

func (rbt *RedBlackTree) insertFixTree(n *Node) {
	if GetParent(n) == nil {
		// root does not exist -> add BLACK root
		n.Color = Black
	} else if GetParent(n).Color == Black {
		// root exist and parent is black. so we inserted red node and didn't changed any rules
		// tree is still valid, so just return
		return
	} else if uncle := GetUncle(n); uncle != nil && uncle.Color == Red {
		uncle.Color = Black
		GetParent(n).Color = Black
		GetGrandParent(n).Color = Red
		// recursively fix tree with grandparent as n, in case grandparent was root and is now red
		rbt.insertFixTree(GetGrandParent(n))
	} else {
		// uncle is BLACK (either a triangle or a line)
		rbt.rotateParent(n)
	}
}

func (rbt *RedBlackTree) rotateParent(n *Node) {
	p := GetParent(n)
	g := GetGrandParent(n)

	if n == p.right && p == g.left {
		rbt.rotateLeft(p)
		n = n.left
	} else if n == p.left && p == g.right {
		rbt.rotateRight(p)
		n = n.right
	}

	rbt.rotateGrandParentAndRecolor(n)
}

func (rbt *RedBlackTree) rotateGrandParentAndRecolor(n *Node) {
	p := GetParent(n)
	g := GetGrandParent(n)

	if n == p.left {
		rbt.rotateRight(g)
	} else {
		rbt.rotateLeft(g)
	}

	p.Color = Black
	g.Color = Red
}
