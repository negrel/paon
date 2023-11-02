package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

type Span struct {
	*pdkwidgets.BaseWidget

	text string
}

func NewSpan(text string) *Span {
	w := &Span{
		text: text,
	}

	w.BaseWidget = pdkwidgets.NewBaseWidget(
		pdkwidgets.Wrap(w),
		pdkwidgets.LayoutFunc(
			func(co layout.Constraint) geometry.Size {
				return LayoutSpan(w.text, co)
			},
		),
		pdkwidgets.DrawerFunc(func(surface draw.Surface) {
			DrawSpan(w.text, surface)
		}),
	)

	return w
}

func (s *Span) SetText(txt string) {
	s.text = txt
}

func (s *Span) Text() string {
	return s.text
}

func LayoutSpan(text string, co layout.Constraint) geometry.Size {
	return geometry.NewSize(len(text), 1)
}

func DrawSpan(text string, surface draw.Surface) {
	// TODO: iterate over grapheme instead of runes.
	for i, c := range text {
		surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style:   draw.CellStyle{},
			Content: c,
		})
	}
}
