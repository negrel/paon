package events

import (
	"github.com/negrel/paon/geometry"
)

var MousePressEventType = NewType("MousePress")

// MousePressListener returns a listener that will call the given handler
// on mouse press
func MousePressListener(handler func(Event, MouseEventData)) (Type, Listener) {
	return MousePressEventType, NewListenerFunc(func(event Event) {
		handler(event, event.Data.(MouseEventData))
	})
}

// NewMousePress returns a new mouse press event.
func NewMousePress(pos geometry.Vec2D, buttons ButtonMask, mods ModMask) Event {
	data := MouseEventData{
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}

	return NewEvent(MousePressEventType, data)
}
