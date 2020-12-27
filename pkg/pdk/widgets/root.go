package widgets

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"

	"github.com/negrel/paon/internal/tree"
)

type rLayout = Layout

type Root interface {
	tree.Root
	Layout
}

var _ tree.Root = &root{}

// root define the root node of the Widget tree.
type root struct {
	rLayout
	children Widget
}

func NewRoot(child Widget) *root {
	r := &root{
		rLayout: newLayout(tree.NewRoot(child)),
	}

	return r
}

func (r *root) Render(screen geometry.Size) draw.Patch {
	canvas := draw.MakePatch(geometry.Rect(0, 0, screen.Width(), screen.Height()))
	ctx := render.MakeContext(r, &canvas)

	//r.Layout(ctx)
	//r.Draw(ctx)

	return canvas
}
