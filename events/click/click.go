package click

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
)

var EventType = events.NewType("Click")

// Listener returns an events.Listener that will call the given handler
// on click events.
func Listener(handler func(Event)) (events.Type, events.Handler) {
	return EventType, events.HandlerFunc(func(event events.Event) {
		assert.IsType(event, Event{})
		handler(event.(Event))
	})
}

// Event define a click event.
type Event struct {
	events.Event
	Position geometry.Vec2D
}

// New returns a new click event.
func New(pos geometry.Vec2D) Event {
	return Event{
		Event:    events.NewEvent(EventType),
		Position: pos,
	}
}

// String implements the fmt.Stringer interface.
func (r Event) String() string {
	return fmt.Sprintf("%s{Position: %+v}", r.Event.Type(), r.Position)
}
