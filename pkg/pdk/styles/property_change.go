package styles

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// PropertyChangeListener convert the given event handler as a generic events.Listener.
func PropertyChangeListener(handler func(setProperty EventPropertyChange)) *events.Listener {
	l := events.Listener{
		Type: eventTypePropertyChange,
		Handle: func(event events.Event) {
			assert.IsType(event, EventPropertyChange{})
			handler(event.(EventPropertyChange))
		},
	}

	return &l
}

var eventTypePropertyChange = events.MakeType("set-property")

func EventTypePropertyChange() events.Type {
	return eventTypePropertyChange
}

type EventPropertyChange struct {
	events.Event
	Old, New property.Property
}

func makeEventPropertyChange(old, new property.Property) EventPropertyChange {
	return EventPropertyChange{
		Event: events.MakeEvent(eventTypePropertyChange),
		Old:   old,
		New:   new,
	}
}
