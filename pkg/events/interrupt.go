package events

import "github.com/negrel/paon/pkg/pdk/events"

var _TypeInterrupt = events.MakeType("interrupt")

func TypeInterrupt() events.Type {
	return _TypeInterrupt
}

var _ events.Event = Interrupt{}

// Interrupt is a generic wakeup event. It can be used to request
// a redraw. It can carry an arbitrary payload, as well.
type Interrupt struct {
	events.Event
	Data interface{}
}

// MakeInterrupt return a new Interrupt events.Event.
func MakeInterrupt(data interface{}) Interrupt {
	return Interrupt{
		Event: events.MakeEvent(_TypeInterrupt),
		Data:  data,
	}
}
