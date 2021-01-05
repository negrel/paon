package events

var _ Event = Interrupt{}

// Interrupt is a generic wakeup event. It can be used to request
// a redraw. It can carry an arbitrary payload, as well.
type Interrupt struct {
	Event
	Data interface{}
}

// MakeInterrupt return a new Interrupt events.Event.
func MakeInterrupt(data interface{}) Interrupt {
	return Interrupt{
		Event: MakeEvent(TypeInterrupt),
		Data:  data,
	}
}
