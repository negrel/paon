package events

import (
	"time"
)

// Event is a generic interface for all events
type Event interface {
	When() int64
	Type() EventType
}

type event struct {
	eType     EventType
	timeStamp int64
}

func makeEvent(eventType EventType) event {
	return event{
		eType:     eventType,
		timeStamp: time.Now().UnixNano(),
	}
}

// Type returns the type of the event.
func (e event) Type() EventType {
	return e.eType
}

// When returns the timestamp of the event.
func (e event) When() int64 {
	return e.timeStamp
}
