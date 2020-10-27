package events

var _ Event = InterruptEvent{}

// InterruptEvent is a generic wakeup event. Its can be used to request
// a redraw. It can carry an arbitrary payload, as well.
type InterruptEvent struct {
	event
	Data interface{}
}

// MakeInterruptEvent return a new InterruptEvent object.
func MakeInterruptEvent(data interface{}) InterruptEvent {
	return InterruptEvent{
		event: makeEvent(InterruptEventType),
		Data:  data,
	}
}
