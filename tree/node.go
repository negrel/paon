package tree

// Node is a node in the tree
type Node struct {
	parent      *Node
	id          string
	classList   []string
	isConnected bool
}

func newNode() {
	return &Node{}
}

func (n *Node) isRoot() bool {
	return n.parent == nil && n.isConnected
}

func AppendChild()
