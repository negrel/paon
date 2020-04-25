package events

import (
	"time"
)

// EventType define the type of an event.
type EventType int

// Event define a generic interface for events.
type Event interface {
	When() time.Time
	Type() EventType
}

// List of existing event type.
const (
	ErrorEventType EventType = iota - 2
	InterruptEventType
	ResizeEventType
	ClickEventType
	ScrollEventType
	KeyboardEventType
)

type event struct {
	evType    EventType
	timeStamp time.Time
}

func (e *event) Type() EventType {
	return e.evType
}

func (e *event) When() time.Time {
	return e.timeStamp
}
