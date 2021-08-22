package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/events"
)

var redrawEventType = events.NewType("redraw")

// RedrawEventType returns the events.Type of RedrawEvent events.
func RedrawEventType() events.Type {
	return redrawEventType
}

// RedrawListener returns an events.Listener that will call the given handler
// on redraw events.
func RedrawListener(handler func(RedrawEvent)) *events.Listener {
	return &events.Listener{
		Type: redrawEventType,
		Handle: func(event events.Event) {
			assert.IsType(event, NewRedrawEvent(nil))
			handler(event.(RedrawEvent))
		},
	}
}

var _ events.Event = RedrawEvent{}

// RedrawEvent is triggered when a Widget need to be redrawn.
type RedrawEvent struct {
	events.Event

	Drawer draw.Drawer
}

// NewRedrawEvent returns a new RedrawEvent.
func NewRedrawEvent(drawer draw.Drawer) RedrawEvent {
	return RedrawEvent{
		Event:  events.NewEvent(redrawEventType),
		Drawer: drawer,
	}
}
