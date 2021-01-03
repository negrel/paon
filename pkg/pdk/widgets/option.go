package widgets

import (
	"github.com/negrel/paon/internal/events"
	pdkstyle "github.com/negrel/paon/pkg/pdk/style"
	"github.com/negrel/paon/pkg/pdk/style/property"
)

type Option func(widget *widget)

// Bind the given variable to the widget.
func Bind(variable *Widget) Option {
	return func(widget *widget) {
		*variable = widget
	}
}

// ID set the widget id (must be unique across your entire application).
func ID(id string) Option {
	return func(widget *widget) {
		widget.id = id
	}
}

// Listener add the given events.Listener for the given events.EventType to your widget.
func Listener(eventType events.EventType, listener events.Listener) Option {
	return func(widget *widget) {
		widget.AddEventListener(eventType, &listener)
	}
}

// Themes apply the given themes to the widget. Those theme can be shared across multiple
// widget and won't be modified by yourWidget.Theme().Set(property) method.
func Themes(themes ...pdkstyle.Theme) Option {
	return func(widget *widget) {
		widget.theme.themes = themes
	}
}

// Theme apply the given property to the widget theme. This theme is not shareable
// and can only be modified using the yourWidget.Theme().Set(property) method.
func Theme(props ...property.Property) Option {
	return func(widget *widget) {
		for _, prop := range props {
			widget.theme.Set(prop)
		}
	}
}
