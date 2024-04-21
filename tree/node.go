package tree

import (
	"errors"
	"fmt"

	"github.com/negrel/paon/events"
)

type NodeData interface {
	events.Target

	setParentNode(any)
	setNextNode(any)
	setPreviousNode(any)
}

type Node[T NodeData] interface {
	UnwrapNode() *T
	ParentNode() *T
	RootNode() *T
	NextNode() *T
	PreviousNode() *T
	FirstChildNode() *T
	LastChildNode() *T
	AppendChildNode(*T) error
	InsertBeforeNode(newChild, reference *T) error
	RemoveChildNode(*T) error
	IsAncestorOfNode(*T) bool
	IsDescendantOfNode(*T) bool
}

// node define a single node in a tree.
type node[T NodeData] struct {
	value T

	parent Node[T]

	// Siblings
	next     Node[T]
	previous Node[T]

	// Children
	firstChild Node[T]
	lastChild  Node[T]
}

// NewNode returns a new node containing the given data.
func NewNode[T NodeData](data T) Node[T] {
	return newNode(data)
}

func newNode[T NodeData](data T) *node[T] {
	return &node[T]{
		value:      data,
		parent:     nil,
		next:       nil,
		previous:   nil,
		firstChild: nil,
		lastChild:  nil,
	}
}

// Unwrap returns the underlying data contained in the node.
func (n *node[T]) Unwrap() T {
	return n.value
}

// Swap swaps underlying data contained in the node and returns old one.
func (n *node[T]) Swap(v T) T {
	old := n.value
	n.value = v
	return old
}

// NextNode returns the next sibling of this node.
func (n *node[T]) NextNode() Node[T] {
	return n.next
}

// Previous implements node.
func (n *node[T]) PreviousNode() *T {
	return n.previous.UnwrapNode()
}

// Parent returns parent node.
func (n *node[T]) ParentNode() *T {
	return n.parent.UnwrapNode()
}

// RootNode returns root node of the tree.
func (n *node[T]) RootNode() *T {
	if n.parent != nil {
		return n.parent.RootNode()
	}

	return n.Unwrap()
}

// FirstChild returns first direct child of this node.
func (n *node[T]) FirstChild() Node[T] {
	return n.firstChild
}

// LastChild returns last direct child of this node.
func (n *node[T]) LastChild() Node[T] {
	return n.lastChild
}

// Appends the given child to the list of child node. An error is returned
// if the given child is an ancestor of this node.
// When an error is returned, the child is left unchanged
func (n *node[T]) AppendChild(newChild Node[T]) (err error) {
	if err = n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	n.appendChild(newChild)

	return nil
}

func (n *node[T]) appendChild(newChild Node[T]) {
	n.prepareChildForInsertion(newChild)
	newChild.UnwrapNode().setParentNode(n)

	if n.lastChild != nil {
		n.lastChild.UnwrapNode().setNextNode(newChild)
		newChild.UnwrapNode().setPreviousNode(n.lastChild)
	} else {
		n.firstChild = newChild
	}

	n.lastChild = newChild
}

func (n *node[T]) ensurePreInsertionValidity(child Node[T]) error {
	if child == nil {
		return errors.New("child is nil")
	}

	// check if child is not a parentnode of pn
	if n.IsDescendantOf(child) {
		return errors.New("child contains the parentnode")
	}

	return nil
}

func (n *node[T]) prepareChildForInsertion(newChild Node[T]) {
	if parent := newChild.ParentNode(); parent != nil {
		_ = parent.RemoveChild(newChild)
	}
}

// Inserts the given child before the given reference child node. If the
// reference is nil, the child is appended. An error is returned
// if the given child is an ancestor of this node or if the reference
// is not a direct child of this.
// When an error is returned, the child is left unchanged
func (n *node[T]) InsertBefore(newChild, reference Node[T]) error {
	// InsertBeforenode(node, nil) is equal to AppendChildnode(node)
	if reference == nil {
		return n.AppendChild(newChild)
	}
	if referenceIsNotChild := n != reference.ParentNode(); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	n.insertBefore(reference, newChild)
	return nil
}

func (n *node[T]) insertBefore(reference, newChild Node[T]) {
	n.prepareChildForInsertion(newChild)
	newChild.setParentNode(n)

	if previous := reference.PreviousNode(); previous != nil {
		previous.setNextNode(newChild)
		newChild.setPreviousNode(previous)
	} else {
		n.firstChild = newChild
	}
	reference.setPreviousNode(newChild)
	newChild.setNextNode(reference)
}

// Removes the given direct child node of this. Returns an error otherwise.
// When an error is returned, the child is left unchanged
func (n *node[T]) RemoveChild(child Node[T]) error {
	// if not a child
	if n != child.ParentNode() {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	n.removeChild(child)

	return nil
}

func (n *node[T]) removeChild(child Node[T]) {
	// Removing parentnode & root link
	child.setParentNode(nil)

	// Removing siblings link
	next := child.NextNode()
	prev := child.PreviousNode()
	if next != nil {
		child.setNextNode(nil)
		next.setPreviousNode(prev)
	} else {
		n.lastChild = prev
	}

	if prev != nil {
		child.setPreviousNode(nil)
		prev.setNextNode(next)
	} else {
		n.firstChild = next
	}
}

// IsAncestorOf returns true if the given node is a child of this one.
func (n *node[T]) IsAncestorOf(other Node[T]) bool {
	if other == nil {
		return false
	}

	return other.IsDescendantOf(n)
}

// IsDescendantOf returns true if this node is a descendant of the given one.
func (n *node[T]) IsDescendantOf(parent Node[T]) bool {
	if parent == nil {
		return false
	}

	var node Node[T] = n
	for node != nil {
		if node == parent {
			return true
		}

		node = node.ParentNode()
	}

	return false
}

func (n *node[T]) setNextNode(nn Node[T]) {
	n.next = nn
}

func (n *node[T]) setPreviousNode(nn Node[T]) {
	n.previous = nn
}

func (n *node[T]) setParentNode(nn Node[T]) {
	n.parent = nn
}
