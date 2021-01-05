package events

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
)

// ClickListener convert the given event handler as a generic Listener.
func ClickListener(handler func(Click)) *Listener {
	l := Listener(func(event Event) {
		ce, ok := event.(Click)

		if !ok {
			log.Warnf("click listener expected %v, but got %v", TypeClick, event.Type())
			return
		}

		handler(ce)
	})

	return &l
}

var _ Event = Click{}

// Click is triggered when the user click inside the rendering surface.
type Click struct {
	Event
	Position geometry.Point
}

// MakeClick returns a new Click events.Event.
func MakeClick(position geometry.Point) Click {
	return Click{
		Event:    MakeEvent(TypeClick),
		Position: position,
	}
}
