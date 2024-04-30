package events

import (
	"github.com/negrel/paon/geometry"
)

// ResizeEventData define data contained in events of type resize.
type ResizeEventData struct {
	NewSize geometry.Size
}

var ResizeEventType = NewType("Resize")

// ResizeListener returns a listener that will call the given handler
// on resize events.
func ResizeListener(handler func(Event, ResizeEventData)) (Type, Listener) {
	return ResizeEventType, NewListenerFunc(func(event Event) {
		handler(event, event.Data.(ResizeEventData))
	})
}

// NewResizeEvent returns a new resize event.
func NewResizeEvent(newSize geometry.Size) Event {
	data := ResizeEventData{newSize}
	return NewEvent(ResizeEventType, data)
}
