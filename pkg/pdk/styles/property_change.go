package styles

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// PropertyChangeListener convert the given event handler as a generic events.Listener.
func PropertyChangeListener(handler func(setProperty EventPropertyChange)) *events.Listener {
	l := events.Listener{
		Type: eventTypeSetProperty,
		Handle: func(event events.Event) {
			spe, ok := event.(EventPropertyChange)

			if !ok {
				log.Warnf("click listener expected %v, but got %v", EventTypePropertyChange, event.Type())
				return
			}

			handler(spe)
		},
	}

	return &l
}

var eventTypeSetProperty = events.MakeType("set-property")

func EventTypePropertyChange() events.Type {
	return eventTypeSetProperty
}

type EventPropertyChange struct {
	events.Event
	Old, New property.Property
}

func makeEventSetProperty(old, new property.Property) EventPropertyChange {
	return EventPropertyChange{
		Event: events.MakeEvent(eventTypeSetProperty),
		Old:   old,
		New:   new,
	}
}
