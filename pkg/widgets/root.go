package widgets

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"

	"github.com/negrel/paon/internal/tree"
)

type rLayout = Layout

// Root define the root node of the Widget tree.
type Root struct {
	rLayout
	children Widget
}

func NewRoot(child Widget) *Root {
	r := &Root{
		rLayout: newLayout(tree.NewRoot(child)),
	}
	InnerDrawFunc(func(ctx render.Context) {})(r)

	return r
}

func (r *Root) Render(screen geometry.Size) draw.Canvas {
	canvas := draw.MakeCanvas(geometry.Rect(0, 0, screen.Width(), screen.Height()))
	ctx := render.MakeContext(r, &canvas)

	r.Layout(ctx)
	r.Draw(ctx)

	return canvas
}
