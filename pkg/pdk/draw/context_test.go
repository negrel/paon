package draw

import (
	"github.com/golang/mock/gomock"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/render"
	"github.com/negrel/paon/pkg/pdk/styles/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext_SetFillColor(t *testing.T) {
	ctx := newContext(nil, geometry.Rectangle{})
	red := value.ColorFromHex(0xFF0000)
	ctx.SetFillColor(red)

	assert.Equal(t, red, ctx.FillColor())
}

func TestContext_FillRectangle(t *testing.T) {
	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Bounds().Return(geometry.Rect(0, 0, 20, 20)).AnyTimes()

	fillRect := geometry.Rect(5, 7, 7, 9)
	fillColor := value.ColorFromRGB(255, 0, 0)

	cells := make([]*render.Cell, 0, 20)
	for i := fillRect.Min.X(); i < fillRect.Max.X(); i++ {
		for j := fillRect.Min.Y(); j < fillRect.Max.Y(); j++ {
			cell := &render.Cell{}
			cells = append(cells, cell)
			canvas.EXPECT().Get(gomock.Eq(geometry.Pt(i, j))).Return(cell)
		}
	}

	ctx := newContext(canvas, canvas.Bounds())
	ctx.SetFillColor(fillColor)
	ctx.FillRectangle(fillRect)
	ctx.Commit()

	for _, cell := range cells {
		assert.Equal(t, fillColor, cell.Style.Background)
	}
}

func TestContext_Commit(t *testing.T) {
	ok := false

	ctx := newContext(nil, geometry.Rectangle{})
	ctx.ops = append(ctx.ops, func(Canvas) {
		ok = true
	})
	assert.False(t, ok)

	ctx.Commit()
	assert.True(t, ok)
}

func TestContext_DoubleCommit(t *testing.T) {
	counter := 0

	ctx := newContext(nil, geometry.Rectangle{})
	ctx.ops = append(ctx.ops, func(Canvas) {
		counter++
	})
	ctx.Commit()
	ctx.Commit()

	assert.Equal(t, 1, counter)
}
