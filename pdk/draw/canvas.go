package draw

import (
	stdcontext "context"

	"github.com/negrel/paon/internal/geometry"
)

// Canvas define a bounded drawing surface.
type Canvas interface {
	// Bounds returns the bounds of this Canvas.
	Bounds() geometry.Rectangle

	// Get returns the cell at the given position.
	// If the position is out of bound, an Cell with default value
	// is returned.
	Get(geometry.Point) Cell

	// Set sets the cell at the given position.
	Set(geometry.Point, Cell)

	// NewContext creates a new Context that can draw on this Canvas within the given bounds.
	NewContext(ctx stdcontext.Context, bounds geometry.Rectangle) *Context
}
