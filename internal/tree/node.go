package tree

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/negrel/debuggo/pkg/log"
)

// Node define an element in a Node tree.
type Node interface {
	fmt.Stringer

	ID() uuid.UUID
	Name() string

	IsSame(Node) bool

	// Next sibling.
	Next() Node
	setNext(Node)

	// Previous sibling.
	Previous() Node
	setPrevious(Node)

	// ParentNode is the direct parentNode of the Node.
	Parent() ParentNode
	setParent(ParentNode)

	// Root define the root of the Node tree.
	Root() Root
	setRoot(Root)

	isDescendantOf(node Node) bool

	// Whether the Node is connected to an active Node tree.
	isConnected() bool
}

var _ Node = &node{}

type node struct {
	name string
	id   uuid.UUID

	next     Node
	previous Node
	parent   ParentNode
	root     Root
}

func NewNode(name string) Node {
	return newNode(name)
}

func newNode(name string) *node {
	log.Debugln("creating a", name, "node")

	return &node{
		name: name,
		id:   uuid.New(),
	}
}

func (n *node) String() string {
	return fmt.Sprintf("%v-%v", n.name, n.id)
}

func (n *node) ID() uuid.UUID {
	return n.id
}

func (n *node) Name() string {
	return n.name
}

func (n *node) IsSame(other Node) bool {
	return n.ID() == other.ID()
}

func (n *node) isDescendantOf(parent Node) bool {
	var node Node = n
	for node != nil {
		if node.IsSame(parent) {
			return true
		}

		node = node.Parent()
	}

	return false
}

func (n *node) Next() Node {
	return n.next
}

func (n *node) setNext(next Node) {
	n.next = next
}

func (n *node) Previous() Node {
	return n.previous
}

func (n *node) setPrevious(previous Node) {
	n.previous = previous
}

func (n *node) Parent() ParentNode {
	return n.parent
}

func (n *node) setParent(parent ParentNode) {
	n.parent = parent
}

func (n *node) Root() Root {
	return n.root
}

func (n *node) setRoot(root Root) {
	n.root = root
}
func (n *node) isConnected() bool {
	return n.root != nil
}
