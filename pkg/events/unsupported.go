package events

import "github.com/negrel/paon/pkg/pdk/events"

var _TypeUnsupported = events.MakeType("unsupported")

func TypeUnsupported() events.Type {
	return _TypeUnsupported
}

// Unsupported define any events that is not in the list of the supported events.
type Unsupported struct {
	events.Event
	str string
}

// MakeUnsupported return a new Unsupported object.
func MakeUnsupported(content string) Unsupported {
	return Unsupported{
		Event: events.MakeEvent(_TypeUnsupported),
		str:   content,
	}
}
