package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

// Target define an object that can receive events and may have listeners for them.
type Target interface {
	AddEventListener(*Listener)
	RemoveEventListener(*Listener)
	DispatchEvent(event Event)
}

var _ Target = &target{}

// target is an implementation of the Target interface.
type target map[Type][]*Listener

// MakeTarget return an event target
func MakeTarget() Target {
	return target(
		make(map[Type][]*Listener),
	)
}

func (t target) AddEventListener(listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")

	if t[listener.Type] == nil {
		t[listener.Type] = make([]*Listener, 0, 8)
	}

	t[listener.Type] = append(t[listener.Type], listener)
}

// RemoveEventListener removes an event listener of a specific event type from the target.
func (t target) RemoveEventListener(listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")

	if t[listener.Type] == nil {
		log.Infof("can't remove event listener %v, no such event listener registered for %v event type", listener, listener.Type)
		return
	}

	for i, l := range t[listener.Type] {
		if l == listener {
			t[listener.Type] = append(t[listener.Type][:i], t[listener.Type][i+1:]...)
			return
		}
	}

	log.Infof("can't remove event listener %v, not found", listener)
}

// DispatchEvent dispatch the given event to listeners.
func (t target) DispatchEvent(event Event) {
	for _, listener := range t[event.Type()] {
		listener.Handle(event)
	}
}
