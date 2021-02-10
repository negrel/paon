package painter

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
)

var _ draw.Painter = block{}

type block struct {
}

func makeBlock() draw.Painter {
	return block{}
}

// Layout implements the draw.Painter interface.
func (b block) Layout(ctx draw.Context) {
	box := ApplyMargin(ctx.Object.Style(), ctx.Constraint.Max)
	box = ApplyBorder(ctx.Object.Style(), box)

	width := box.Width()

	height := ComputeObjectHeight(ctx.Object)
	height = math.Max(height, ctx.Constraint.Min.Height())
	height = math.Min(height, ctx.Constraint.Max.Height())

	ctx.Canvas.Resize(geometry.MakeSize(width, height))
}

// Draw implements the draw.Painter interface.
func (b block) Draw(ctx draw.Context) {

}
