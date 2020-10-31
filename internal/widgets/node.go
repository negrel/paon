package widgets

type Node interface {
	Widget

	// Next sibling.
	Next() Node
	setNext(Node)

	// Previous sibling.
	Previous() Node
	setPrevious(Node)

	// Parent is the direct parent of the Node.
	Parent() Parent
	setParent(Parent)

	// Root define the root of the Node tree.
	Root() Parent
	setRoot(Parent)

	isDescendantOf(node Node) bool

	// Whether the Node is connected to an active Node tree.
	isConnected() bool
}

var _ Node = &node{}

type node struct {
	*widget

	next     Node
	previous Node
	parent   Parent
	root     Parent
}

func (n node) isDescendantOf(parent Node) bool {
	var node Node = n
	for node != nil {
		if node.isSame(parent) {
			return true
		}

		node = node.Parent()
	}

	return false
}

func newNode(name string) *node {
	return &node{
		widget: newWidget(name),
	}
}

func (n node) Next() Node {
	return n.next
}

func (n node) setNext(next Node) {
	n.next = next
}

func (n node) Previous() Node {
	return n.previous
}

func (n node) setPrevious(previous Node) {
	n.previous = previous
}

func (n node) Parent() Parent {
	return n.parent
}

func (n node) setParent(parent Parent) {
	n.parent = parent
}

func (n node) Root() Parent {
	return n.root
}

func (n node) setRoot(root Parent) {
	n.root = root
}
func (n node) isConnected() bool {
	return n.root != nil
}
