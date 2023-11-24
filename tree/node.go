package tree

import (
	"errors"
	"fmt"
)

// NodeAccessor define objects tied to a Node.
type NodeAccessor interface {
	Node() *Node
}

// Node define a single Node in a tree.
type Node struct {
	value any

	parent *Node

	// Siblings
	next     *Node
	previous *Node

	// Children
	firstChild *Node
	lastChild  *Node
}

// NewNode returns a new Node containing the given data.
func NewNode(data any) *Node {
	return &Node{
		value:      data,
		parent:     nil,
		next:       nil,
		previous:   nil,
		firstChild: nil,
		lastChild:  nil,
	}
}

// Node implements NodeAccessor.
func (n *Node) Node() *Node {
	return n
}

// Unwrap returns the underlying data contained in the node.
func (n *Node) Unwrap() any {
	return n.value
}

// Swap swaps underlying data contained in the node and returns old one.
func (n *Node) Swap(v any) any {
	old := n.value
	n.value = v
	return old
}

// NextSibling returns the next sibling of this node.
func (n *Node) Next() *Node {
	return n.next
}

// Previous implements Node.
func (n *Node) Previous() *Node {
	return n.previous
}

// Parent returns parent node.
func (n *Node) Parent() *Node {
	return n.parent
}

// Root returns root node of the tree.
func (n *Node) Root() *Node {
	if n.parent != nil {
		return n.parent.Root()
	}

	return n
}

// FirstChild returns first direct child of this node.
func (n *Node) FirstChild() *Node {
	return n.firstChild
}

// LastChild returns last direct child of this node.
func (n *Node) LastChild() *Node {
	return n.lastChild
}

// Appends the given child to the list of child node. An error is returned
// if the given child is an ancestor of this node.
// When an error is returned, the child is left unchanged
func (n *Node) AppendChild(newChild *Node) (err error) {
	if err = n.ensurePreInsertionValidity(newChild); err != nil {
		return fmt.Errorf("can't append child, %v", err)
	}
	n.appendChild(newChild)

	return nil
}

func (n *Node) appendChild(newChild *Node) {
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

func (n *Node) ensurePreInsertionValidity(child *Node) error {
	if child == nil {
		return errors.New("child is nil")
	}

	// check if child is not a parentNode of pn
	if n.IsDescendantOf(child) {
		return errors.New("child contains the parentNode")
	}

	return nil
}

func (n *Node) prepareChildForInsertion(newChild *Node) {
	if parent := newChild.Parent(); parent != nil {
		_ = parent.RemoveChild(newChild)
	}
}

// Inserts the given child before the given reference child Node. If the
// reference is nil, the child is appended. An error is returned
// if the given child is an ancestor of this node or if the reference
// is not a direct child of this.
// When an error is returned, the child is left unchanged
func (n *Node) InsertBefore(newChild, reference *Node) error {
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

func (n *Node) insertBefore(reference, newChild *Node) {
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
func (n *Node) RemoveChild(child *Node) error {
	// if not a child
	if n != child.Parent() {
		return errors.New("can't remove child, the node is not a child of this node")
	}

	n.removeChild(child)

	return nil
}

func (n *Node) removeChild(child *Node) {
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
func (n *Node) IsAncestorOf(other *Node) bool {
	if other == nil {
		return false
	}

	return other.IsDescendantOf(n)
}

// IsDescendantOf returns true if this node is a descendant of the given one.
func (n *Node) IsDescendantOf(parent *Node) bool {
	if parent == nil {
		return false
	}

	var node *Node = n
	for node != nil {
		if node == parent {
			return true
		}

		node = node.Parent()
	}

	return false
}
