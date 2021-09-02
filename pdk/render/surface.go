package render

import (
	"bytes"
	"io"
	"math"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
)

func makeBuf(rect geometry.Rectangle) []draw.Cell {
	return make([]draw.Cell, rect.Area())
}

var _ draw.Canvas = BufferSurface{}

// BufferSurface is an implementation of the Canvas interface.
type BufferSurface struct {
	cells  []draw.Cell
	bounds geometry.Rectangle
}

// NewBufferSurface returns a new BufferSurface with the given bounds.
func NewBufferSurface(bounds geometry.Rectangle) BufferSurface {
	return BufferSurface{
		cells:  makeBuf(bounds),
		bounds: bounds,
	}
}

// Bounds implements the Canvas interface.
func (bs BufferSurface) Bounds() geometry.Rectangle {
	return bs.bounds
}

func (bs BufferSurface) index(pt geometry.Point) int {
	return pt.Y()*bs.bounds.Width() + pt.X()
}

// Get implements the Canvas interface.
func (bs BufferSurface) Get(pt geometry.Point) draw.Cell {
	return bs.get(pt)
}

func (bs BufferSurface) get(pt geometry.Point) draw.Cell {
	pt = pt.Add(bs.bounds.Min)
	if bs.bounds.Contains(pt) {
		return bs.cells[bs.index(pt)]
	}

	return draw.ZeroCell()
}

// Set implements the Canvas interface.
func (bs BufferSurface) Set(pt geometry.Point, cell draw.Cell) {
	bs.set(pt, cell)
}

func (bs BufferSurface) set(pt geometry.Point, cell draw.Cell) {
	pt = pt.Add(bs.bounds.Min)

	if bs.bounds.Contains(pt) {
		bs.cells[bs.index(pt)] = cell
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

var _ draw.Canvas = InfiniteSurface{}

// InfiniteSurface define a surface with maximum bounds size.
type InfiniteSurface struct {
	origin *geometry.Point

	regions map[int]map[int]BufferSurface
}

// NewInfiniteSurface ...
func NewInfiniteSurface(origin geometry.Point) InfiniteSurface {
	return InfiniteSurface{
		origin:  &origin,
		regions: make(map[int]map[int]BufferSurface),
	}
}

// Bounds implement the draw.Canvas interface.
func (is InfiniteSurface) Bounds() geometry.Rectangle {
	return geometry.Rectangle{
		Min: *is.origin,
		Max: is.origin.Add(geometry.Pt(math.MaxInt, math.MaxInt)),
	}
}

// Get implement the draw.Canvas interface.
func (is InfiniteSurface) Get(pt geometry.Point) draw.Cell {
	return is.GetSurface(pt).Get(pt)
}

// GetSurface returns the surface containing the given point. This method is here
// for performance purpose. If you're going to draw within a single region only, you may
// want to use this method to avoid surface lookup on each Get/Set call.
func (is InfiniteSurface) GetSurface(pt geometry.Point) BufferSurface {
	var row map[int]BufferSurface
	var layer BufferSurface

	if r, ok := is.regions[pt.Y()%RegionHeight]; ok {
		row = r
	} else {
		row = make(map[int]BufferSurface)
		is.regions[pt.Y()%RegionHeight] = row
	}

	if l, ok := row[pt.X()%RegionWidth]; ok {
		layer = l
	} else {
		layer = NewBufferSurface(geometry.Rectangle{
			Min: geometry.Point{},
			Max: geometry.Pt(RegionWidth, RegionHeight),
		})
		row[pt.X()%RegionWidth] = layer
	}

	return layer
}

// Set implement the draw.Canvas interface.
func (is InfiniteSurface) Set(pt geometry.Point, c draw.Cell) {
	is.GetSurface(pt).Set(pt, c)
}
