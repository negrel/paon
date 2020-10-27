package events

import (
	"fmt"
)

// UnsupportedEvent define any events that is not in the list supported events.
type UnsupportedEvent struct {
	event
	str string
}

// MakeUnsupportedEvent return a new UnsupportedEvent object.
func MakeUnsupportedEvent(content string) UnsupportedEvent {
	return UnsupportedEvent{
		event: makeEvent(UnsupportedEventType),
		str:   content,
	}
}

// String implements the fmt.Stringer interface.
func (ue UnsupportedEvent) String() string {
	return fmt.Sprintf("%v: %v", ue.eType, ue.str)
}
