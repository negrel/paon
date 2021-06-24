package tree

var _ Node = Root{}

// Root define the root of a node tree.
type Root struct {
	*node
}

// NewRoot returns a new Root node containing the given data.
func NewRoot(data interface{}) Root {
	return Root{
		node: newNode(data),
	}
}

// Root returns itself to implements the Node interface.
func (r Root) Root() Node {
	return r
}

// SetParent implements the Node interface.
// This method will always panic since root node can't have a parent.
func (r Root) SetParent(Node) {
	panic("root node can't have a parent")
}

// Parent implements the Node interface.
// A nil value is always returned since root node can't have a parent.
func (r Root) Parent() Node {
	return nil
}

// AppendChild implments the Node interface.
func (r Root) AppendChild(newChild Node) error {
	err := r.node.AppendChild(newChild)
	if err != nil {
		return err
	}
	newChild.SetParent(r)

	return nil
}

// InsertBefore implements the Node interface.
func (r Root) InsertBefore(reference, newChild Node) error {
	err := r.node.InsertBefore(reference, newChild)
	if err != nil {
		return err
	}
	newChild.SetParent(r)

	return nil
}
