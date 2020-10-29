package events

import (
	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/utils"
)

// ClickListener convert the given event handler as a generic Listener.
func ClickListener(handler func(ClickEvent)) *Listener {
	l := Listener(func(event Event) {
		ce, ok := event.(ClickEvent)

		if !ok {
			log.Warnf("click listener expected %v, but got %v", ClickEventType, event.Type())
			return
		}

		handler(ce)
	})

	return &l
}

var _ Event = ClickEvent{}

// ClickEvent is triggered when the user click inside the rendering surface.
type ClickEvent struct {
	event
	Position utils.Point
}

// MakeResizeEvent returns a new ClickEvent object.
func MakeClickEvent(position utils.Point) ClickEvent {
	return ClickEvent{
		event:    makeEvent(ClickEventType),
		Position: position,
	}
}
