package tree

import (
	"errors"
	"fmt"
)

// Node define a single Node in a tree.
type Node[T any] struct {
	value T

	parent *Node[T]

	// Siblings
	next     *Node[T]
	previous *Node[T]

	// Children
	firstChild *Node[T]
	lastChild  *Node[T]
}

// NewNode returns a new Node containing the given data.
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		value:      data,
		parent:     nil,
		next:       nil,
		previous:   nil,
		firstChild: nil,
		lastChild:  nil,
	}
}

// Unwrap returns the underlying data contained in the node.
func (n *Node[T]) Unwrap() T {
	return n.value
}

// NextSibling returns the next sibling of this node.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Previous implements Node.
func (n *Node[T]) Previous() *Node[T] {
	return n.previous
}

// Parent returns parent node.
func (n *Node[T]) Parent() *Node[T] {
	return n.parent
}

// Root returns root node of the tree.
func (n *Node[T]) Root() *Node[T] {
	if n.parent != nil {
		return n.parent.Root()
	}

	return n
}

// FirstChild returns first direct child of this node.
func (n *Node[T]) FirstChild() *Node[T] {
	return n.firstChild
}

// LastChild returns last direct child of this node.
func (n *Node[T]) LastChild() *Node[T] {
	return n.lastChild
}

// Appends the given child to the list of child node. An error is returned
// if the given child is an ancestor of this node.
// When an error is returned, the child is left unchanged
func (n *Node[T]) AppendChild(newChild *Node[T]) (err error) {
	if err = n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	n.appendChild(newChild)

	return nil
}

func (n *Node[T]) appendChild(newChild *Node[T]) {
	n.prepareChildForInsertion(newChild)
	newChild.parent = n

	if n.lastChild != nil {
		n.lastChild.next = newChild
		newChild.previous = n.lastChild
	} else {
		n.firstChild = newChild
	}

	n.lastChild = newChild
}

func (n *Node[T]) ensurePreInsertionValidity(child *Node[T]) error {
	if child == nil {
		return errors.New("child is nil")
	}

	// check if child is not a parentNode of pn
	if n.IsDescendantOf(child) {
		return errors.New("child contains the parentNode")
	}

	return nil
}

func (n *Node[T]) prepareChildForInsertion(newChild *Node[T]) {
	if parent := newChild.Parent(); parent != nil {
		_ = parent.RemoveChild(newChild)
	}
}

// Inserts the given child before the given reference child Node. If the
// reference is nil, the child is appended. An error is returned
// if the given child is an ancestor of this node or if the reference
// is not a direct child of this.
// When an error is returned, the child is left unchanged
func (n *Node[T]) InsertBefore(newChild, reference *Node[T]) error {
	// InsertBeforeNode(node, nil) is equal to AppendChildNode(node)
	if reference == nil {
		return n.AppendChild(newChild)
	}
	if referenceIsNotChild := n != reference.Parent(); referenceIsNotChild {
		return errors.New("can't insert child, the given reference is not a child of this node")
	}

	if err := n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't insert child, %v", err)
	}

	n.insertBefore(reference, newChild)
	return nil
}

func (n *Node[T]) insertBefore(reference, newChild *Node[T]) {
	n.prepareChildForInsertion(newChild)
	newChild.parent = n

	if previous := reference.Previous(); previous != nil {
		previous.next = newChild
		newChild.previous = previous
	} else {
		n.firstChild = newChild
	}
	reference.previous = newChild
	newChild.next = reference
}

// Removes the given direct child node of this. Returns an error otherwise.
// When an error is returned, the child is left unchanged
func (n *Node[T]) RemoveChild(child *Node[T]) error {
	// if not a child
	if n != child.Parent() {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	n.removeChild(child)

	return nil
}

func (n *Node[T]) removeChild(child *Node[T]) {
	// Removing parentNode & root link
	child.parent = nil

	// Removing siblings link
	next := child.next
	prev := child.previous
	if next != nil {
		child.next = nil
		next.previous = prev
	} else {
		n.lastChild = prev
	}

	if prev != nil {
		child.previous = nil
		prev.next = next
	} else {
		n.firstChild = next
	}
}

// IsAncestorOf returns true if the given node is a child of this one.
func (n *Node[T]) IsAncestorOf(other *Node[T]) bool {
	if other == nil {
		return false
	}

	return other.IsDescendantOf(n)
}

// IsDescendantOf returns true if this node is a descendant of the given one.
func (n *Node[T]) IsDescendantOf(parent *Node[T]) bool {
	if parent == nil {
		return false
	}

	var node *Node[T] = n
	for node != nil {
		if node == parent {
			return true
		}

		node = node.Parent()
	}

	return false
}
