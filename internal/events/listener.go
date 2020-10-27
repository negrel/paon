package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

// Listener represents an object that can handle an event dispatched by a Target object.
type Listener func(event Event)

// NewListener return a new listener object
func NewListener(eventHandler interface{}) *Listener {
	l := makeListener(eventHandler)
	return &l
}

func makeListener(eventHandler interface{}) Listener {
	assert.NotNil(eventHandler, "event handler must be non-nil")

	switch handler := eventHandler.(type) {
	case func(ClickEvent):
		return func(event Event) {
			e, ok := event.(ClickEvent)
			if ok {
				handler(e)
			}
		}

	case func(KeyboardEvent):
		return func(event Event) {
			e, ok := event.(KeyboardEvent)
			if ok {
				handler(e)
			}
		}

	default:
		log.Panicf("%v is not a valid event handler\n", eventHandler)
		return nil
	}
}
