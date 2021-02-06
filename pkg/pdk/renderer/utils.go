package renderer

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

func computeObjectHeightOrWidth(obj render.Object, p, min, max property.ID) int {
	result := -1

	if p := obj.Get(p); p != nil {
		result = p.(property.Unit).Value
	}

	if max := obj.Get(min); max != nil {
		maxR := max.(property.Unit).Value
		result = math.Min(result, maxR)
	}

	if min := obj.Get(max); min != nil {
		minR := min.(property.Unit).Value
		result = math.Max(result, minR)
	}
	return result
}

// ComputeObjectWidth returns the width of the given render.Object based on it's style.
func ComputeObjectWidth(obj render.Object) int {
	return computeObjectHeightOrWidth(obj, property.IDWidth, property.IDMinWidth, property.IDMaxWidth)
}

// ComputeObjectHeight returns the height of the given render.Object based on it's style..
func ComputeObjectHeight(obj render.Object) int {
	return computeObjectHeightOrWidth(obj, property.IDHeight, property.IDMinHeight, property.IDMaxHeight)
}

// ComputeWidth computes the width for the given render.Context based on
// the render.Object style and the render.Constraint.
func ComputeWidth(ctx render.Context) int {
	w := ComputeObjectWidth(ctx.Object)
	w = math.Max(w, ctx.Constraint.Min.Width())
	w = math.Min(w, ctx.Constraint.Max.Width())

	return w
}

// ComputeHeight computes the height for the given render.Context based on
// the render.Object style and the render.Constraint.
func ComputeHeight(ctx render.Context) int {
	h := ComputeObjectHeight(ctx.Object)
	h = math.Max(h, ctx.Constraint.Min.Height())
	h = math.Min(h, ctx.Constraint.Max.Height())

	return h
}

// Draw draws on the given draw.Canvas
func Draw(canvas draw.Canvas, from, to geometry.Point, fn func(point geometry.Point, cell *draw.Cell)) {
	for i := from.Y(); i < to.Y(); i++ {
		for j := from.X(); j < to.Y(); j++ {
			pt := geometry.Pt(j, i)
			fn(pt, canvas.Get(pt))
		}
	}
}

// ApplyMarginLeft applies the given left margin of the canvas.
func ApplyMarginLeft(canvas draw.Canvas, size int) {
	canvas.Bounds.Min = canvas.Bounds.Min.Add(geometry.Pt(size, 0))
}

// ApplyMarginRight applies the given right margin of the canvas.
func ApplyMarginRight(canvas draw.Canvas, size int) {
	canvas.Bounds.Max = canvas.Bounds.Max.Add(geometry.Pt(-size, 0))
}

// ApplyMarginX applies the given margin on the left and the right of the canvas.
func ApplyMarginX(canvas draw.Canvas, size int) {
	ApplyMarginLeft(canvas, size)
	ApplyMarginRight(canvas, size)
}

// ApplyMarginTop applies the given top margin of the canvas.
func ApplyMarginTop(canvas draw.Canvas, size int) {
	canvas.Bounds.Min = canvas.Bounds.Max.Add(geometry.Pt(0, size))
}

// ApplyMarginBottom applies the given bottom margin of the canvas.
func ApplyMarginBottom(canvas draw.Canvas, size int) {
	canvas.Bounds.Max = canvas.Bounds.Max.Add(geometry.Pt(0, -size))
}

// ApplyMarginY applies the given margin on the top and the bottom of the canvas.
func ApplyMarginY(canvas draw.Canvas, size int) {
	ApplyMarginTop(canvas, size)
	ApplyMarginBottom(canvas, size)
}

// ApplyMargin applies the given margin on the left, top, right and the bottom of the canvas.
func ApplyMargin(canvas draw.Canvas, size int) {
	ApplyMarginX(canvas, size)
	ApplyMarginY(canvas, size)
}
