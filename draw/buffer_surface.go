package draw

import (
	"bytes"
	"io"

	"github.com/negrel/paon/geometry"
)

var _ Surface = BufferSurface{}

// BufferSurface is a dumpable in memory implementation of the Surface interface.
type BufferSurface struct {
	cells  []Cell
	bounds geometry.Rectangle
}

// NewBufferSurface returns a new BufferSurface with the given bounds.
func NewBufferSurface(size geometry.Size) BufferSurface {
	bounds := geometry.Rectangle{Origin: geometry.Vec2D{}, RectSize: size}

	return BufferSurface{
		cells:  make([]Cell, bounds.Area()),
		bounds: bounds,
	}
}

// Bounds implements the Surface interface.
func (bs BufferSurface) Size() geometry.Size {
	return bs.bounds.Size()
}

func (bs BufferSurface) index(v2 geometry.Vec2D) int {
	return v2.Y*bs.bounds.Size().Width + v2.X
}

// Get implements the Surface interface.
func (bs BufferSurface) Get(v2 geometry.Vec2D) Cell {
	return bs.get(v2)
}

func (bs BufferSurface) get(v2 geometry.Vec2D) Cell {
	v2 = v2.Add(bs.bounds.TopLeft())
	if bs.bounds.Contains(v2) {
		return bs.cells[bs.index(v2)]
	}

	return Cell{}
}

// Set implements the Surface interface.
func (bs BufferSurface) Set(v2 geometry.Vec2D, cell Cell) {
	bs.set(v2, cell)
}

func (bs BufferSurface) set(v2 geometry.Vec2D, cell Cell) {
	v2 = v2.Add(bs.bounds.TopLeft())

	if bs.bounds.Contains(v2) {
		bs.cells[bs.index(v2)] = cell
	}
}

// Dump the layer to the given io writer
func (bs BufferSurface) Dump(w io.Writer) error {
	var buf bytes.Buffer

	for i, cell := range bs.cells {
		if i != 0 && i%bs.bounds.Size().Width == 0 {
			buf.WriteRune('\n')
			_, err := w.Write(buf.Bytes())
			if err != nil {
				return err
			}
			buf.Reset()
		}

		if cell.Content != 0 {
			buf.WriteRune(cell.Content)
		} else {
			buf.WriteRune(' ')
		}
	}

	buf.WriteRune('\n')
	_, err := w.Write(buf.Bytes())
	return err
}
