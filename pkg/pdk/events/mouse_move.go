package events

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
)

// MouseMoveListener convert the given event handler as a generic Listener.
func MouseMoveListener(handler func(MouseMove)) *Listener {
	l := Listener{
		Type: TypeMouseMove,
		Handle: func(event Event) {
			ce, ok := event.(MouseMove)

			if !ok {
				log.Warnf("mouse-move listener expected %v, but got %v", TypeMouseMove, event.Type())
				return
			}

			handler(ce)
		},
	}

	return &l
}

var _ Event = MouseMove{}

// MouseMove is triggered when the user click inside the rendering surface.
type MouseMove struct {
	Event
	Position geometry.Point
}

// MakeMouseMove returns a new MouseMove events.Event.
func MakeMouseMove(position geometry.Point) MouseMove {
	return MouseMove{
		Event:    MakeEvent(TypeMouseMove),
		Position: position,
	}
}
