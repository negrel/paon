package styles

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// SetPropertyListener convert the given event handler as a generic events.Listener.
func SetPropertyListener(handler func(setProperty EventSetProperty)) *events.Listener {
	l := events.Listener{
		Type: eventTypeSetProperty,
		Handle: func(event events.Event) {
			spe, ok := event.(EventSetProperty)

			if !ok {
				log.Warnf("click listener expected %v, but got %v", EventTypeSetProperty, event.Type())
				return
			}

			handler(spe)
		},
	}

	return &l
}

var eventTypeSetProperty = events.MakeType("set-property")

func EventTypeSetProperty() events.Type {
	return eventTypeSetProperty
}

type EventSetProperty struct {
	events.Event
	Old, New property.Property
}

func makeEventSetProperty(old, new property.Property) EventSetProperty {
	return EventSetProperty{
		Event: events.MakeEvent(eventTypeSetProperty),
		Old:   old,
		New:   new,
	}
}
