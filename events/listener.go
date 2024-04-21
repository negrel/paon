package events

// Listener define an objectc listening for events.
type Listener interface {
	HandleEvent(event Event)
}

// ListenerFunc is wrapper around handler function that implements Listener.
type ListenerFunc struct {
	listenerPtr *func(event Event)
}

// NewListenerFunc returns a new
func NewListenerFunc(fn func(event Event)) ListenerFunc {
	return ListenerFunc{listenerPtr: &fn}
}

// HandleEvent implements Listener.
func (lf ListenerFunc) HandleEvent(event Event) {
	(*lf.listenerPtr)(event)
}

