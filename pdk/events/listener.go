package events

import "fmt"

// Listener define an event handler that listens for a events of a specific type.
type Listener struct {
	Type   Type
	Handle func(event Event)
}

func (l *Listener) String() string {
	return fmt.Sprintf("(%v)-event-listener", l.Type)
}

// TreeListener define a listener for TreeEvents.
type TreeListener struct {
	Listener
	Phase Phase
}
