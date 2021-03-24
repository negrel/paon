package tree

import "github.com/negrel/paon/pkg/pdk/id"

// SetParentOf is an exported function that sets the parent of a given Node.
// This function should be used to change the ParentNode of a Node after a call
// to an insertion method if it didn't return an error.
func SetParentOf(node Node, parent ParentNode) {
	node.setParentNode(parent)
}

// Node define an element in a Node tree.
type Node interface {
	id.Identifiable

	// Return true if the given Node is the same as this.
	IsSame(Node) bool

	// Next sibling.
	NextNode() Node
	setNextNode(Node)

	// Previous sibling.
	PreviousNode() Node
	setPreviousNode(Node)

	// ParentNode is the direct parentNode of the Node.
	ParentNode() ParentNode
	setParentNode(ParentNode)

	// RootNode define the root of the Node tree.
	RootNode() ParentNode

	// IsDescendantOf return true if this is a descendant of the given Node.
	IsDescendantOf(node Node) bool
}

var _ Node = &node{}

type node struct {
	id id.ID

	next     Node
	previous Node
	parent   ParentNode
}

func NewNode() Node {
	return newNode()
}

func newNode() *node {
	return &node{
		id: id.Make(),
	}
}

func (n *node) ID() id.ID {
	return n.id
}

func (n *node) IsSame(other Node) bool {
	return n.ID() == other.ID()
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

		node = node.ParentNode()
	}

	return false
}

func (n *node) NextNode() Node {
	return n.next
}

func (n *node) setNextNode(next Node) {
	n.next = next
}

func (n *node) PreviousNode() Node {
	return n.previous
}

func (n *node) setPreviousNode(previous Node) {
	n.previous = previous
}

func (n *node) ParentNode() ParentNode {
	return n.parent
}

func (n *node) setParentNode(parent ParentNode) {
	n.parent = parent
}

func (n *node) RootNode() ParentNode {
	if n.parent != nil {
		return n.parent.RootNode()
	}

	return nil
}

// String implements the fmt.Stringer interface.
func (n *node) String() string {
	return string(n.id)
}
