package tree

// Node define an element in a Node tree.
type Node interface {
	nodeID() NodeID

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
	RootNode() Root

	IsDescendantOf(node Node) bool
}

var _ Node = &node{}

type node struct {
	cache

	id NodeID

	next     Node
	previous Node
	parent   ParentNode
}

func NewNode() Node {
	return newNode()
}

func newNode() *node {
	return &node{
		id: makeNodeID(),
	}
}

func (n *node) nodeID() NodeID {
	return n.id
}

func (n *node) IsSame(other Node) bool {
	return n.nodeID() == other.nodeID()
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
	n.clearCache()
	n.parent = parent
}

func (n *node) clearCache() {
	n.cache = cache{}
}

func (n *node) RootNode() Root {
	var root Root

	if n.cache.root != nil {
		root = n.cache.root
	} else {
		root = n.rootNode()
		n.cache.root = root
	}

	return root
}

func (n *node) rootNode() Root {
	if n.parent != nil {
		return n.parent.RootNode()
	}

	return nil
}
