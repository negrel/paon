package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/layout"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

type baseWidgetOption struct {
	*BaseWidget

	listeners   []events.Listener
	nodeOptions []treevents.NodeOption
}

// WidgetOption define an option for BaseWidget.
type WidgetOption func(*baseWidgetOption)

// NodeOptions adds the given NodeOptions to options list used to create underlying
// node.
func NodeOptions(options ...treevents.NodeOption) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.nodeOptions = append(bwo.nodeOptions, options...)
	}
}

// Wrap returns a WidgetOption that sets the internal data used by the tree.Node.
// This data is accessible throught the tree.Node.Unwrap method.
// This options should only be used by structs that embed a BaseWidget.
func Wrap(data Widget) WidgetOption {
	assert.NotNil(data)

	return NodeOptions(treevents.Wrap(data))
}

// Listener returns a WidgetOption that append the given listener to the internal
// events.Target.
func Listener(etype events.Type, h events.Handler) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.listeners = append(bwo.listeners, events.Listener{
			EventType: etype,
			Handler:   h,
		})
	}
}

// Drawer returns a WidgetOption that define the drawer of the widget.
func Drawer(drawer draw.Drawer) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.drawer = drawer
	}
}

// DrawerFunc returns a WidgetOption that define the drawer of the widget.
func DrawerFunc(d func(_ draw.Surface)) WidgetOption {
	return Drawer(draw.DrawerFunc(d))
}

// LayoutLayout returns a WidgetOption that define the layout of the widget.
func LayoutLayout(l layout.Layout) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.layout = layout.NewCache(l)
	}
}

// LayoutFunc returns a WidgetOption that define the layout of the widget.
func LayoutFunc(l func(layout.Constraint) geometry.Size) WidgetOption {
	return LayoutLayout(layout.LayoutFunc(l))
}
