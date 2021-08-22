package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/layout"
)

var reflowEventType = events.NewType("reflow")

// ReflowEventType returns the events.Type of ReflowEvent events.
func ReflowEventType() events.Type {
	return reflowEventType
}

// ReflowListener returns an events.Listener that will call the given handler
// on reflow events.
func ReflowListener(handler func(ReflowEvent)) *events.Listener {
	return &events.Listener{
		Type: reflowEventType,
		Handle: func(event events.Event) {
			assert.IsType(event, NewReflowEvent(nil))
			handler(event.(ReflowEvent))
		},
	}
}

var _ events.Event = ReflowEvent{}

// ReflowEvent is triggered when a Widget need a layout.
type ReflowEvent struct {
	events.Event
	Manager layout.Manager
}

// NewReflowEvent returns a new ReflowEvent.
func NewReflowEvent(manager layout.Manager) ReflowEvent {
	return ReflowEvent{
		Event:   events.NewEvent(reflowEventType),
		Manager: manager,
	}
}
