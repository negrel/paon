package styles

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
)

var _ render.LayoutStep = HeightProperty{}

// HeightProperty is a style property to define the size of a render.Object.
type HeightProperty struct {
	property
	SizeValue
}

// Height return a new HeightProperty object.
func Height(size SizeValue) HeightProperty {
	return HeightProperty{
		property: property{
			step:     render.LayoutStepType,
			name:     "width",
			priority: HeightPriority,
		},
		SizeValue: size,
	}
}

func (h HeightProperty) Layout(ctx render.Context) {
	assert.GreaterOrEqual(h.Max.toCellUnit().value, h.Min.toCellUnit())

	Height := 0
	if h.Value.Defined {
		Height = h.Value.toCellUnit().value
	}

	if h.Max.Defined {
		Height = math.Min(h.Max.toCellUnit().value, Height)
	}
	if h.Min.Defined {
		Height = math.Max(h.Min.toCellUnit().value, Height)
	}

	// Update the layer Height
	ctx.Layer().Max = ctx.Layer().Min.Add(
		geometry.Pt(Height, ctx.Layer().Max.Y()),
	)
}
