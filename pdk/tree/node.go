package tree

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/id"
)

// Node define an element in a Node tree.
type Node interface {
	id.Identifiable

	// Unwrap returns the data stored in this Node.
	Unwrap() interface{}

	// Returns true if the given Node is the same as this.
	IsSame(Node) bool

	// Next sibling.
	Next() Node
	SetNext(Node)

	// Previous sibling.
	Previous() Node
	SetPrevious(Node)

	// ParentNode is the direct parentNode of the Node.
	Parent() Node
	SetParent(Node)

	// RootNode defines the root of the Node tree.
	Root() Node

	// Returns the first child Node of this.
	FirstChild() Node

	// Returns the last child Node of this.
	LastChild() Node

	// Appends the given child to the list of child Node. An error is returned
	// if the given child is an ancestor of this Node.
	AppendChild(newChild Node) error

	// Inserts the given child before the given reference child Node. If the
	// reference is nil, the child is appended. An error is returned
	// if the given child is an ancestor of this Node or if the reference
	// is not a direct child of this.
	InsertBefore(reference, newChild Node) error

	// Removes the given direct child Node of this. Returns an error otherwise.
	RemoveChild(child Node) error

	// IsAncestorOf returns true if the given Node is a child of this.
	IsAncestorOf(child Node) bool

	// IsDescendantOf returns true if this is a descendant of the given Node.
	IsDescendantOf(node Node) bool
}

type node struct {
	*leafNode

	firstChild Node
	lastChild  Node
}

// NewNode returns a Node wrapping the given data.
func NewNode(data interface{}) Node {
	return newNode(data)
}

func newNode(data interface{}) *node {
	n := &node{
		leafNode: newLeafNode(data),
	}

	return n
}

func (n *node) FirstChild() Node {
	return n.firstChild
}

func (n *node) LastChild() Node {
	return n.lastChild
}

func (n *node) AppendChild(newChild Node) (err error) {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	if err = n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	n.appendChild(newChild)

	return nil
}

func (n *node) appendChild(newChild Node) {
	n.prepareChildForInsertion(newChild)
	newChild.SetParent(n)

	if n.lastChild != nil {
		n.lastChild.SetNext(newChild)
		newChild.SetPrevious(n.lastChild)
	} else {
		n.firstChild = newChild
	}

	n.lastChild = newChild
}

func (n *node) ensurePreInsertionValidity(child Node) error {
	if child == nil {
		return errors.New("child is nil")
	}

	// check if child is not a parentNode of pn
	if child.IsAncestorOf(n) {
		return errors.New("child contains the parentNode")
	}

	return nil
}

func (n *node) prepareChildForInsertion(newChild Node) {
	if parent := newChild.Parent(); parent != nil {
		err := parent.RemoveChild(newChild)
		assert.Nil(err)
	}
	assert.Nil(newChild.Root())
	assert.Nil(newChild.Parent())
	assert.Nil(newChild.Previous())
	assert.Nil(newChild.Next())
}

func (n *node) InsertBefore(reference, newChild Node) error {
	assert.NotNil(newChild, "child must be non-nil to be appended")

	// InsertBeforeNode(nil, node) is equal to AppendChildNode(node)
	if reference == nil {
		return n.AppendChild(newChild)
	}
	if referenceIsNotChild := !n.IsSame(reference.Parent()); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	n.insertBefore(reference, newChild)
	return nil
}

func (n *node) insertBefore(reference, newChild Node) {
	n.prepareChildForInsertion(newChild)

	if previous := reference.Previous(); previous != nil {
		previous.SetNext(newChild)
		newChild.SetPrevious(previous)
	} else {
		n.firstChild = newChild
	}
	reference.SetPrevious(newChild)
	newChild.SetNext(reference)

	newChild.SetParent(n)
}

func (n *node) RemoveChild(child Node) error {
	assert.NotNil(child, "child must be non-nil to be removed")

	// if not a child of pn
	if !n.IsSame(child.Parent()) {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	n.removeChild(child)

	return nil
}

func (n *node) removeChild(child Node) {
	// Removing siblings link
	next := child.Next()
	prev := child.Previous()
	if next != nil {
		child.SetNext(nil)
		next.SetPrevious(prev)
	} else {
		n.lastChild = prev
	}

	if prev != nil {
		child.SetPrevious(nil)
		prev.SetNext(next)
	} else {
		n.firstChild = next
	}
	// Removing parentNode & root link
	child.SetParent(nil)
}
