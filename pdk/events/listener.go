package events

// ListenerFunc is a function that implements the Listener interface.
type ListenerFunc func(event Event)

// HandleEvent implements the Listener interface.
func (lf ListenerFunc) HandleEvent(event Event) {
	lf(event)
}

// Listener define an event handler that listens for a events of a specific type.
type Listener interface {
	HandleEvent(event Event)
}
