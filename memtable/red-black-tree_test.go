package memtable

import (
	"testing"
)

func TestRedBlackTree_rotateLeft(t *testing.T) {
	child := Node{key: "3"}
	rightNode := Node{key: "2"}
	node := Node{key: "1"}
	root := Node{key: "0"}
	root.left = &node
	node.parent = &root
	node.right = &rightNode
	rightNode.parent = &node
	rightNode.left = &child
	child.parent = &rightNode

	rbt := RedBlackTree{Root: &root}

	rbt.rotateLeft(&node)

	if root.left != &rightNode {
		t.Error("Root left was not rotated properly, expected:", &rightNode, "received:", root.left)
	}
	if rightNode.parent != &root {
		t.Error("rightNode parent was not rotated properly, expected:", &root, "received:", rightNode.parent)
	}
	if rightNode.left != &node {
		t.Error("rightNode left was not rotated properly, expected:", &node, "received:", rightNode.left)
	}
	if node.parent != &rightNode {
		t.Error("Nodes parent was not rotated properly, expected:", &rightNode, "received:", node.parent)
	}
	if node.right != &child {
		t.Error("Nodes right was not rotated properly, expected:", &child, "received:", node.right)
	}
	if child.parent != &node {
		t.Error("Child parent was not rotated properly, expected:", &node, "received:", child.parent)
	}
}

func TestRedBlackTree_rotateRight(t *testing.T) {
	child := Node{key: "3"}
	leftNode := Node{key: "2"}
	node := Node{key: "1"}
	root := Node{key: "0"}
	root.left = &node
	node.parent = &root
	node.left = &leftNode
	leftNode.parent = &node
	leftNode.right = &child
	child.parent = &leftNode

	rbt := RedBlackTree{Root: &root}

	rbt.rotateRight(&node)

	if root.left != &leftNode {
		t.Error("Root left was not rotated properly, expected:", &leftNode, "received:", root.left)
	}
	if leftNode.parent != &root {
		t.Error("leftNode parent was not rotated properly, expected:", &root, "received:", leftNode.parent)
	}
	if leftNode.right != &node {
		t.Error("leftNode right was not rotated properly, expected:", &node, "received:", leftNode.right)
	}
	if node.parent != &leftNode {
		t.Error("Nodes parent was not rotated properly, expected:", &leftNode, "received:", node.parent)
	}
	if node.left != &child {
		t.Error("Nodes left was not rotated properly, expected:", &child, "received:", node.left)
	}
	if child.parent != &node {
		t.Error("Child parent was not rotated properly, expected:", &node, "received:", child.parent)
	}
}

func TestRedBlackTree_insertNode(t *testing.T) {

	t.Run("adds node in proper position", func(t *testing.T) {
		root := Node{key: "5"}
		rbt := RedBlackTree{Root: &root}

		node2 := Node{key: "2"}
		rbt.insertNode(&node2, &root)
		if root.left != &node2 {
			t.Error("Invalid node insertion, expected:", &node2, "received:", root.left)
		}

		node3 := Node{key: "3"}
		rbt.insertNode(&node3, &root)
		if node2.right != &node3 {
			t.Error("Invalid node insertion, expected:", &node3, "received:", node2.right)
		}

		node6 := Node{key: "6"}
		rbt.insertNode(&node6, &root)
		if root.right != &node6 {
			t.Error("Invalid node insertion, expected:", &node6, "received:", root.right)
		}

		node8 := Node{key: "8"}
		rbt.insertNode(&node8, &root)
		if node6.right != &node8 {
			t.Error("Invalid node insertion, expected:", &node8, "received:", node6.right)
		}

		node1 := Node{key: "1"}
		rbt.insertNode(&node1, &root)
		if node2.left != &node1 {
			t.Error("Invalid node insertion, expected:", &node1, "received:", node2.left)
		}

		node7 := Node{key: "7"}
		rbt.insertNode(&node7, &root)
		if node8.left != &node7 {
			t.Error("Invalid node insertion, expected:", &node7, "received:", node8.left)
		}

		node5 := Node{key: "5"}
		rbt.insertNode(&node5, &root)
		if node6.left != &node5 {
			t.Error("Invalid node insertion, expected:", &node5, "received:", node6.left)
		}
	})
}

