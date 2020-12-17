package styles

import (
	"github.com/negrel/debuggo/pkg/log"
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
			name:     "height",
			priority: HeightPriority,
		},
		SizeValue: size,
	}
}

func (h HeightProperty) Layout(ctx render.Context) {
	height := ctx.Layer().Height()

	if h.Value.Defined {
		height = h.Value.toCellUnit().value
	}

	if h.Max.Defined {
		height = math.Min(h.Max.toCellUnit().value, height)
	}
	if h.Min.Defined {
		height = math.Max(h.Min.toCellUnit().value, height)
	}

	// Update the layer height
	ctx.Layer().Max = ctx.Layer().Min.Add(
		geometry.Pt(ctx.Layer().Max.X(), height),
	)

	log.Debug("LAYER ", ctx.Layer())
}
