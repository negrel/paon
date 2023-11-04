package mouse

import (
	"fmt"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var PressEventType = events.NewType("MousePress")

// Listener returns an events.Listener that will call the given handler
// on mouse press events.
func PressListener(handler func(PressEvent)) (events.Type, events.Handler) {
	return PressEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(PressEvent))
	})
}

// PressEvent define a mouse press event.
type PressEvent struct {
	events.Event
	// Mouse press absolute position.
	AbsPosition geometry.Vec2D
	// Position relative to parent position.
	RelPosition geometry.Vec2D
	// List of buttons that were pressed or wheel motions.
	Buttons ButtonMask
	// Modifiers is the modifiers that were present with the key press. Note
	// that not all platforms and terminals support this equally well, and some
	// cases we will not not know for sure. Hence, applications should avoid
	// using this in most circumstances.
	Modifiers keypress.ModMask
}

// NewPress returns a new mouse press event.
func NewPress(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask) PressEvent {
	return PressEvent{
		Event:       events.NewEvent(PressEventType),
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}
}

// String implements the fmt.Stringer interface.
func (pe PressEvent) String() string {
	return fmt.Sprintf("%s{Position: %+v}", pe.Event.Type(), pe.AbsPosition)
}
