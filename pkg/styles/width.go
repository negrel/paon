package styles

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
)

var _ render.LayoutStep = WidthProperty{}

// WidthProperty is a style property to define the size of a render.Object.
type WidthProperty struct {
	property
	SizeValue
}

// Width return a new WidthProperty object.
func Width(size SizeValue) WidthProperty {
	return WidthProperty{
		property: property{
			step:     render.LayoutStepType,
			name:     "width",
			priority: WidthPriority,
		},
		SizeValue: size,
	}
}

func (w WidthProperty) Layout(ctx render.Context) {
	width := ctx.Layer().Width()

	if w.Value.Defined {
		width = w.Value.toCellUnit().value
	}

	if w.Max.Defined {
		width = math.Min(w.Max.toCellUnit().value, width)
	}
	if w.Min.Defined {
		width = math.Max(w.Min.toCellUnit().value, width)
	}

	// Update the layer width
	ctx.Layer().Max = ctx.Layer().Min.Add(
		geometry.Pt(width, ctx.Layer().Max.Y()),
	)
}
