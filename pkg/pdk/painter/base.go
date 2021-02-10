package painter

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

// Frame define a Rectangle with a rectangular hole in it.
type Frame struct {
	Outer geometry.Rectangle
	Inner geometry.Rectangle
}

// Base is the base used for all renderers. Base have many useful method to avoid to
// rewriting same code for multiple renderers.
type Base struct {
}

func toCellUnit(unit property.Unit) value.Unit {
	return value.Unit{}
}

// UnitProp returns the value of the given property as a CellUnit value.
func UnitProp(style styles.Style, id property.ID) int {
	p := style.Get(id)
	if p == nil {
		return 0
	}

	unit, ok := p.(property.Unit)
	log.Debugf("property %s is not a property.Unit and is thus considered as a 0", id)
	if !ok {
		return 0
	}

	return toCellUnit(unit).Value
}

func computeObjectHeightOrWidth(obj draw.Object, p, min, max property.ID) int {
	result := -1
	style := obj.Style()

	result = UnitProp(style, p)
	result = math.Min(result, UnitProp(style, min))
	result = math.Max(result, UnitProp(style, max))

	return result
}

// ComputeObjectWidth returns the width of the given draw.Object based on it's style.
func ComputeObjectWidth(obj draw.Object) int {
	return computeObjectHeightOrWidth(obj, property.IDWidth, property.IDMinWidth, property.IDMaxWidth)
}

// ComputeObjectHeight returns the height of the given draw.Object based on it's style.
func ComputeObjectHeight(obj draw.Object) int {
	return computeObjectHeightOrWidth(obj, property.IDHeight, property.IDMinHeight, property.IDMaxHeight)
}

// ComputeWidth computes the width for the given draw.Context based on
// the draw.Object style and the draw.Constraint.
func ComputeWidth(ctx draw.Context) int {
	w := ComputeObjectWidth(ctx.Object)
	w = math.Max(w, ctx.Constraint.Min.Width())
	w = math.Min(w, ctx.Constraint.Max.Width())

	return w
}

// ComputeHeight computes the height for the given draw.Context based on
// the draw.Object style and the draw.Constraint.
func ComputeHeight(ctx draw.Context) int {
	h := ComputeObjectHeight(ctx.Object)
	h = math.Max(h, ctx.Constraint.Min.Height())
	h = math.Min(h, ctx.Constraint.Max.Height())

	return h
}

// Draw draws on the given render.Buffer
func Draw(canvas render.Buffer, from, to geometry.Point, fn func(point geometry.Point, cell *render.Cell)) {
	for i := from.Y(); i < to.Y(); i++ {
		for j := from.X(); j < to.Y(); j++ {
			pt := geometry.Pt(j, i)
			fn(pt, canvas.Get(pt))
		}
	}
}

// ApplyMargin applies the margins defined in the given styles.Style on the given
// geometry.Rectangle target and returns it.
func ApplyMargin(style styles.Style, target geometry.Rectangle) geometry.Rectangle {
	ml := UnitProp(style, property.IDMarginLeft)
	mt := UnitProp(style, property.IDMarginTop)
	target.Min = target.TopLeft().Add(geometry.Pt(ml, mt))

	mr := UnitProp(style, property.IDMarginRight)
	mb := UnitProp(style, property.IDMarginBottom)
	target.Max = target.BottomRight().Sub(geometry.Pt(mr, mb))

	return target
}

// ApplyBorder applies the margins defined in the given styles.Style on the given
// geometry.Rectangle target and returns it.
func ApplyBorder(style styles.Style, target geometry.Rectangle) geometry.Rectangle {
	bl := UnitProp(style, property.IDBorderLeft)
	bt := UnitProp(style, property.IDBorderTop)
	target.Min = target.TopLeft().Add(geometry.Pt(bl, bt))

	br := UnitProp(style, property.IDBorderRight)
	bb := UnitProp(style, property.IDBorderBottom)
	target.Max = target.BottomRight().Sub(geometry.Pt(br, bb))

	return target
}
