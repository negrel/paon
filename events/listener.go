package events

import "github.com/negrel/paon/id"

// Handler define an event handler.
type Handler interface {
	id.Identifiable
	HandleEvent(event Event)
}

type handlerFunc struct {
	id      id.ID
	handler func(event Event)
}

func HandlerFunc(handler func(event Event)) Handler {
	return handlerFunc{id.New(), handler}
}

func (hf handlerFunc) HandleEvent(event Event) {
	hf.handler(event)
}

// ID implements the id.Identifiable interface.
func (hf handlerFunc) ID() id.ID {
	return hf.id
}

// IsSame implements the id.Identifiable interface.
func (hf handlerFunc) IsSame(other id.Identifiable) bool {
	return hf.id == other.ID()
}

// Listener define a pair event type and handler.
type Listener struct {
	EventType Type
	Handler   Handler
}
