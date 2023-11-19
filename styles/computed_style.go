package styles

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/widgets/border"
)

// ComputedStyle contains styling properties used for rendering.
type ComputedStyle struct {
	MarginStyle MarginPaddingStyle

	BorderStyle BorderStyle

	PaddingStyle MarginPaddingStyle
	draw.CellStyle

	ExtrasStyle any
}

type MarginPaddingStyle struct {
	Left, Top, Right, Bottom int
}

type BorderStyle struct {
	border.Border
	draw.CellStyle
}
