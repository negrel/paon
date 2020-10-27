package events

var _ Event = WheelEvent{}

// WheelEvent define a user interaction with the
type WheelEvent struct {
	event
}

// MakeWheelEvent returns a new WheelEvent object.
func MakeWheelEvent() WheelEvent {
	return WheelEvent{
		event: makeEvent(WheelEventType),
	}
}
