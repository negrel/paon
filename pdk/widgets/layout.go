package widgets

import (
	"github.com/negrel/paon/pdk/tree"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

// Layout is an extension of the Widget interface that adds the support for
// children widgets. Howerver, it is strongly recommended to create custom
// layouts using the BaseLayout implementation.
type Layout interface {
	Widget

	// AppendChild appends the given Widget to this Layout
	// children list.
	// An error is returned if the child to insert is an ancestor of this Layout.
	AppendChild(Widget) error

	// InsertBefore inserts the given Widget before the given reference.
	// If the reference is nil, the child is appended.
	// An error is returned if the child to insert is an ancestor of this Layout.
	// If the child have a parent, the Widget is removed from it.
	InsertBefore(reference, newChild Widget) error

	// RemoveChild removes the given direct child of this Layout.
	// An error is returned if the child is nil or not a direct
	// child of this.
	RemoveChild(Widget) error

	// FirstChild returns the first child in the children list
	// of this Layout.
	FirstChild() Widget

	// LastChild returns the last child in the children list
	// of this Layout.
	LastChild() Widget
}

var _ Layout = &BaseLayout{}

// BaseLayout define a basic implementation of the Layout interface.
// BaseLayout can either be used alone (see NewBaseLayout for the required options)
// or in composite struct.
// BaseLayout is an extension of BaseWidget.
// BaseLayout takes care of the following things for you:
// - Dispatch LifeCycleEvents to children
type BaseLayout struct {
	*BaseWidget
}

// NewBaseLayout returns a new BaseLayout object configured with the given
// options.
// When embedding this layout in a composite struct, wrap the struct instance
// using Wrap widget options.
func NewBaseLayout(options ...LayoutOption) *BaseLayout {
	layout := newBaseLayout(options...)

	return layout
}

func newBaseLayout(options ...LayoutOption) *BaseLayout {
	layout := &BaseLayout{}
	layoutConf := &baseLayoutOption{
		BaseLayout:        layout,
		widgetConstructor: NewBaseWidget,
		widgetOptions: []WidgetOption{
			Wrap(layout),
			NodeOptions(treevents.NodeConstructor(tree.NewNode)),
		},
	}

	for _, option := range options {
		option(layoutConf)
	}

	layout.BaseWidget = layoutConf.widgetConstructor(layoutConf.widgetOptions...)

	return layout
}

// FirstChild implements the Layout interface.
func (bl *BaseLayout) FirstChild() Widget {
	return widgetOrNil(bl.node.FirstChild())
}

// LastChild implements the Layout interface.
func (bl *BaseLayout) LastChild() Widget {
	return widgetOrNil(bl.node.LastChild())
}

// AppendChild implements the Layout interface.
func (bl *BaseLayout) AppendChild(newChild Widget) error {
	return bl.node.AppendChild(newChild.Node())
}

// InsertBefore implements the Layout interface.
func (bl *BaseLayout) InsertBefore(reference, newChild Widget) error {
	return bl.node.InsertBefore(nodeOrNil(reference), nodeOrNil(newChild))
}

// RemoveChild implements the Layout interface.
func (bl *BaseLayout) RemoveChild(child Widget) error {
	return bl.node.RemoveChild(nodeOrNil(child))
}
