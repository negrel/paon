package events

import (
	"github.com/negrel/debuggo/pkg/assert"
)

// EventType is the type of an Event
type EventType int

// List of existing event type.
const (
	ErrorEventType EventType = iota - 2
	UnsupportedEventType
	InterruptEventType
	ResizeEventType
	ClickEventType
	WheelEventType
	KeyboardEventType
)

var eventTypeName = map[EventType]string{
	ErrorEventType:       "ErrorEvent",
	UnsupportedEventType: "UnsupportedEvent",
	InterruptEventType:   "InterruptEvent",
	ResizeEventType:      "ResizeEvent",
	ClickEventType:       "ClickEvent",
	WheelEventType:       "WheelEvent",
	KeyboardEventType:    "KeyboardEvent",
}

// String implements the fmt.Stringer interface.
func (et EventType) String() string {
	name, ok := eventTypeName[et]
	assert.Truef(ok, "%v is an invalid event type", et)

	return name
}
