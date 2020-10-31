package widgets

import (
	"github.com/negrel/paon/internal/events"
)

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

// DispatchEvent implements the events.EventTarget interface.
func (nl *NodeLayout) DispatchEvent(event events.Event) {
	nl.Node.Target.DispatchEvent(event)

	for _, widget := range nl.widgetList.list {
		widget.DispatchEvent(event)
	}
}
