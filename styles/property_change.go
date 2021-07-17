package styles

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
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

var eventTypePropertyChange = events.NewType("property-change")

// EventTypePropertyChange returns the events.Type of PropertyChangeEvent.
func EventTypePropertyChange() events.Type {
	return eventTypePropertyChange
}

// EventPropertyChange is triggered when the style property of a Style change.
type EventPropertyChange struct {
	events.Event
	Old, New property.Property
}

func newEventPropertyChange(old, new property.Property) EventPropertyChange {
	return EventPropertyChange{
		Event: events.NewEvent(eventTypePropertyChange),
		Old:   old,
		New:   new,
	}
}
