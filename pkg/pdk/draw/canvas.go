package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/render"
)

type Canvas interface {
	// Bounds returns the bounds of this Canvas.
	Bounds() geometry.Rectangle

	// Get returns the cell at the given position.
	// A nil pointer is returned if the given position is not
	// within this Canvas bounds.
	Get(geometry.Point) *render.Cell

	// SubCanvas returns a new Canvas sharing the same underlying
	// render.Cell but with different bounds.
	SubCanvas(bounds geometry.Rectangle) Canvas

	// Draw applies the Drawer in this Canvas.
	Draw(Drawer)

	// Patch returns a render.Patch object ready to be renderer.
	Patch() render.Patch

	// NewContext create a new context to draw on this Canvas within the given bounds.
	NewContext(bounds geometry.Rectangle) Context
}
