package span

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/widgets"
)

type Widget struct {
	*widgets.BaseWidget

	text string
}

func New(text string) *Widget {
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
			Draw(w.text, surface)
		}),
	)

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

func Draw(text string, surface draw.Surface) {
	// TODO: iterate over grapheme instead of runes.
	for i, c := range text {
		surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style:   draw.CellStyle{},
			Content: c,
		})
	}
}
