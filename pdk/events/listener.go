package events

// HandlerFunc is a function that implements the Listener interface.
type HandlerFunc func(event Event)

// HandleEvent implements the Listener interface.
func (lf HandlerFunc) HandleEvent(event Event) {
	lf(event)
}

// Handler define an event handler.
type Handler interface {
	HandleEvent(event Event)
}

// Listener define a pair event type and handler.
type Listener struct {
	EventType Type
	Handler   Handler
}
