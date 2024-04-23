package events

import (
	"github.com/negrel/paon/geometry"
)

// Copied from tcell

// ButtonMask is a mask of mouse buttons and wheel events. Mouse button presses
// are normally delivered as both press and release events. Mouse wheel events
// are normally just single impulse events. Windows supports up to eight
// separate buttons, but XTerm can only support mouse buttons 1-3.
// Its not unheard of for terminals to support only one or two buttons
// (think Macs). Old terminals, and true emulations (such as vt100) won't support
// mice at all, of course.
type ButtonMask int16

// These are the actual button values. Note that tcell version 1.x reversed buttons
// two and three on *nix based terminals. We use button 1 as the primary, and
// button 2 as the secondary, and button 3 (which is often missing) as the middle.
const (
	Button1 ButtonMask = 1 << iota // Usually the left (primary) mouse button.
	Button2                        // Usually the right (secondary) mouse button.
	Button3                        // Usually the middle mouse button.
	Button4                        // Often a side button (thumb/next).
	Button5                        // Often a side button (thumb/prev).
	Button6
	Button7
	Button8
	ButtonNone ButtonMask = 0 // No button.

	ButtonPrimary   = Button1
	ButtonSecondary = Button2
	ButtonMiddle    = Button3
)

// Not copied from tcell

var MouseEventType = NewType("Mouse")

// MouseEventListener returns a listener that will call the given handler
// on mouse
func MouseEventListener(handler func(Event, MouseEventData)) (Type, Listener) {
	return MouseEventType, NewListenerFunc(func(event Event) {
		handler(event, event.Data.(MouseEventData))
	})
}

// MouseEventData define data contained in all mouse event.
type MouseEventData struct {
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
	Modifiers ModMask
}

// NewMouseEvent returns a new mouse event.
func NewMouseEvent(pos geometry.Vec2D, buttons ButtonMask, mods ModMask) Event {
	data := MouseEventData{
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}

	return NewEvent(MouseEventType, data)
}

// RelativePosition implements PointerEventData.
func (med MouseEventData) RelativePosition() geometry.Vec2D {
	return med.RelPosition
}

// WithPositionRelativeToOrigin implements PointerEventData.
func (med MouseEventData) WithPositionRelativeToOrigin(origin geometry.Vec2D) PointerEventData {
	med.RelPosition = med.RelPosition.Sub(origin)
	return med
}
