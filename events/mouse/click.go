package mouse

import (
	"fmt"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var ClickEventType = events.NewType("Click")

// ClickListener returns an events.Listener that will call the given handler
// on click events.
func ClickListener(handler func(ClickEvent)) (events.Type, events.Handler) {
	return ClickEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(ClickEvent))
	})
}

// ClickEvent define a click event.
type ClickEvent struct {
	events.Event
	// Mouse absolute position.
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

// NewClick returns a new click event.
func NewClick(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask) ClickEvent {
	return ClickEvent{
		Event:       events.NewEvent(ClickEventType),
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}
}

// String implements the fmt.Stringer interface.
func (ce ClickEvent) String() string {
	return fmt.Sprintf("%s{Position: %+v}", ce.Event.Type(), ce.AbsPosition)
}
