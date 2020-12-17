package styles

import (
	"github.com/gdamore/tcell"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
)

var _ render.DrawStep = BackgroundColorProperty{}

// WidthProperty is a style property to define the size of a render.Object.
type BackgroundColorProperty struct {
	property
	tcell.Color
}

// BGColor return a new BackgroundColorProperty object.
func BGColor(color tcell.Color) BackgroundColorProperty {
	return BackgroundColorProperty{
		property: property{
			step:     render.DrawStepType,
			name:     "background-color",
			priority: BackgroundColorPriority,
		},
		Color: color,
	}
}

func (bg BackgroundColorProperty) Draw(ctx render.Context) {
	canvas := ctx.Canvas()
	for i := 0; i < canvas.Height(); i++ {
		for j := 0; j < canvas.Width(); j++ {
			cell := canvas.Get(geometry.Pt(j, i))
			cell.Style = cell.Style.Background(bg.Color)
		}
	}
}
