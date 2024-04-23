package events

import (
	"github.com/negrel/paon/geometry"
)

var MouseUpEventType = NewType("MouseUp")

// MouseUpListener returns a listener that will call the given handler
// on mouse up
func MouseUpListener(handler func(Event, MouseEventData)) (Type, Listener) {
	return MouseUpEventType, NewListenerFunc(func(event Event) {
		handler(event, event.Data.(MouseEventData))
	})
}

// NewMouseUp returns a new mouse up event.
func NewMouseUp(pos geometry.Vec2D, buttons ButtonMask, mods ModMask) Event {
	data := MouseEventData{
		AbsPosition: pos,
		RelPosition: pos,
		Buttons:     buttons,
		Modifiers:   mods,
	}

	return NewEvent(MouseUpEventType, data)
}
