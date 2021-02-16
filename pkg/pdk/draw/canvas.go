package draw

import "github.com/negrel/paon/internal/geometry"

type Canvas interface {
	// Bounds returns the bounds of this Canvas.
	Bounds() geometry.Rectangle

	// Get returns the cell at the given position.
	// A nil pointer is returned if the given position is not
	// within this Canvas bounds.
	Get(geometry.Point) *Cell

	// SubCanvas returns a new Canvas sharing the same underlying
	// Cell but with different bounds.
	SubCanvas(bounds geometry.Rectangle) Canvas

	// Draw applies the Drawer in this Canvas.
	Draw(Drawer)
}
