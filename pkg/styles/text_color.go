package styles

import (
	"github.com/gdamore/tcell"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
)

var _ render.DrawStep = TextColorProperty{}

// WidthProperty is a style property to define the size of a render.Object.
type TextColorProperty struct {
	property
	tcell.Color
}

// ColorProperty return a new ColorProperty object.
func TextColor(color tcell.Color) TextColorProperty {
	return TextColorProperty{
		property: property{
			step:     render.DrawStepType,
			name:     "text-color",
			priority: TextColorPriority,
		},
		Color: color,
	}
}

func (c TextColorProperty) Draw(ctx render.Context) {
	canvas := ctx.Canvas()
	for i := 0; i < canvas.Height(); i++ {
		for j := 0; j < canvas.Width(); j++ {
			cell := canvas.Get(geometry.Pt(j, i))
			cell.Style = cell.Style.Foreground(c.Color)
		}
	}
}
