package widgets

import (
	"github.com/negrel/paon/internal/events"
	pdkstyle "github.com/negrel/paon/pkg/pdk/style"
	"github.com/negrel/paon/pkg/pdk/style/property"
)

type Option func(widget *widget)

func ID(id string) Option {
	return func(widget *widget) {
		widget.id = id
	}
}

func Listener(eventType events.EventType, listener events.Listener) Option {
	return func(widget *widget) {
		widget.AddEventListener(eventType, &listener)
	}
}

func Themes(themes ...pdkstyle.Theme) Option {
	return func(widget *widget) {
		widget.theme.themes = themes
	}
}

func Props(props ...property.Property) Option {
	return func(widget *widget) {
		for _, prop := range props {
			widget.theme.Set(prop)
		}
	}
}
