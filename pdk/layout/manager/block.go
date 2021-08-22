package manager

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/layout"
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
func Block(style styles.Style, constraint layout.Constraint, fallback layout.Manager) *layout.Box {
	assert.NotNil(style)

	constraint = constraint.NewFromStyle(style)
	width := constraint.MinSize.Width()
	height := constraint.MinSize.Height()

	widthProp, hasWidth := layout.UnitProp(style, property.WidthID())
	if hasWidth {
		width = constraint.ApplyOnWidthProp(widthProp)
	}

	heightProp, hasHeight := layout.UnitProp(style, property.HeightID())
	if hasHeight {
		height = constraint.ApplyOnHeightProp(heightProp)
	}

	if !hasWidth || !hasHeight {
		return fallback.Layout(constraint)
	}

	// At this stage minWidth == maxWidth && minHeight == maxHeight
	box := layout.NewBox(geometry.NewSize(width, height)).
		ApplyMargin(style).
		ApplyBorder(style).
		ApplyPadding(style)

	return box
}
