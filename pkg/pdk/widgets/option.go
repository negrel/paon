package widgets

import (
	"github.com/negrel/paon/pkg/pdk/draw"
	"github.com/negrel/paon/pkg/pdk/flows"
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

type Option func(widget *widget)

func MergeOptions(opts []Option, extension ...Option) []Option {
	for _, opt := range extension {
		opts = append(opts, opt)
	}

	return opts
}

// Bind binds the given variable to the widget.
func Bind(variable *Widget) Option {
	return func(widget *widget) {
		*variable = widget
	}
}

func Algo(algorithm func(flows.Constraint) flows.BoxModel) Option {
	return Flowable(flows.Algorithm(algorithm))
}

func Flowable(flowable flows.Flowable) Option {
	return func(widget *widget) {
		widget.Flowable = flowable
	}
}

func Script(script draw.Script) Option {
	return Drawable(script)
}

func Drawable(drawable draw.Drawable) Option {
	return func(widget *widget) {
		widget.Drawable = drawable
	}
}

// Styles applies the given styles to the widget. Those theme can be shared across multiple
// widget and won't be modified by any theme modifier method. You can still remove styles
// m widgets using the DelStyle method.
func Styles(styles ...pdkstyle.Style) Option {
	return func(widget *widget) {
		for _, style := range styles {
			widget.theme.AddStyle(style)
		}
	}
}

// Props applies the given properties to the widget theme. This theme is not shareable
// and can only be modified using Set/Del method on the theme of the widget.
func Props(props ...property.Property) Option {
	return func(widget *widget) {
		for _, prop := range props {
			widget.theme.Set(prop)
		}
	}
}
