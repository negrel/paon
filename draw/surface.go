package draw

import (
	"github.com/negrel/paon/geometry"
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

var _ Surface = SubSurface{}

// SubSurface define a bounded wrapper around a Surface type.
// Every cells sets/gets from this SubSurface is relative to the
// origin of the SubSurface bounds.
type SubSurface struct {
	srf    Surface
	bounds geometry.Rectangle
}

// NewSubSurface returns a new SubSurface object that wraps the given
// Surface and use the given geometry.Rectangle as bounds.
func NewSubSurface(s Surface, bounds geometry.Rectangle) SubSurface {
	if srf, ok := s.(SubSurface); ok {
		s = srf.srf
		bounds = geometry.Rectangle{
			Min: srf.bounds.Min.Add(bounds.Min),
			Max: srf.bounds.Min.Add(bounds.Max),
		}
	}

	return SubSurface{
		srf:    s,
		bounds: bounds,
	}
}

// Size implements the geometry.Sized interface.
func (ss SubSurface) Size() geometry.Size {
	return ss.bounds.Size()
}

// Get implements the Surface interface.
func (ss SubSurface) Get(v2 geometry.Vec2D) Cell {
	v2 = v2.Add(ss.bounds.Min)
	return ss.srf.Get(v2)
}

// Set implements the Surface interface.
func (ss SubSurface) Set(v2 geometry.Vec2D, cell Cell) {
	v2 = v2.Add(ss.bounds.Min)
	ss.srf.Set(v2, cell)
}
