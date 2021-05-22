package draw

import (
	"github.com/negrel/paon/internal/geometry"
)

type Canvas interface {
	// Bounds returns the bounds of this Canvas.
	Bounds() geometry.Rectangle

	// Get returns the cell at the given position.
	// If the position is out of bound, an Cell with default value
	// is returned.
	Get(geometry.Point) Cell

	// Set sets the cell at the given position.
	Set(geometry.Point, Cell)

	// NewContext creates a new Context that can on this Canvas within the given bounds.
	NewContext(bounds geometry.Rectangle) Context
}
