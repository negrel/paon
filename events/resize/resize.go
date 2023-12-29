package resize

import (
	"fmt"

	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
)

var EventType = events.NewType("resize")

// Listener returns an events.Listener that will call the given handler
// on resize events.
func Listener(handler func(Event)) (events.Type, events.Handler) {
	return EventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(Event))
	})
}

// Event define a resize event.
type Event struct {
	events.Event
	Old, New geometry.Size
}

// New returns a new resize event.
func New(old, new geometry.Size) Event {
	return Event{
		Event: events.NewEvent(EventType),
		Old:   old,
		New:   new,
	}
}

// String implements the fmt.Stringer interface.
func (r Event) String() string {
	return fmt.Sprintf("%s{Old: %v, New: %v}", r.Event.Type(), r.Old, r.New)
}

// IsWider returns true if the new terminal size is wider.
func (r Event) IsWider() bool {
	return r.Old.Size().Width < r.New.Size().Width
}

// IsGreater returns true if the new terminal size is greater.
func (r Event) IsGreater() bool {
	return r.Old.Size().Height < r.New.Size().Height
}