func setupTree(keys []string) *RedBlackTree {
	rbt := new(RedBlackTree)

	for i := 0; i < len(keys); i++ {
		rbt.Insert(keys[i])
	}
	return rbt
}
func TestRedBlackTree_Insert(t *testing.T) {

	t.Run("rotates around grandparent", func(t *testing.T) {
		keys := []string{"7", "6", "5"}
		rbt := setupTree(keys[:])

		type result struct {
			expectedVal, receivedVal     string
			expectedColor, receivedColor color
		}

		results := []result{
			{"6", rbt.Root.key, Black, rbt.Root.Color},
			{"5", rbt.Root.left.key, Red, rbt.Root.left.Color},
			{"7", rbt.Root.right.key, Red, rbt.Root.right.Color},
		}

		for i := 0; i < len(results); i++ {
			r := results[i]
			if r.expectedVal != r.receivedVal || r.expectedColor != r.receivedColor {
				t.Error("Invalid node key or color, expected:", r.expectedVal, r.expectedColor, "received:", r.receivedVal, r.receivedColor)
			}
		}

		if actual := rbt.Size; actual != uint(len(keys)) {
			t.Error("Invalid tree size, expected:", len(keys), "received:", actual)
		}
	})

	t.Run("multiple rotations and recolors", func(t *testing.T) {
		keys := []string{"100", "50", "20", "40", "45", "120", "110", "90", "111"}
		rbt := setupTree(keys[:])

		type result struct {
			expectedVal, receivedVal     string
			expectedColor, receivedColor color
		}

		results := []result{
			{"20", rbt.Root.key, Black, rbt.Root.Color},
			{"110", rbt.Root.left.key, Red, rbt.Root.left.Color},
			{"100", rbt.Root.left.left.key, Black, rbt.Root.left.left.Color},
			{"120", rbt.Root.left.right.key, Black, rbt.Root.left.right.Color},
			{"45", rbt.Root.right.key, Red, rbt.Root.right.Color},
			{"40", rbt.Root.right.left.key, Black, rbt.Root.right.left.Color},
			{"50", rbt.Root.right.right.key, Black, rbt.Root.right.right.Color},
			{"90", rbt.Root.right.right.right.key, Red, rbt.Root.right.right.right.Color},
			{"111", rbt.Root.left.right.left.key, Red, rbt.Root.left.right.left.Color},
		}

		for i := 0; i < len(results); i++ {
			r := results[i]
			if r.expectedVal != r.receivedVal || r.expectedColor != r.receivedColor {
				t.Error("Invalid node key or color, expected:", r.expectedVal, r.expectedColor, "received:", r.receivedVal, r.receivedColor)
			}
		}

		if actual := rbt.Size; actual != uint(len(keys)) {
			t.Error("Invalid tree size, expected:", len(keys), "received:", actual)
		}
	})
}

func TestRedBlackTree_Search(t *testing.T) {

	t.Run("finds an existing key", func(t *testing.T) {
		keys := []string{"100", "50", "20", "40", "45", "120", "110", "90", "111"}
		rbt := setupTree(keys[:])

		result := rbt.Search("111", rbt.Root)

		if result != "111" {
			t.Error("Invalid search result, expected:", "111", "received:", result)
		}
	})

	t.Run("returns \"\" when key does not exist", func(t *testing.T) {
		keys := []string{"100", "50", "20", "40", "45", "120", "110", "90", "111"}
		rbt := setupTree(keys[:])

		result := rbt.Search("99", rbt.Root)

		if result != "" {
			t.Error("Invalid search result, expected:", "", "received:", result)
		}
	})
}
