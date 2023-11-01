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
	)

	return w
}

func (s *Span) Render(co layout.Constraint, surface draw.Surface) geometry.Size {
	// TODO: iterate over grapheme instead of runes.
	for i, c := range s.text {
		surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
			Style:   draw.CellStyle{},
			Content: c,
		})
	}

	return geometry.NewSize(len(s.text), 1)
}
