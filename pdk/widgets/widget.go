package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/tree"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// Any types that implement the Widget interface can be added to the widget tree. However, it is strongly
// recommended to create custom widgets using the BaseWidget implementation.
type Widget interface {
	id.Identifiable
	events.Target
	layout.Layout
	draw.Drawer

	// LifeCycleStage returns the current LifeCycleStage of this Widget.
	LifeCycleStage() treevents.LifeCycleStage

	// Node returns the underlying tree.Node used by this Widget.
	Node() treevents.Node

	// Parent returns the parent Layout of this Node.
	Parent() Layout

	// Root returns the root of the widget tree.
	Root() *Root

	// NextSibling returns the next sibling widget in the parent child list.
	NextSibling() Widget

	// PreviousSibling returns the previous sibling widget in the parent child list.
	PreviousSibling() Widget
}

func widgetOrNil(n tree.Node) Widget {
	if n == nil {
		return nil
	}

	return n.Unwrap().(Widget)
}

func nodeOrNil(w Widget) treevents.Node {
	if w == nil {
		return nil
	}

	return w.Node()
}

var _ Widget = &BaseWidget{}

// Private treevents.Node type that can be embedded in private field.
type node treevents.Node

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
// BaseWidget takes care of the following things for you:
// - Constant access time to the root.
// - Caching the layout.BoxModel
// - Limit the draw area of the context to the widget border box.
type BaseWidget struct {
	node

	layout layout.Layout
	drawer draw.Drawer
}

// NewBaseWidget returns a new BaseWidget object configured with
// the given options.
// The LayoutAlgo and Drawer widget options are required.
// To embed this layout in composite struct, use the Wrap widget options.
func NewBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := newBaseWidget(options...)

	return widget
}

func newBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := &BaseWidget{}
	widgetConf := &baseWidgetOption{
		BaseWidget:  widget,
		nodeOptions: []treevents.NodeOption{treevents.Wrap(widget)},
	}

	for _, option := range options {
		option(widgetConf)
	}

	widget.node = treevents.NewBaseNode(widgetConf.nodeOptions...)

	for _, listener := range widgetConf.listeners {
		widget.AddEventListener(listener.EventType, listener.Handler)
	}

	return widget
}

// Layout implements Layout.
func (bw *BaseWidget) Layout(co layout.Constraint) geometry.Rectangle {
	rect := bw.layout.Layout(co)
	return rect
}

// Draw implements Drawer.
func (bw *BaseWidget) Draw(surface draw.Surface) {
	bw.drawer.Draw(surface)
}

// ID implements the id.Identifiable interface.
func (bw *BaseWidget) ID() id.ID {
	return bw.node.ID()
}

// IsSame implements the id.Identifiable interface.
func (bw *BaseWidget) IsSame(other id.Identifiable) bool {
	return bw.node.IsSame(other)
}

// Widget is an helper method that returns the Widget wrapped by the internal tree.Node.
// This method act kinda like the `this` keyword in Java.
func (bw *BaseWidget) Widget() Widget {
	return widgetOrNil(bw.node)
}

// Node implements the Widget interface.
func (bw *BaseWidget) Node() treevents.Node {
	return bw.node
}

// Parent implements the Widget interface.
func (bw *BaseWidget) Parent() Layout {
	if parent := bw.node.Parent(); parent != nil {
		return parent.Unwrap().(Layout)
	}

	return nil
}

// Root implements the Widget interface.
func (bw *BaseWidget) Root() *Root {
	return bw.node.Root().Unwrap().(*Root)
}

// NextSibling implements the Widget interface.
func (bw *BaseWidget) NextSibling() Widget {
	return widgetOrNil(bw.node.Next())
}

// PreviousSibling implements the Widget interface.
func (bw *BaseWidget) PreviousSibling() Widget {
	return widgetOrNil(bw.node.Previous())
}
