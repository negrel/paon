package events

// listenerFunc is a function that implements the Listener interface.
type listenerFunc func(event Event)

// ListenerFunc returns a new Listener that wraps the given function.
func ListenerFunc(fn func(event Event)) Listener {
	a := listenerFunc(fn)
	return &a
}

// HandleEvent implements the Listener interface.
func (lf *listenerFunc) HandleEvent(event Event) {
	(*lf)(event)
}

// Listener define an event handler that listens for a events of a specific type.
type Listener interface {
	HandleEvent(event Event)
}
