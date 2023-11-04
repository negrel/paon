package mouse

import (
	"fmt"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var UpEventType = events.NewType("MouseUp")

// UpListener returns an events.Listener that will call the given handler
// on mouse up events.
func UpListener(handler func(UpEvent)) (events.Type, events.Handler) {
	return UpEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(UpEvent))
	})
}

// UpEvent define a mouse up event.
type UpEvent struct {
	events.Event
	// Mouse up absolute position.
	AbsPosition geometry.Vec2D
	// Position relative to element position.
	RelPosition geometry.Vec2D
	// List of buttons that were pressed or wheel motions.
	Buttons ButtonMask
	// Modifiers is the modifiers that were present with the key press. Note
	// that not all platforms and terminals support this equally well, and some
	// cases we will not not know for sure. Hence, applications should avoid
	// using this in most circumstances.
	Modifiers keypress.ModMask
}

// NewUp returns a new mouse up event.
func NewUp(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask) UpEvent {
	return UpEvent{
		Event:       events.NewEvent(UpEventType),
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}
}

// String implements the fmt.Stringer interface.
func (ue UpEvent) String() string {
	return fmt.Sprintf("%s{Position: %+v}", ue.Event.Type(), ue.AbsPosition)
}
