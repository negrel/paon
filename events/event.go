package events

import (
	"time"
)

type Event struct {
	eType  Type
	time   time.Time
	target Target
	Data   any
}

// NewEvent returns a new Event object of the given type. This function
// should be used as a base for real Event objects.
func NewEvent(eventType Type, data any) Event {
	return newEvent(eventType, data)
}

func newEvent(eventType Type, data any) Event {
	return Event{
		eType:  eventType,
		time:   time.Now(),
		target: nil,
		Data:   data,
	}
}

// EventType returns the type of the event.
func (e Event) EventType() Type {
	return e.eType
}

// When returns the timestamp of the event.
func (e Event) When() time.Time {
	return e.time
}

// Target returns target set previously using WithTarget.
func (e Event) Target() Target {
	return e.target
}

// WithTarget creates a copy of this event with a different target and returns
// it.
func (e Event) WithTarget(t Target) Event {
	e.target = t
	return e
}
