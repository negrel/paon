package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
)

type baseWidgetOption struct {
	*BaseWidget

	data      Widget
	listeners []events.Listener
}

// WidgetOption define an option for BaseWidget.
type WidgetOption func(*baseWidgetOption)

// Wrap returns a WidgetOption that sets the internal data contained in tree.Node.
// This data is accessible throught the tree.Node.Unwrap method.
// This options should be used by structs that embed BaseWidget.
func Wrap(data Widget) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.data = data
	}
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
		bwo.layout = l
	}
}

// LayoutFunc returns a WidgetOption that define the layout of the widget.
func LayoutFunc(l func(layout.Constraint) geometry.Size) WidgetOption {
	return LayoutLayout(layout.LayoutFunc(l))
}

// Style returns a WidgetOption that define the widget style.
func Style(s styles.Style) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.style = s
	}
}
