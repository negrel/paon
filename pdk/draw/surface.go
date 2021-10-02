package draw

import (
	"github.com/negrel/paon/internal/geometry"
)

// Surface define a bounded drawing surface.
type Surface interface {
	geometry.Sized

	// Get returns the cell at the given position.
	// If the position is out of bound, a Cell with default value
	// is returned.
	Get(geometry.Vec2D) Cell

	// Set sets the cell at the given position.
	// If the position is out of bounds, the cell is dropped.
	Set(geometry.Vec2D, Cell)
}
