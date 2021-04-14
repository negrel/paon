package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
)

var _TypeMouseMove = events.MakeType("mouse-move")

func TypeMouseMove() events.Type {
	return _TypeMouseMove
}

// MouseMoveListener convert the given event handler as a generic Listener.
func MouseMoveListener(handler func(MouseMove)) *events.Listener {
	l := events.Listener{
		Type: _TypeMouseMove,
		Handle: func(event events.Event) {
			assert.IsType(event, MakeMouseMove(geometry.Point{}))
			handler(event.(MouseMove))
		},
	}

	return &l
}

var _ events.Event = MouseMove{}

// MouseMove is triggered when the user click inside the rendering surface.
type MouseMove struct {
	events.Event
	Position geometry.Point
}

// MakeMouseMove returns a new MouseMove events.Event.
func MakeMouseMove(position geometry.Point) MouseMove {
	return MouseMove{
		Event:    events.MakeEvent(_TypeMouseMove),
		Position: position,
	}
}
