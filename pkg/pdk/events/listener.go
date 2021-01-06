package events

import "fmt"

// Listener represents an object that can Handle an event dispatched by an Target object.
type Listener struct {
	Type   Type
	Handle func(event Event)
}

func (l *Listener) String() string {
	return fmt.Sprintf("(%v)-event-listener", l.Type)
}
