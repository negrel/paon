package themes

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/events"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
)

func ThemeChangeListener(handler func(event EventThemeChange)) *events.Listener {
	l := events.Listener{
		Type: eventTypeThemeChange,
		Handle: func(event events.Event) {
			assert.IsType(event, makeEventThemeChange(nil, true))
			handler(event.(EventThemeChange))
		},
	}

	return &l
}

var eventTypeThemeChange = events.MakeType("theme-change")

func EventTypeThemeChange() events.Type {
	return eventTypeThemeChange
}

var _ events.Event = EventThemeChange{}

type EventThemeChange struct {
	events.Event
	pdkstyles.Style
	DeletedStyle bool
}

func makeEventThemeChange(style pdkstyles.Style, deleted bool) EventThemeChange {
	return EventThemeChange{
		Event:        events.MakeEvent(eventTypeThemeChange),
		Style:        style,
		DeletedStyle: deleted,
	}
}
