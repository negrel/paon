package events

import "github.com/negrel/paon/pkg/pdk/events"

var _TypeWheel = events.MakeType("wheel")

func TypeWheel() events.Type {
	return _TypeWheel
}

var _ events.Event = Wheel{}

// Wheel define a user interaction with the
type Wheel struct {
	events.Event
}

// MakeWheel returns a new Wheel object.
func MakeWheel() Wheel {
	return Wheel{
		Event: events.MakeEvent(_TypeWheel),
	}
}
