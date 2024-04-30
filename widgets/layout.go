package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/tree"
)

// Layout define widgets that can have children.
type Layout interface {
	Widget

	FirstChild() Widget
	LastChild() Widget
	AppendChild(Widget) error
	InsertBefore(newChild, reference Widget) error
	RemoveChild(Widget) error
	IsAncestorOf(Widget) bool
	IsDescendantOf(Widget) bool
}

func nodeOrNil(w Widget) *tree.Node[Widget] {
	if w == nil {
		return nil
	}

	return w.Node()
}

func widgetOrNil(n *tree.Node[Widget]) Widget {
	if n == nil {
		return nil
	}

	return n.Unwrap()
}

// BaseLayout define a minimal/incomplete and unopiniated layout implementation
// that should be embedded into more complex implementation.
type BaseLayout struct {
	BaseWidget
	ChildrenLayout ChildrenLayout
}

// NewBaseLayout returns a new base layout.
func NewBaseLayout(embedder Widget) BaseLayout {
	bl := BaseLayout{
		BaseWidget:     NewBaseWidget(embedder),
		ChildrenLayout: ChildrenLayout{},
	}

	return bl
}

// Draw implements draw.Drawer.
func (bl *BaseLayout) Draw(srf draw.Surface) {
	bl.ChildrenLayout.Draw(srf)
}

// FirstChild returns first child of this layout.
func (bl *BaseLayout) FirstChild() Widget {
	return widgetOrNil(bl.node.FirstChild())
}

// LastChild returns last child of this layout.
func (bl *BaseLayout) LastChild() Widget {
	return widgetOrNil(bl.node.LastChild())
}

// AppendChild appends the given widget to the list of child widget. An error is
// returned if the given widget is an ancestor of this node. When an error is
// returned, given widget is left unchanged.
func (bl *BaseLayout) AppendChild(w Widget) error {
	return bl.node.AppendChild(nodeOrNil(w))
}

func (bl *BaseLayout) InsertBefore(newChild, reference Widget) error {
	return bl.node.InsertBefore(nodeOrNil(newChild), nodeOrNil(reference))
}

func (bl *BaseLayout) RemoveChild(child Widget) error {
	return bl.node.RemoveChild(nodeOrNil(child))
}

func (bl *BaseLayout) IsAncestorOf(other Widget) bool {
	return bl.node.IsAncestorOf(nodeOrNil(other))
}

func (bl *BaseLayout) IsDescendantOf(other Widget) bool {
	return bl.node.IsDescendantOf(nodeOrNil(other))
}
