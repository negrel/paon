package themes

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/pkg/pdk/events"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
)

func ThemeChangeListener(handler func(event EventThemeChange)) *events.Listener {
	l := events.Listener{
		Type: eventTypeThemeChange,
		Handle: func(event events.Event) {
			tce, ok := event.(EventThemeChange)

			if !ok {
				log.Warnf("\"%v\" listener expected, but got %v", eventTypeThemeChange, event.Type())
				return
			}

			handler(tce)
		},
	}

	return &l
}

var eventTypeThemeChange = events.MakeType("theme-change")

func EventTypeThemeChange() events.Type {
	return eventTypeThemeChange
}

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
