package draw

import (
	"github.com/negrel/paon/internal/geometry"
)

// Canvas define a bounded drawing surface.
type Canvas interface {
	// Bounds returns the bounds of this Canvas.
	Bounds() geometry.Rectangle

	// Get returns the cell at the given position.
	// If the position is out of bound, a Cell with default value
	// is returned.
	Get(geometry.Point) Cell

	// Set sets the cell at the given position.
	// If the position is out of bounds, the cell is dropped.
	Set(geometry.Point, Cell)
}

var _ Canvas = SubCanvas{}

// SubCanvas define a subset of a Canvas. It get/set cells
// relative to it's own bound.
type SubCanvas struct {
	bounds geometry.Rectangle
	canvas Canvas
}

// NewSubCanvas wraps the given Canvas to rebound it under the given subbounds.
func NewSubCanvas(c Canvas, subbounds geometry.Rectangle) SubCanvas {
	if canvas, ok := c.(SubCanvas); ok {
		c = canvas.canvas
		subbounds = geometry.Rectangle{
			Min: canvas.bounds.Min.Add(subbounds.Min),
			Max: canvas.bounds.Max.Add(subbounds.Max),
		}
	}

	return SubCanvas{
		bounds: c.Bounds().Mask(subbounds),
		canvas: c,
	}
}

// Bounds implement the Canvas interface.
func (sc SubCanvas) Bounds() geometry.Rectangle {
	return sc.bounds
}

// Get implements the Canvas interface.
func (sc SubCanvas) Get(pt geometry.Point) Cell {
	pt = pt.Add(sc.bounds.Min)
	if sc.bounds.Contains(pt) {
		return sc.canvas.Get(pt)
	}

	return ZeroCell()
}

// Set implements the Canvas interface.
func (sc SubCanvas) Set(pt geometry.Point, cell Cell) {
	pt = pt.Add(sc.bounds.Min)
	sc.canvas.Set(pt, cell)
}
