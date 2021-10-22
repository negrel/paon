package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/tree"
	treevents "github.com/negrel/paon/pdk/tree/events"
	"github.com/negrel/paon/styles"
)

type baseWidgetOption struct {
	*BaseWidget

	nodeOptions  []treevents.NodeOption
	defaultStyle styles.Style
}

// WidgetOption define an option for BaseWidget.
type WidgetOption func(*baseWidgetOption)

// NodeConstructor returns a WidgetOption that sets the internal tree.Node constructor used by the
// Widget.
func NodeConstructor(constructor func(data interface{}) tree.Node) WidgetOption {
	assert.NotNil(constructor)

	return func(bwo *baseWidgetOption) {
		bwo.nodeOptions = append(bwo.nodeOptions, treevents.NodeConstructor(constructor))
	}
}

// Wrap returns a WidgetOption that sets the internal data used by the tree.Node.
// This data is accessible through the tree.Node.Unwrap method.
// This options should only be used by structs that embed a BaseWidget.
func Wrap(data Widget) WidgetOption {
	assert.NotNil(data)

	return func(bwo *baseWidgetOption) {
		bwo.nodeOptions = append(bwo.nodeOptions, treevents.Wrap(data))
	}
}

// Target returns a WidgetOption that sets the internal events.Target object
// used by the Widget. This can be useful if you want to add listener before the Widget default
// one.
func Target(target events.Target) WidgetOption {
	assert.NotNil(target)

	return func(bwo *baseWidgetOption) {
		bwo.nodeOptions = append(bwo.nodeOptions, treevents.EventTarget(target))
	}
}

// DefaultStyle returns a WidgetOption that sets default style properties of
// the widget.
func DefaultStyle(style styles.Style) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.defaultStyle = style
	}
}
