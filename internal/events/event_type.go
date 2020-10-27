package events

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
