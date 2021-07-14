package manager

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// Block is a layout helper that can be used by a Manager.
// Block updates the size Constraint using the min/max width/height
// properties while respecting the original Constraint.
// If the style is fully sized (has a defined width and height)
// a BoxModel using these properties and the margins, borders and
// paddings props is returned.
// Otherwise, the fallback manager is called. The fallback should
// also applies margins, borders and paddings in order to have consistent
// layout.
func Block(style styles.Style, constraint layout.Constraint, fallback layout.Manager) layout.BoxModel {
	assert.NotNil(style)

	minWidth, maxWidth := layout.MinMaxWidth(style, constraint)
	minHeight, maxHeight := layout.MinMaxHeight(style, constraint)

	widthProp, hasWidth := layout.UnitProp(style, property.WidthID())
	if hasWidth {
		width := constraint.ToCellUnit(widthProp.Value)
		minWidth = math.Max(minWidth, width)
		maxWidth = math.Min(maxWidth, width)
	}

	heightProp, hasHeight := layout.UnitProp(style, property.HeightID())
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
	box := layout.NewBox(geometry.Rectangle{
		Min: constraint.Min.Min,
		Max: constraint.Min.Min.Add(geometry.Pt(maxWidth, maxHeight)),
	})
	box.ApplyMargin(style)
	box.ApplyBorder(style)
	box.ApplyPadding(style)

	return box
}
