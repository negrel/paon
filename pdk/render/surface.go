package render

import (
	"bytes"
	"io"
	"math"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
)

var _ draw.Surface = BufferSurface{}

// BufferSurface is an implementation of the Canvas interface.
type BufferSurface struct {
	cells  []draw.Cell
	bounds geometry.Rectangle
}

// NewBufferSurface returns a new BufferSurface with the given bounds.
func NewBufferSurface(bounds geometry.Rectangle) BufferSurface {
	return BufferSurface{
		cells:  make([]draw.Cell, bounds.Area()),
		bounds: bounds,
	}
}

// Bounds implements the Canvas interface.
func (bs BufferSurface) Bounds() geometry.Rectangle {
	return bs.bounds
}

func (bs BufferSurface) index(v2 geometry.Vec2D) int {
	return v2.Y()*bs.bounds.Width() + v2.X()
}

// Get implements the Canvas interface.
func (bs BufferSurface) Get(v2 geometry.Vec2D) draw.Cell {
	return bs.get(v2)
}

func (bs BufferSurface) get(v2 geometry.Vec2D) draw.Cell {
	v2 = v2.Add(bs.bounds.Min)
	if bs.bounds.Contains(v2) {
		return bs.cells[bs.index(v2)]
	}

	return draw.ZeroCell()
}

// Set implements the Canvas interface.
func (bs BufferSurface) Set(v2 geometry.Vec2D, cell draw.Cell) {
	bs.set(v2, cell)
}

func (bs BufferSurface) set(v2 geometry.Vec2D, cell draw.Cell) {
	v2 = v2.Add(bs.bounds.Min)

	if bs.bounds.Contains(v2) {
		bs.cells[bs.index(v2)] = cell
	}
}

// Dump the layer to the given io writer
func (bs BufferSurface) Dump(w io.Writer) {
	var buf bytes.Buffer

	for i, cell := range bs.cells {
		if i%bs.bounds.Width() == 0 {
			buf.WriteRune('\n')
			w.Write(buf.Bytes())
			buf.Reset()
		}
		buf.WriteRune(cell.Content)
	}
}

// Region size for infinite layers
const (
	RegionWidth  = 1024
	RegionHeight = 1024
)

var _ draw.Surface = InfiniteSurface{}

// InfiniteSurface define a surface with maximum bounds size.
type InfiniteSurface struct {
	origin geometry.Vec2D

	regions map[int]map[int]BufferSurface
}

// NewInfiniteSurface ...
func NewInfiniteSurface(origin geometry.Vec2D) InfiniteSurface {
	return InfiniteSurface{
		origin:  origin,
		regions: make(map[int]map[int]BufferSurface),
	}
}

// Bounds implement the draw.Surface interface.
func (is InfiniteSurface) Bounds() geometry.Rectangle {
	return geometry.Rectangle{
		Min: is.origin,
		Max: is.origin.Add(geometry.NewVec2D(math.MaxInt, math.MaxInt)),
	}
}

// Get implement the draw.Surface interface.
func (is InfiniteSurface) Get(v2 geometry.Vec2D) draw.Cell {
	return is.GetSurface(v2).Get(v2)
}

// GetSurface returns the surface containing the given Vec2D. This method is here
// for performance purpose. If you're going to draw within a single region only, you may
// want to use this method to avoid surface lookup on each Get/Set call.
func (is InfiniteSurface) GetSurface(v2 geometry.Vec2D) BufferSurface {
	var row map[int]BufferSurface
	var layer BufferSurface

	if r, ok := is.regions[v2.Y()/RegionHeight]; ok {
		row = r
	} else {
		row = make(map[int]BufferSurface)
		is.regions[v2.Y()/RegionHeight] = row
	}

	if l, ok := row[v2.X()/RegionWidth]; ok {
		layer = l
	} else {
		layer = NewBufferSurface(geometry.Rectangle{
			Min: geometry.Vec2D{},
			Max: geometry.NewVec2D(RegionWidth, RegionHeight),
		})
		row[v2.X()/RegionWidth] = layer
	}

	return layer
}

// Set implement the draw.Surface interface.
func (is InfiniteSurface) Set(v2 geometry.Vec2D, c draw.Cell) {
	is.GetSurface(v2).Set(v2, c)
}
