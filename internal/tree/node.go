package tree

type NodeType int8

const (
	ElementNode NodeType = iota + 1
	TextNode
	DocumentNode
)

// Node are part of the Node tree that compose the Document.
type Node interface {
	Owner() *Document
	setOwner(*Document)

	Parent() ParentNode
	setParent(ParentNode)

	Previous() Node
	setPrevious(Node)

	Next() Node
	setNext(Node)

	Type() NodeType

	isConnected() bool
	isParentNode() bool
	contains(Node) bool

	isAncestorOf(Node) bool
	isDescendantOf(ParentNode) bool

	isSame(Node) bool
}

var _ Node = &node{}

type node struct {
	owner       **Document
	parent      ParentNode
	previous    Node
	next        Node
	nodeType    NodeType
	isContainer bool
}

func (n *node) Owner() *Document {
	if n.owner != nil {
		return *n.owner
	}

	return nil
}

func (n *node) setOwner(owner *Document) {
	n.owner = &owner
}

func (n *node) Parent() ParentNode {
	return n.parent
}

func (n *node) setParent(parent ParentNode) {
	n.parent = parent
}
func (n *node) Previous() Node {
	return n.previous
}

func (n *node) setPrevious(previous Node) {
	n.previous = previous
}

func (n *node) Next() Node {
	return n.next
}

func (n *node) setNext(next Node) {
	n.next = next
}

func (n *node) Type() NodeType {
	return n.nodeType
}

func (n *node) isConnected() bool {
	return n.owner != nil
}

func (n *node) isParentNode() bool {
	return n.isContainer
}

func (n *node) contains(node Node) bool {
	if node == nil || !n.isContainer {
		return false
	}

	return n == node || n.isAncestorOf(node)
}

func (n *node) isAncestorOf(node Node) bool {
	if node == nil || !n.isContainer {
		return false
	}
	if n.isConnected() != node.isConnected() || node.Owner() != n.Owner() {
		return false
	}

	for ancestor := node.Parent(); ancestor != nil; ancestor = node.Parent() {
		if ancestor.isSame(n) {
			return true
		}
	}

	return false
}

func (n *node) isDescendantOf(node ParentNode) bool {
	return node.isAncestorOf(n)
}

func (n *node) isSame(other Node) bool {
	return n == other
}
