package painter

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/draw"
)

type inline struct {
	Base
}

func makeInline() draw.Painter {
	return inline{}
}

func (i inline) performLayout(ctx draw.Context) {
	width := i.ComputeWidth(ctx)
	assert.GreaterOrEqual(width, 0, "widget width can't be a negative number")

	height := i.ComputeHeight(ctx)
	assert.GreaterOrEqual(height, 0, "widget height can't be a negative number")
}

// Layout implements the draw.Painter interface.
func (i inline) Layout(ctx draw.Context) {

}

// Draw implements the draw.Painter interface.
func (i inline) Draw(ctx draw.Context) {
	panic("implement me")
}
