package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/tree"
)

// Widget define objects part of widget tree and UI.
type Widget interface {
	events.Target
	Renderable

	// Node associated to the widget. This is mostly for internal use,
	// applications code shouldn't rely on this and use widget base method instead
	// (NextSibling, PreviousSibling, etc).
	Node() *tree.Node[Widget]

	Root() *Root
	Parent() Widget
	NextSibling() Widget
	PreviousSibling() Widget
}

// IsMounted returns true if widgets root is non nil.
func IsMounted(w Widget) bool {
	return w.Root() != nil
}

// BaseWidget define a minimal/incomplete widget implementation that should be
// embedded into more complex implementation.
type BaseWidget struct {
	events.Target
	node    tree.Node[Widget]
	IsDirty bool
}

// NewBaseWidget returns a new base widget.
func NewBaseWidget(embedder Widget) BaseWidget {
	bw := BaseWidget{
		Target: events.NewTarget(),
	}
	bw.node = tree.NewNode(embedder)

	return bw
}

// Node returns internal node.
func (bw *BaseWidget) Node() *tree.Node[Widget] {
	return &bw.node
}

// Parent returns parent widget.
func (bw *BaseWidget) Parent() Widget {
	return widgetOrNil(bw.node.Parent())
}

// Root returns root of widget tree. Nil is returned if widget is not mounted.
func (bw *BaseWidget) Root() *Root {
	w := widgetOrNil(bw.node.Root())
	if root, isRoot := w.(*Root); isRoot {
		return root
	}

	return nil
}

// NextSibling returns next sibling widget if any.
func (bw *BaseWidget) NextSibling() Widget {
	return widgetOrNil(bw.node.Next())
}

// PreviousSibling returns previous sibling widget if any.
func (bw *BaseWidget) PreviousSibling() Widget {
	return widgetOrNil(bw.node.Previous())
}

// DispatchEvent implements events.Target.
func (bw *BaseWidget) DispatchEvent(ev events.Event) {
	bw.Target.DispatchEvent(ev.WithTarget(bw.node.Unwrap()))
}

// NeedRender notify runtime that this widget need to be rendered again.
func (bw *BaseWidget) NeedRender() {
	root := bw.Root()
	if root != nil {
		root.DispatchEvent(events.NewEvent(NeedRenderEventType, nil))
	}
}
