package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/tree"
	"github.com/negrel/paon/styles"
)

type baseWidgetOption struct {
	*BaseWidget

	nodeConstructor func(data interface{}) tree.Node
	data            interface{}
	listeners       []*events.Listener
}

// WidgetOption define an option for BaseWidget.
type WidgetOption func(*baseWidgetOption)

func initialLCS(lcs LifeCycleStage) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.stage = lcs
	}
}

// NodeConstructor returns a WidgetOption that sets the internal tree.Node constructor used by the
// Widget.
func NodeConstructor(constructor func(data interface{}) tree.Node) WidgetOption {
	assert.NotNil(constructor)

	return func(bwo *baseWidgetOption) {
		bwo.nodeConstructor = constructor
	}
}

// Wrap returns a WidgetOption that sets the internal data used by the tree.Node.
// This data is accessible throught the tree.Node.Unwrap method.
// This options should only be used by structs that embed a BaseWidget.
func Wrap(data Widget) WidgetOption {
	assert.NotNil(data)

	return func(bwo *baseWidgetOption) {
		bwo.data = data
	}
}

// Target returns a WidgetOption that sets the internal events.Target object
// used by the Widget. This can be useful if you want to add listener before the Widget default
// one.
func Target(target events.Target) WidgetOption {
	assert.NotNil(target)

	return func(bwo *baseWidgetOption) {
		bwo.Target = target
	}
}

// Listeners returns a WidgetOption that prepends the given listeners to the internal
// events.Target.
func Listeners(listeners ...*events.Listener) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.listeners = append(bwo.listeners, listeners...)
	}
}

// LayoutManager returns a WidgetOption that sets the layout algorithm
// used to compute the flow of the widget.
func LayoutManager(man layout.Manager) WidgetOption {
	assert.NotNil(man)

	return func(bwo *baseWidgetOption) {
		bwo.cache = layout.NewCache(man)
	}
}

// Drawer returns a WidgetOption that sets the drawer used to
// draw the Widget on a Canvas.
func Drawer(drawer draw.Drawer) WidgetOption {
	assert.NotNil(drawer)

	return func(bwo *baseWidgetOption) {
		bwo.drawer = drawer
	}
}

// Theme returns a WidgetOption that sets the Theme used by the
// Widget.
func Theme(theme styles.Theme) WidgetOption {
	assert.NotNil(theme)

	return func(bwo *baseWidgetOption) {
		bwo.theme = theme
	}
}
