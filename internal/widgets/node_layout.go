package widgets

var _ Layout = &NodeLayout{}

// NodeLayout define a minimal branch node in the Widget tree.
type NodeLayout struct {
	*Node
	*widgetList
}

// Children return the children of the layout.
func (nl *NodeLayout) Children() []Widget {
	return nl.list[:]
}

// NewNodeLayout return a new NodeLayout object to embed in custom
// layout widget.
func NewNodeLayout(name string) *NodeLayout {
	return &NodeLayout{
		Node:       NewNodeWidget(name),
		widgetList: NewWidgetList(),
	}
}
