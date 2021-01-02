package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

// Listener represents an object that can handle an event dispatched by an EventTarget object.
type Listener func(event Event)

// EventTarget define an object that can receive events and may have listeners for them.
type EventTarget interface {
	AddEventListener(EventType, *Listener)
	RemoveEventListener(EventType, *Listener)
	DispatchEvent(event Event)
}

var _ EventTarget = &Target{}

// Target is an implementation of the EventTarget interface.
type Target map[EventType][]*Listener

// MakeTarget return an event target
func MakeTarget() Target {
	return make(map[EventType][]*Listener)
}

func (t Target) AddEventListener(eventType EventType, listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")

	if t[eventType] == nil {
		t[eventType] = make([]*Listener, 0, 8)
	}

	t[eventType] = append(t[eventType], listener)
}

// RemoveEventListener removes an event listener of a specific event type from the Target.
func (t Target) RemoveEventListener(eventType EventType, listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")

	if t[eventType] == nil {
		log.Infof("can't remove event listener %v, no event listener for %v event type", listener, eventType.String())
		return
	}

	for i, l := range t[eventType] {
		if l == listener {
			t[eventType] = append(t[eventType][:i], t[eventType][i+1:]...)
			return
		}
	}

	log.Infof("can't remove event listener %v, not found")
}

// DispatchEvent dispatch the given event to listeners.
func (t Target) DispatchEvent(event Event) {
	for _, listener := range t[event.Type()] {
		(*listener)(event)
	}
}
