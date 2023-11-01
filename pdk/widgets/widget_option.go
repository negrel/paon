package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

type baseWidgetOption struct {
	*BaseWidget

	data        interface{}
	listeners   []events.Listener
	nodeOptions []treevents.NodeOption
}

// WidgetOption define an option for BaseWidget.
type WidgetOption func(*baseWidgetOption)

func initialLCS(lcs treevents.LifeCycleStage) WidgetOption {
	return func(bwo *baseWidgetOption) {
		bwo.stage = lcs
	}
}

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

// Target returns a WidgetOption that sets the internal events.Target object
// used by the Widget. This can be useful if you want to add listener before the Widget default
// one.
func Target(target events.Target) WidgetOption {
	assert.NotNil(target)

	return func(bwo *baseWidgetOption) {
		bwo.Target = target
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
