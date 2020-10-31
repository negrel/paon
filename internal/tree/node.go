package tree

type NodeType int8

const (
	ElementNode NodeType = iota + 1
	TextNode
	DocumentNode
)

type Node interface {
	Owner() Node
	setOwner(Node)

	Parent() ContainerNode
	setParent(ContainerNode)

	Previous() Node
	setPrevious(Node)

	Next() Node
	setNext(Node)

	Type() NodeType

	isConnected() bool
	isContainerNode() bool
	contains(Node) bool

	isAncestorOf(Node) bool
	isDescendantOf(ContainerNode) bool

	isSame(Node) bool
}

var _ Node = &node{}

type node struct {
	owner       Node
	parent      ContainerNode
	previous    Node
	next        Node
	nType       NodeType
	isContainer bool
}

func (n *node) Owner() Node {
	return n.owner
}

func (n *node) setOwner(owner Node) {
	n.owner = owner
}

func (n *node) Parent() ContainerNode {
	return n.parent
}

func (n *node) setParent(parent ContainerNode) {
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
	return n.nType
}

func (n *node) isConnected() bool {
	return n.owner != nil
}

func (n *node) isContainerNode() bool {
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

func (n *node) isDescendantOf(node ContainerNode) bool {
	return node.isAncestorOf(n)
}

func (n *node) isSame(other Node) bool {
	return n == other
}
