package styles

import (
	"github.com/negrel/paon/draw"
)

// ComputedStyle contains styling properties used for rendering.
type ComputedStyle struct {
	MarginStyle MarginPaddingStyle

	BordersStyle BordersStyle

	PaddingStyle MarginPaddingStyle

	draw.CellStyle

	ExtrasStyle any
}

// Compute implements Style.
func (cs ComputedStyle) Compute() ComputedStyle {
	return cs
}

// MarginPaddingStyle define margin and padding style properties.
type MarginPaddingStyle struct {
	Left, Top, Right, Bottom int
}

// BordersStyle define borders style properties.
type BordersStyle struct {
	Top    BorderSide
	Bottom BorderSide
	Left   BorderSide
	Right  BorderSide
}
