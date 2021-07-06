package layout

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Manager calculates a BoxModel based on the given constraint.
type Manager interface {
	Layout(Constraint) BoxModel
}

var _ Manager = ManagerFn(nil)

// ManagerFn is a simple function that implements the Manager interface.
type ManagerFn func(Constraint) BoxModel

// Layout implements the Manager interface.
func (mf ManagerFn) Layout(c Constraint) BoxModel {
	return mf(c)
}

// Block is a basic layout algorithm that returns a BoxModel using the
// width, height properties (min, max) and the Constraint.
// Sometimes one of those properties aren't defined and the fallback Algorithm is used.
// This algorithm must be wrapped in a function to be used as an Algorithm.
func Block(style styles.Style, constraint Constraint, fallback Manager) BoxModel {
	assert.NotNil(style)

	minWidth, maxWidth := computeMinMaxWidth(style, constraint)
	minHeight, maxHeight := computeMinMaxHeight(style, constraint)

	widthProp, hasWidth := getUnitProp(style, property.Width())
	if hasWidth {
		width := constraint.ToCellUnit(widthProp.Value)
		minWidth = math.Max(minWidth, width)
		maxWidth = math.Min(maxWidth, width)
	}

	heightProp, hasHeight := getUnitProp(style, property.Height())
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
