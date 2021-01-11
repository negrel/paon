package widgets

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
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

// Styles apply the given styles to the widget. Those theme can be shared across multiple
// widget and won't be modified by any theme modifier method. You can still remove styles
// m widgets using the DelStyle method.
func Styles(styles ...pdkstyle.Style) Option {
	return func(widget *widget) {
		for _, style := range styles {
			widget.theme.AddStyle(style)
		}
	}
}

// Props apply the given properties to the widget theme. This theme is not shareable
// and can only be modified using Set/Del method on the theme of the widget.
func Props(props ...property.Property) Option {
	return func(widget *widget) {
		for _, prop := range props {
			widget.theme.Set(prop)
		}
	}
}
