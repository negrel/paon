package mouse

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var UpEventType = events.NewType("MouseUp")

// UpListener returns an events.Listener that will call the given handler
// on mouse up events.
func UpListener(handler func(Event)) (events.Type, events.Handler) {
	return UpEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(Event))
	})
}

// NewUp returns a new mouse up event.
func NewUp(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask) Event {
	return Event{
		Event:       events.NewEvent(UpEventType),
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}
}
