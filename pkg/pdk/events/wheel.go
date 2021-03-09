package events

var _ Event = Wheel{}

// Wheel define a user interaction with the
type Wheel struct {
	Event
}

// MakeWheel returns a new Wheel object.
func MakeWheel() Wheel {
	return Wheel{
		Event: MakeEvent(TypeWheel()),
	}
}
