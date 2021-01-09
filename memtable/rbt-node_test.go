package memtable

import (
	"os"
	"testing"
)

var grandpar = Node{key: "4"}
var unc = Node{key: "3"}
var par = Node{key: "5"}
var child = Node{key: "6"}

func setupNodes() {
	grandpar.left = &unc
	grandpar.right = &par

	unc.parent = &grandpar

	par.parent = &grandpar
	par.right = &child

	child.parent = &par
}

func TestMain(m *testing.M) {
	setupNodes()

	code := m.Run()
	os.Exit(code)
}

func TestNode_EstimateSize(t *testing.T) {
	t.Run("Default size", func(t *testing.T) {
		n := Node{}
		size := n.EstimateSize()
		if size != 64 {
			t.Errorf("Default size of Node is not correct, expected: %v, received %v", 64, size)
		}
	})

	t.Run("Size with key, value strings", func(t *testing.T) {
		n := Node{key: "kung", value: "fu"}
		size := n.EstimateSize()
		if size != 70 {
			t.Errorf("Default size of Node is not correct, expected: %v, received %v", 70, size)
		}
	})
}

func TestGetParent(t *testing.T) {
	t.Run("with parent", func(t *testing.T) {
		parent := GetParent(&child)
		if parent != &par {
			t.Errorf("GetParent has failed, expected: %+v, received %+v", par, parent)
		}
	})

	t.Run("without parent", func(t *testing.T) {
		parent := GetParent(&grandpar)
		if parent != nil {
			t.Errorf("GetParent without parent has failed, expected: %+v, received %+v", nil, parent)
		}
	})
}

func TestGetSibling(t *testing.T) {
	t.Run("with sibling", func(t *testing.T) {
		sibling := GetSibling(&par)
		if sibling != &unc {
			t.Errorf("GetSibling has failed, expected: %+v, received %+v", unc, sibling)
		}
	})

	t.Run("without sibling", func(t *testing.T) {
		sibling := GetSibling(&child)
		if sibling != nil {
			t.Errorf("GetSibling without sibling has failed, expected: %+v, received %+v", nil, sibling)
		}
	})
}

func TestGetGrandParent(t *testing.T) {
	t.Run("with grandparent", func(t *testing.T) {
		grandParent := GetGrandParent(&child)
		if grandParent != &grandpar {
			t.Errorf("GetGrandParent has failed, expected: %+v, received %+v", grandpar, grandParent)
		}
	})

	t.Run("without grandparent", func(t *testing.T) {
		grandParent := GetGrandParent(&par)
		if grandParent != nil {
			t.Errorf("GetGrandParent without grandparent has failed, expected: %+v, received %+v", nil, grandParent)
		}
	})
}

func TestGetUncle(t *testing.T) {
	t.Run("with uncle", func(t *testing.T) {
		uncle := GetUncle(&child)
		if uncle != &unc {
			t.Errorf("GetUncle has failed, expected: %+v, received %+v", unc, uncle)
		}
	})

	t.Run("without uncle", func(t *testing.T) {
		uncle := GetUncle(&par)
		if uncle != nil {
			t.Errorf("GetUncle without uncle has failed, expected: %+v, received %+v", nil, uncle)
		}
	})
}
