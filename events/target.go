package events

import (
	"sync"

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
type target struct {
	sync.RWMutex
	listeners map[Type][]*Listener
}

// NewTarget return a new event Target with no listeners.
func NewTarget() Target {
	return &target{
		listeners: make(map[Type][]*Listener),
	}
}

func (t *target) AddEventListener(listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")
	t.Lock()
	defer t.Unlock()

	if t.listeners[listener.Type] == nil {
		t.listeners[listener.Type] = make([]*Listener, 0, 8)
	}

	t.listeners[listener.Type] = append(t.listeners[listener.Type], listener)
}

// RemoveEventListener removes an event listener of a specific event type from the target.
func (t *target) RemoveEventListener(listener *Listener) {
	assert.NotNil(listener, "listener must be not nil")
	t.Lock()
	defer t.Unlock()

	if t.listeners[listener.Type] == nil {
		log.Infof("can't remove event listener %v, no such event listener registered for %v event type", listener, listener.Type)
		return
	}

	for i, l := range t.listeners[listener.Type] {
		if l == listener {
			t.listeners[listener.Type] = append(t.listeners[listener.Type][:i], t.listeners[listener.Type][i+1:]...)
			return
		}
	}

	log.Infof("can't remove event listener %v, not found", listener)
}

// DispatchEvent dispatch the given event to listeners.
func (t *target) DispatchEvent(event Event) {
	t.RLock()
	defer t.RUnlock()

	for _, listener := range t.listeners[event.Type()] {
		listener.Handle(event)
	}
}
