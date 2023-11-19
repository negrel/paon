package button

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/widgets"
	"github.com/negrel/paon/widgets/span"
)

type Option func(*Widget)

// WithStyle return an option that sets button widget style.
func WithStyle(style widgets.Style) Option {
	return func(w *Widget) {
		w.style.InnerStyle = style
	}
}

func OnClick(handler func(event mouse.ClickEvent)) Option {
	return func(w *Widget) {
		w.AddEventListener(mouse.ClickListener(handler))
	}
}

// Widget define a clickable widget button.
type Widget struct {
	*widgets.BaseWidget

	style widgets.InheritStyle
	text  string
}

// New returns a new button widget configured with the given options.
func New(text string, options ...Option) *Widget {
	w := &Widget{
		text: text,
	}

	w.BaseWidget = widgets.NewBaseWidget(w)
	w.style = widgets.InheritStyle{
		Widget:     w,
		InnerStyle: widgets.Style{},
	}

	w.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		w.style.InnerStyle = w.style.InnerStyle.Reverse(true)
	}))

	w.AddEventListener(mouse.ClickListener(func(event mouse.ClickEvent) {
		w.style.InnerStyle = w.style.InnerStyle.Reverse(false)
	}))

	for _, applyOption := range options {
		applyOption(w)
	}

	return w
}

// Layout implements layout.Layout.
func (w *Widget) Layout(co layout.Constraint) geometry.Size {
	return styles.Layout(w.style.Compute(), co, layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
		return span.Layout(w.text, co)
	}))
}

// Draw implements draw.Drawer.
func (w *Widget) Draw(surface draw.Surface) {
	style := w.style.Compute()

	surface = styles.Draw(style, surface)

	span.Draw(surface, w.text, style.CellStyle)
}

// Style implements styles.Styled.
func (w *Widget) Style() styles.Style {
	return w.style
}
