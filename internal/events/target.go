package events

import (
	"github.com/negrel/debuggo/pkg/assert"
)

type _Target interface {
	AddEventListener(EventType, *Listener)
	RemoveEventListener(EventType, *Listener)
	DispatchEvent(event Event)
}

// Target define an object that can receive events and may have listeners for them.
type Target map[EventType]map[*Listener]struct{}

// MakeTarget return an event target
func MakeTarget() Target {
	return make(map[EventType]map[*Listener]struct{})
}

// AddEventListener registers an event handler of a specific event type on the Target.
func (t Target) AddEventListener(eventType EventType, listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")

	if t[eventType] == nil {
		t[eventType] = make(map[*Listener]struct{})
	}

	t[eventType][listener] = struct{}{}
}

// RemoveEventListener removes an event listener of a specific event type from the Target.
func (t Target) RemoveEventListener(eventType EventType, listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")
	delete(t[eventType], listener)
}

// DispatchEvent dispatch the given event to listeners.
func (t Target) DispatchEvent(event Event) {
	for listener := range t[event.Type()] {
		(*listener)(event)
	}
}
