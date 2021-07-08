package events

import (
	"time"
)

// Event is a generic interface for all events
type Event interface {
	When() int64
	Type() Type
}

type event struct {
	eType     Type
	timeStamp int64
}

// NewEvent returns a new Event object of the given type. This function
// should be used as a base for real Event objects.
func NewEvent(eventType Type) Event {
	return newEvent(eventType)
}

func newEvent(eventType Type) event {
	return event{
		eType:     eventType,
		timeStamp: time.Now().UnixNano(),
	}
}

// Type returns the type of the event.
func (e event) Type() Type {
	return e.eType
}

// When returns the timestamp of the event.
func (e event) When() int64 {
	return e.timeStamp
}
