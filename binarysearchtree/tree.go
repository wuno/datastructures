package binarysearchtree

import (
	"errors"
)

// Tree struct that holds our nodes and provides extra functionality
type Tree struct {
	Root *Node
}

// Insert method for the Tree struct
func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

// Find method for the Tree struct
func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

// Delete method for the Tree struct
func (t *Tree) Delete(s string) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// Traverse method for the Tree struct
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}
