package events

import (
	"github.com/negrel/paon/geometry"
)

// ResizeEventData define data contained in events of type resize.
type ResizeEventData struct {
	Old, New geometry.Size
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
func NewResizeEvent(old, new geometry.Size) Event {
	data := ResizeEventData{old, new}
	return NewEvent(ResizeEventType, data)
}

// IsWider returns true if the new terminal size is wider.
func (red ResizeEventData) IsWider() bool {
	return red.Old.Size().Width < red.New.Size().Width
}

// IsGreater returns true if the new terminal size is greater.
func (red ResizeEventData) IsGreater() bool {
	return red.Old.Size().Height < red.New.Size().Height
}
