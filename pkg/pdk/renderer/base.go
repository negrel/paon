package renderer

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

// Frame define a Rectangle with a rectangular hole in it.
type Frame struct {
	Outer geometry.Rectangle
	Inner geometry.Rectangle
}

// Base is the base used for all renderers. Base have many useful method to avoid to
// rewriting same code for multiple renderers.
type Base struct {
	Margin  Frame
	Padding Frame
}

func (b *Base) computeObjectHeightOrWidth(obj render.Object, p, min, max property.ID) int {
	result := -1
	style := obj.Style()

	if p := style.Get(p); p != nil {
		result = p.(property.Unit).Value
	}

	if max := style.Get(min); max != nil {
		maxR := max.(property.Unit).Value
		result = math.Min(result, maxR)
	}

	if min := style.Get(max); min != nil {
		minR := min.(property.Unit).Value
		result = math.Max(result, minR)
	}
	return result
}

// ComputeObjectWidth returns the width of the given render.Object based on it's style.
func (b *Base) ComputeObjectWidth(obj render.Object) int {
	return b.computeObjectHeightOrWidth(obj, property.IDWidth, property.IDMinWidth, property.IDMaxWidth)
}

// ComputeObjectHeight returns the height of the given render.Object based on it's style.
func (b *Base) ComputeObjectHeight(obj render.Object) int {
	return b.computeObjectHeightOrWidth(obj, property.IDHeight, property.IDMinHeight, property.IDMaxHeight)
}

// ComputeWidth computes the width for the given render.Context based on
// the render.Object style and the render.Constraint.
func (b *Base) ComputeWidth(ctx render.Context) int {
	w := b.ComputeObjectWidth(ctx.Object)
	w = math.Max(w, ctx.Constraint.Min.Width())
	w = math.Min(w, ctx.Constraint.Max.Width())

	return w
}

// ComputeHeight computes the height for the given render.Context based on
// the render.Object style and the render.Constraint.
func (b *Base) ComputeHeight(ctx render.Context) int {
	h := b.ComputeObjectHeight(ctx.Object)
	h = math.Max(h, ctx.Constraint.Min.Height())
	h = math.Min(h, ctx.Constraint.Max.Height())

	return h
}

// Draw draws on the given draw.Canvas
func (b *Base) Draw(canvas draw.Canvas, from, to geometry.Point, fn func(point geometry.Point, cell *draw.Cell)) {
	for i := from.Y(); i < to.Y(); i++ {
		for j := from.X(); j < to.Y(); j++ {
			pt := geometry.Pt(j, i)
			fn(pt, canvas.Get(pt))
		}
	}
}

// ApplyMarginLeft applies the given left margin of the canvas.
func (b *Base) ApplyMarginLeft(canvas draw.Canvas, size int) {
	b.Margin.Outer.Min = canvas.Bounds.Min
	b.Margin.Inner.Min = canvas.Bounds.Min.Add(geometry.Pt(size, 0))

	canvas.Bounds.Min = b.Margin.Inner.Min
}

// ApplyMarginRight applies the given right margin of the canvas.
func (b *Base) ApplyMarginRight(canvas draw.Canvas, size int) {
	b.Margin.Outer.Max = canvas.Bounds.Max
	b.Margin.Inner.Max = canvas.Bounds.Max.Add(geometry.Pt(-size, 0))

	canvas.Bounds.Max = b.Margin.Inner.Max
}

// ApplyMarginX applies the given margin on the left and the right of the canvas.
func (b *Base) ApplyMarginX(canvas draw.Canvas, size int) {
	b.ApplyMarginLeft(canvas, size)
	b.ApplyMarginRight(canvas, size)
}

// ApplyMarginTop applies the given top margin of the canvas.
func (b *Base) ApplyMarginTop(canvas draw.Canvas, size int) {
	b.Margin.Outer.Min = canvas.Bounds.Min
	b.Margin.Inner.Min = canvas.Bounds.Max.Add(geometry.Pt(0, size))

	canvas.Bounds.Min = b.Margin.Inner.Min
}

// ApplyMarginBottom applies the given bottom margin of the canvas.
func (b *Base) ApplyMarginBottom(canvas draw.Canvas, size int) {
	b.Margin.Outer.Max = canvas.Bounds.Max
	b.Margin.Inner.Max = canvas.Bounds.Max.Add(geometry.Pt(0, -size))

	canvas.Bounds.Max = b.Margin.Inner.Max

}

// ApplyMarginY applies the given margin on the top and the bottom of the canvas.
func (b *Base) ApplyMarginY(canvas draw.Canvas, size int) {
	b.ApplyMarginTop(canvas, size)
	b.ApplyMarginBottom(canvas, size)
}

// ApplyMargin applies the given margin on the left, top, right and the bottom of the canvas.
func (b *Base) ApplyMargin(canvas draw.Canvas, size int) {
	b.ApplyMarginX(canvas, size)
	b.ApplyMarginY(canvas, size)
}
