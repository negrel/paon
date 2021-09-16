package draw

import (
	"github.com/negrel/paon/internal/geometry"
)

// Surface define a bounded drawing surface.
type Surface interface {
	// Bounds returns the bounds of this Surface.
	Bounds() geometry.Rectangle

	// Get returns the cell at the given position.
	// If the position is out of bound, a Cell with default value
	// is returned.
	Get(geometry.Vec2D) Cell

	// Set sets the cell at the given position.
	// If the position is out of bounds, the cell is dropped.
	Set(geometry.Vec2D, Cell)
}

var _ Surface = SubSurface{}

// SubSurface define a subset of a Surface. It get/set cells
// relative to it's own bound.
type SubSurface struct {
	bounds  geometry.Rectangle
	surface Surface
}

// NewSubSurface wraps the given Surface to rebound it under the given subbounds.
func NewSubSurface(s Surface, subbounds geometry.Rectangle) SubSurface {
	if srf, ok := s.(SubSurface); ok {
		s = srf.surface
		subbounds = geometry.Rectangle{
			Min: srf.bounds.Min.Add(subbounds.Min),
			Max: srf.bounds.Max.Add(subbounds.Max),
		}
	}

	return SubSurface{
		bounds:  s.Bounds().Mask(subbounds),
		surface: s,
	}
}

// Bounds implement the Surface interface.
func (ss SubSurface) Bounds() geometry.Rectangle {
	return ss.bounds
}

// Get implements the Surface interface.
func (ss SubSurface) Get(v2 geometry.Vec2D) Cell {
	v2 = v2.Add(ss.bounds.Min)
	if ss.bounds.Contains(v2) {
		return ss.surface.Get(v2)
	}

	return ZeroCell()
}

// Set implements the Surface interface.
func (ss SubSurface) Set(v2 geometry.Vec2D, cell Cell) {
	v2 = v2.Add(ss.bounds.Min)
	ss.surface.Set(v2, cell)
}
