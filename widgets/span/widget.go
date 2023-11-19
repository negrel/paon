package span

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/widgets"
)

type Option func(*Widget)

// WithStyle define span internal styles.Style.
func WithStyle(style widgets.Style) Option {
	return func(w *Widget) {
		w.style.InnerStyle = style
	}
}

type Widget struct {
	*widgets.BaseWidget

	style widgets.InheritStyle
	text  string
}

func New(text string, options ...Option) *Widget {
	w := &Widget{
		text: text,
	}

	w.BaseWidget = widgets.NewBaseWidget(w)
	w.style = widgets.InheritStyle{
		Widget:     w,
		InnerStyle: widgets.Style{},
	}

	for _, applyOption := range options {
		applyOption(w)
	}

	return w
}

func (w *Widget) SetText(txt string) {
	w.text = txt
}

func (w *Widget) Text() string {
	return w.text
}

// Layout implements layout.Layout.
func (w *Widget) Layout(co layout.Constraint) geometry.Size {
	return styles.Layout(w.style.Compute(), co, layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
		return Layout(w.text, co)
	}))
}

// Draw implements draw.Drawer.
func (w *Widget) Draw(surface draw.Surface) {
	style := w.style.Compute()
	surface = styles.Draw(style, surface)

	Draw(surface, w.text, style.CellStyle)
}

// Style implements styles.Styled.
func (w *Widget) Style() styles.Style {
	return w.style
}

func Layout(text string, co layout.Constraint) geometry.Size {
	return geometry.NewSize(len(text), 1)
}

func Draw(surface draw.Surface, text string, style draw.CellStyle) {
	// TODO: iterate over grapheme instead of runes.
	for i, c := range text {
		surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style:   style,
			Content: c,
		})
	}
}
