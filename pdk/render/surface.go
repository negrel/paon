package render

import (
	"bytes"
	"io"

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
