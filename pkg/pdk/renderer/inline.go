package renderer

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/render"
)

type inline struct {
	Base
}

func makeInline() render.Renderer {
	return inline{}
}

func (i inline) performLayout(ctx render.Context) {
	width := i.ComputeWidth(ctx)
	assert.GreaterOrEqual(width, 0, "widget width can't be a negative number")

	height := i.ComputeHeight(ctx)
	assert.GreaterOrEqual(height, 0, "widget height can't be a negative number")
}

// Layout implements the render.Renderer interface.
func (i inline) Layout(ctx render.Context) {

}

// Draw implements the render.Renderer interface.
func (i inline) Draw(ctx render.Context) {
	panic("implement me")
}
