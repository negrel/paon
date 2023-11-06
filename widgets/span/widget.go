package span

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

// Style define styling options for span rendering.
type Style draw.CellStyle

type Option func(*Widget)

// WithStyle return an option that sets span widget style.
func WithStyle(style Style) Option {
	return func(w *Widget) {
		w.style = style
	}
}

type Widget struct {
	*widgets.BaseWidget

	style Style
	text  string
}

func New(text string, options ...Option) *Widget {
	w := &Widget{
		text: text,
	}

	w.BaseWidget = widgets.NewBaseWidget(
		widgets.Wrap(w),
		widgets.LayoutFunc(
			func(co layout.Constraint) geometry.Size {
				return Layout(w.text, co)
			},
		),
		widgets.DrawerFunc(func(surface draw.Surface) {
			Draw(surface, w.text, w.style)
		}),
	)

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

func Layout(text string, co layout.Constraint) geometry.Size {
	return geometry.NewSize(len(text), 1)
}

func Draw(surface draw.Surface, text string, style Style) {
	// TODO: iterate over grapheme instead of runes.
	for i, c := range text {
		surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style:   draw.CellStyle(style),
			Content: c,
		})
	}
}
