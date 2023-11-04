package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/id"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
	treevents "github.com/negrel/paon/tree/events"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// Any types that implement the Widget interface can be added to the widget tree. However, it is strongly
// recommended to create custom widgets using the BaseWidget implementation.
type Widget interface {
	id.Identifiable
	events.Target
	layout.Layout
	draw.Drawer
	styles.Styled

	// Node returns the underlying Node used by this Widget.
	Node() treevents.Node
}

var _ Widget = &BaseWidget{}

// Private treevents.Node type that can be embedded in private field.
type node treevents.Node

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
type BaseWidget struct {
	node

	layout layout.Layout
	drawer draw.Drawer

	style styles.Style
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
func (bw *BaseWidget) Layout(co layout.Constraint) geometry.Size {
	return bw.layout.Layout(co)
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

// Node implements the Widget interface.
func (bw *BaseWidget) Node() treevents.Node {
	return bw.node
}

// Style implements the styles.Styled interface.
func (bw *BaseWidget) Style() *styles.Style {
	return &bw.style
}
