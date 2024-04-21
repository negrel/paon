package mouse

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var MouseEventType = events.NewType("Mouse")

// EventListener returns an events.Listener that will call the given handler
// on mouse events.
func EventListener(handler func(Event)) (events.Type, events.Handler) {
	return MouseEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(Event))
	})
}

// Event is a generic mouse event. A mouse event can be of events.Type:
//
// * MousePress
//
// * MouseUp
type Event struct {
	events.Event
	// Mouse absolute position.
	AbsPosition geometry.Vec2D
	// Position relative to element position.
	RelPosition geometry.Vec2D
	// List of buttons that were pressed.
	Buttons ButtonMask
	// Modifiers is the modifiers that were present with the key press. Note
	// that not all platforms and terminals support this equally well, and some
	// cases we will not not know for sure. Hence, applications should avoid
	// using this in most circumstances.
	Modifiers keypress.ModMask
}

// NewEvent returns a new mouse event.
func NewEvent(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask) Event {
	return Event{
		Event:       events.NewEvent(MouseEventType),
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}
}
