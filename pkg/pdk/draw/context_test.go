package draw

import (
	"github.com/golang/mock/gomock"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/render"
	"github.com/negrel/paon/pkg/pdk/styles/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext_Commit(t *testing.T) {
	ok := false

	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Lock()
	canvas.EXPECT().Unlock()

	ctx := newContext(canvas, geometry.Rectangle{})
	ctx.ops = append(ctx.ops, func(Canvas) {
		ok = true
	})
	assert.False(t, ok)

	ctx.Commit()
	assert.True(t, ok)
}

func TestContext_Commit_Twice(t *testing.T) {
	counter := 0

	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Lock().Times(2)
	canvas.EXPECT().Unlock().Times(2)

	ctx := newContext(canvas, geometry.Rectangle{})
	ctx.ops = append(ctx.ops, func(Canvas) {
		counter++
	})
	ctx.Commit()
	ctx.Commit()

	assert.Equal(t, 1, counter)
}

func TestContext_SetFillColor(t *testing.T) {
	ctx := newContext(nil, geometry.Rectangle{})
	red := value.ColorFromHex(0xFF0000)
	ctx.SetFillColor(red)

	assert.Equal(t, red, ctx.FillColor())
}

func TestContext_FillTextH(t *testing.T) {
	fillRect := geometry.Rect(0, 0, 20, 20)
	fillColor := value.ColorFromRGB(255, 0, 0)

	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Bounds().Return(fillRect).AnyTimes()
	canvas.EXPECT().Lock()
	canvas.EXPECT().Unlock()

	cells := make([]*render.Cell, 0, fillRect.Height()*fillRect.Width())
	for i := fillRect.Min.Y(); i < fillRect.Max.Y(); i++ {
		for j := fillRect.Min.X(); j < fillRect.Max.X(); j++ {
			cell := &render.Cell{}
			cells = append(cells, cell)
			canvas.EXPECT().Get(gomock.Eq(geometry.Pt(j, i))).Return(cell).AnyTimes()
		}
	}

	ctx := newContext(canvas, canvas.Bounds())
	ctx.SetFillColor(fillColor)
	textOrigin := geometry.Pt(9, 5)
	text := "Hello world"
	ctx.FillTextH(textOrigin, text)
	ctx.Commit()

	for i, char := range text {
		index := textOrigin.Y()*fillRect.Width() + textOrigin.X() + i
		assert.Equal(t, char, cells[index].Content)
	}
}

func TestContext_FillTextV(t *testing.T) {
	fillRect := geometry.Rect(0, 0, 20, 20)
	fillColor := value.ColorFromRGB(255, 0, 0)

	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Bounds().Return(fillRect).AnyTimes()
	canvas.EXPECT().Lock()
	canvas.EXPECT().Unlock()

	cells := make([]*render.Cell, 0, fillRect.Height()*fillRect.Width())
	for i := fillRect.Min.Y(); i < fillRect.Max.Y(); i++ {
		for j := fillRect.Min.X(); j < fillRect.Max.X(); j++ {
			cell := &render.Cell{}
			cells = append(cells, cell)
			canvas.EXPECT().Get(gomock.Eq(geometry.Pt(j, i))).Return(cell).AnyTimes()
		}
	}

	ctx := newContext(canvas, canvas.Bounds())
	ctx.SetFillColor(fillColor)
	textOrigin := geometry.Pt(9, 5)
	text := "Hello world"
	ctx.FillTextV(textOrigin, text)
	ctx.Commit()

	for i, char := range text {
		index := (textOrigin.Y()+i)*fillRect.Width() + textOrigin.X()
		assert.Equal(t, char, cells[index].Content)
	}
}

func TestContext_FillRectangle(t *testing.T) {
	fillRect := geometry.Rect(5, 7, 7, 9)
	fillColor := value.ColorFromRGB(255, 0, 0)

	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Bounds().Return(fillRect).AnyTimes()
	canvas.EXPECT().Lock()
	canvas.EXPECT().Unlock()

	cells := make([]*render.Cell, 0, fillRect.Height()*fillRect.Width())
	for i := fillRect.Min.Y(); i < fillRect.Max.Y(); i++ {
		for j := fillRect.Min.X(); j < fillRect.Max.X(); j++ {
			cell := &render.Cell{}
			cells = append(cells, cell)
			canvas.EXPECT().Get(gomock.Eq(geometry.Pt(j, i))).Return(cell)
		}
	}

	ctx := newContext(canvas, canvas.Bounds())
	ctx.SetFillColor(fillColor)
	ctx.FillRectangle(fillRect)

	for _, cell := range cells {
		assert.NotEqual(t, fillColor, cell.Style.Background)
	}

	ctx.Commit()

	for _, cell := range cells {
		assert.Equal(t, fillColor, cell.Style.Background)
	}
}

func TestContext_FillRectangle_OverText(t *testing.T) {
	fillRect := geometry.Rect(5, 7, 7, 9)
	fillColor := value.ColorFromRGB(255, 0, 0)

	ctrl := gomock.NewController(t)
	canvas := NewMockCanvas(ctrl)
	canvas.EXPECT().Bounds().Return(fillRect).AnyTimes()
	canvas.EXPECT().Lock()
	canvas.EXPECT().Unlock()

	cells := make([]*render.Cell, 0, fillRect.Height()*fillRect.Width())
	for i := fillRect.Min.Y(); i < fillRect.Max.Y(); i++ {
		for j := fillRect.Min.X(); j < fillRect.Max.X(); j++ {
			cell := &render.Cell{Content: rune(i + j)}
			cells = append(cells, cell)
			canvas.EXPECT().Get(gomock.Eq(geometry.Pt(j, i))).Return(cell)
		}
	}

	ctx := newContext(canvas, canvas.Bounds())
	ctx.SetFillColor(fillColor)
	ctx.FillRectangle(fillRect)
	ctx.Commit()

	for _, cell := range cells {
		assert.Equal(t, fillColor, cell.Style.Background)
		assert.Equal(t, rune(0), cell.Content)
	}
}
