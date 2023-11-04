package mouse

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var PressEventType = events.NewType("MousePress")

// Listener returns an events.Listener that will call the given handler
// on mouse press events.
func PressListener(handler func(Event)) (events.Type, events.Handler) {
	return PressEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(Event))
	})
}

// NewPress returns a new mouse press event.
func NewPress(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask) Event {
	return Event{
		Event:       events.NewEvent(PressEventType),
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}
}
