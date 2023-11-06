package button

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/span"
)

// Style define styling options for button rendering.
type Style span.Style

type Option func(*Widget)

// WithStyle return an option that sets button widget style.
func WithStyle(style Style) Option {
	return func(w *Widget) {
		w.style = style
	}
}

func OnClick(handler func(event mouse.ClickEvent)) Option {
	return func(w *Widget) {
		w.AddEventListener(mouse.ClickListener(handler))
	}
}

type Widget struct {
	*widgets.BaseWidget

	style Style
	text  string
}

func New(text string, options ...Option) *Widget {
	w := &Widget{
		style: Style{
			Foreground:    0,
			Background:    0,
			Bold:          false,
			Blink:         false,
			Reverse:       true,
			Underline:     false,
			Dim:           false,
			Italic:        false,
			StrikeThrough: false,
		},
		text: text,
	}

	w.BaseWidget = widgets.NewBaseWidget(
		widgets.Wrap(w),
		widgets.LayoutFunc(
			func(co layout.Constraint) geometry.Size {
				return span.Layout(w.text, co)
			},
		),
		widgets.DrawerFunc(func(surface draw.Surface) {
			span.Draw(surface, w.text, span.Style(w.style))
		}),
	)

	w.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		w.style.Reverse = !w.style.Reverse
	}))

	w.AddEventListener(mouse.ClickListener(func(event mouse.ClickEvent) {
		w.style.Reverse = !w.style.Reverse
	}))

	for _, applyOption := range options {
		applyOption(w)
	}

	return w
}
