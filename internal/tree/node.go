package tree

import (
	"errors"

	"github.com/negrel/paon/pkg/pdk/id"
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

	// RootNode define the root of the Node tree.
	Root() Node

	// Return the first child Node of this.
	FirstChild() Node

	// Return the last child Node of this.
	LastChild() Node

	// Append the given child to the list of child Node. An error is returned
	// if the given child is an ancestor of this Node.
	AppendChild(newChild Node) error

	// Insert the given child before the given reference child Node. If the
	// reference is nil, the child is appended. An error is returned
	// if the given child is an ancestor of this Node or if the reference
	// is not a direct child of this.
	InsertBefore(reference, newChild Node) error

	// Remove the given direct child Node of this. Return an error otherwise.
	RemoveChild(child Node) error

	// IsAncestorOf return true if the given Node is a child of this.
	IsAncestorOf(child Node) bool

	// IsDescendantOf return true if this is a descendant of the given Node.
	IsDescendantOf(node Node) bool
}

var _ Node = &node{}

type node struct {
	id   id.ID
	data interface{}

	next     Node
	previous Node
	parent   Node
}

func NewLeafNode(data interface{}) Node {
	return newLeafNode(data)
}

func newLeafNode(data interface{}) *node {
	return &node{
		id:   id.Make(),
		data: data,
	}
}

func (n *node) Unwrap() interface{} {
	return n.data
}

func (n *node) ID() id.ID {
	return n.id
}

func (n *node) IsSame(other Node) bool {
	if other == nil {
		return false
	}

	return n.ID() == other.ID()
}

func (n *node) Next() Node {
	return n.next
}

func (n *node) SetNext(next Node) {
	n.next = next
}

func (n *node) Previous() Node {
	return n.previous
}

func (n *node) SetPrevious(previous Node) {
	n.previous = previous
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) SetParent(parent Node) {
	n.parent = parent
}

func (n *node) Root() Node {
	if n.parent != nil {
		return n.parent.Root()
	}

	return nil
}

func (n *node) FirstChild() Node {
	return nil
}

func (n *node) LastChild() Node {
	return nil
}

func (n *node) AppendChild(_ Node) error {
	return errors.New("simple node can't have children")
}

func (n *node) InsertBefore(_, _ Node) error {
	return errors.New("simple node can't have children")
}

func (n *node) RemoveChild(Node) error {
	return errors.New("simple node can't have children")
}

func (n *node) IsAncestorOf(other Node) bool {
	if other == nil {
		return false
	}

	return other.IsDescendantOf(n)
}

func (n *node) IsDescendantOf(parent Node) bool {
	if parent == nil {
		return false
	}

	var node Node = n
	for node != nil {
		if node.IsSame(parent) {
			return true
		}

		node = node.Parent()
	}

	return false
}

// String implements the fmt.Stringer interface.
func (n *node) String() string {
	return string(n.id)
}
