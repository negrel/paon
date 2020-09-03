package widgets

var _ Layout = &NodeLayout{}

// NodeLayout define a minimal branch node in the Widget tree.
type NodeLayout struct {
	*Node
	children []Widget
}

// NewNodeLayout return a new NodeLayout object to embed in custom
// layout widget.
func NewNodeLayout(name string) *NodeLayout {
	return &NodeLayout{
		Node:     NewNodeWidget(name),
		children: []Widget{},
	}
}

// AppendChild implements the Layout interface.
func (c NodeLayout) AppendChild(child Widget) {
	panic("implement me")
}

// Children implements the Layout interface.
func (c NodeLayout) Children() []Widget {
	return c.children
}

// IndexOf implements the Layout interface.
func (c NodeLayout) IndexOf(child Widget) int {
	for i, c := range c.children {
		if c == child {
			return i
		}
	}

	return -1
}

// RemoveChild implements the Layout interface.
func (c NodeLayout) RemoveChild(child Widget) {
	panic("implement me")
}
