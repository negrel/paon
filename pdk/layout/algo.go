package layout

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Algo define a flow algorithm used to compute
// a BoxModel based on the given constraint.
type Algo interface {
	Layout(Constraint) BoxModel
}

var _ Algo = AlgoFn(nil)

type AlgoFn func(Constraint) BoxModel

func (af AlgoFn) Layout(c Constraint) BoxModel {
	return af(c)
}

// Block is a basic layout algorithm that returns a BoxModel using the
// width, height properties (min, max) and the Constraint.
// Sometimes one of those properties aren't defined and the fallback Algorithm is used.
// This algorithm must be wrapped in a function to be used as an Algorithm.
func Block(style styles.Style, constraint Constraint, fallback Algo) BoxModel {
	assert.NotNil(style)

	minWidth, maxWidth := ComputeMinMaxWidth(style, constraint)
	minHeight, maxHeight := ComputeMinMaxHeight(style, constraint)

	widthProp, hasWidth := GetUnitProp(style, property.Width())
	if hasWidth {
		width := constraint.ToCellUnit(widthProp.Value)
		minWidth = math.Max(minWidth, width)
		maxWidth = math.Min(maxWidth, width)
	}

	heightProp, hasHeight := GetUnitProp(style, property.Height())
	if hasHeight {
		height := constraint.ToCellUnit(heightProp.Value)
		minHeight = math.Max(minHeight, height)
		maxHeight = math.Min(maxHeight, height)
	}

	if isNotSized := !hasHeight || !hasWidth; isNotSized {
		constraint = constraint.SetMin(
			geometry.Rectangle{
				Min: constraint.Min.Min,
				Max: constraint.Min.Min.Add(geometry.Pt(minWidth, minHeight)),
			},
		).SetMax(
			geometry.Rectangle{
				Min: constraint.Max.Min,
				Max: constraint.Max.Min.Add(geometry.Pt(maxWidth, maxHeight)),
			},
		)

		return fallback.Layout(constraint)
	}

	// At this stage minWidth == maxWidth && minHeight == maxHeight
	box := NewBox(geometry.Rectangle{
		Min: constraint.Min.Min,
		Max: constraint.Min.Min.Add(geometry.Pt(minWidth, maxWidth)),
	})
	box.ApplyMargin(style)
	box.ApplyBorder(style)
	box.ApplyPadding(style)

	return box
}
