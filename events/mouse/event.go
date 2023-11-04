package mouse

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

// Event is a generic mouse event. A mouse event can be an event of type:
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
	// List of buttons that were pressed or wheel motions.
	Buttons ButtonMask
	// Modifiers is the modifiers that were present with the key press. Note
	// that not all platforms and terminals support this equally well, and some
	// cases we will not not know for sure. Hence, applications should avoid
	// using this in most circumstances.
	Modifiers keypress.ModMask
}
