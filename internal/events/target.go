package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

// Listener represents an object that can handle an event dispatched by an Target object.
type Listener func(event Event)

// Target define an object that can receive events and may have listeners for them.
type Target interface {
	AddEventListener(EventType, *Listener)
	RemoveEventListener(EventType, *Listener)
	DispatchEvent(event Event)
}

var _ Target = &target{}

// target is an implementation of the Target interface.
type target map[EventType][]*Listener

// MakeTarget return an event target
func MakeTarget() Target {
	return target(
		make(map[EventType][]*Listener),
	)
}

func (t target) AddEventListener(eventType EventType, listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")

	if t[eventType] == nil {
		t[eventType] = make([]*Listener, 0, 8)
	}

	t[eventType] = append(t[eventType], listener)
}

// RemoveEventListener removes an event listener of a specific event type from the target.
func (t target) RemoveEventListener(eventType EventType, listener *Listener) {
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
func (t target) DispatchEvent(event Event) {
	for _, listener := range t[event.Type()] {
		(*listener)(event)
	}
}
