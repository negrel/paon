package render

import (
	"bytes"
	"io"
	"math"

	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
)

// Large array of zero draw.Cell to copy from using the built-in function
var zeroCellBuf = [8192]draw.Cell{}

type size = geometry.Size

var _ draw.Surface = BufferSurface{}

// BufferSurface is an implementation of the Canvas interface.
type BufferSurface struct {
	size
	cells []draw.Cell
}

// NewBufferSurface returns a new BufferSurface with the given bounds.
func NewBufferSurface(size geometry.Size) BufferSurface {
	return BufferSurface{
		size:  size,
		cells: make([]draw.Cell, size.Width()*size.Height()),
	}
}

// Size implements the geometry.Sized interface.
func (bs BufferSurface) Size() geometry.Size {
	return bs.size
}

// Resize resizes the BufferSurface to the given size and zero all cells.
func (bs BufferSurface) Resize(size geometry.Size) BufferSurface {
	bs.size = size

	newArea := size.Width() * size.Height()

	if newArea-cap(bs.cells) > 0 {
		bs.cells = make([]draw.Cell, newArea)
		return bs
	}

	// Resize the slice
	area := size.Width() * size.Height()
	bs.cells = bs.cells[:area]

	// Fill it with zero cell
	copied := 0
	for copied != len(bs.cells) {
		copied += copy(bs.cells, zeroCellBuf[:])
	}

	return bs
}

func (bs BufferSurface) bounds() geometry.Rectangle {
	return geometry.Rect(0, 0, bs.Size().Width(), bs.Size().Height())
}

func (bs BufferSurface) index(v2 geometry.Vec2D) int {
	return v2.Y()*bs.size.Width() + v2.X()
}

// Get implements the Canvas interface.
func (bs BufferSurface) Get(v2 geometry.Vec2D) draw.Cell {
	if bs.bounds().Contains(v2) {
		return bs.cells[bs.index(v2)]
	}

	return draw.Cell{}
}

// Set implements the Canvas interface.
func (bs BufferSurface) Set(v2 geometry.Vec2D, cell draw.Cell) {
	if bs.bounds().Contains(v2) {
		bs.cells[bs.index(v2)] = cell
	}
}

// Slice returns the underlying drawing cell buffer.
func (bs BufferSurface) Slice() []draw.Cell {
	return bs.cells
}

// Dump the layer to the given io writer
func (bs BufferSurface) Dump(w io.Writer) {
	var buf bytes.Buffer

	for i, cell := range bs.cells {
		if i%bs.size.Width() == 0 {
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

// Size implements the geometry.Sized interface.
func (is InfiniteSurface) Size() geometry.Size {
	return geometry.NewSize(math.MaxInt, math.MaxInt)
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
		layer = NewBufferSurface(geometry.NewSize(RegionWidth, RegionHeight))
		row[v2.X()/RegionWidth] = layer
	}

	return layer
}

// Set implement the draw.Surface interface.
func (is InfiniteSurface) Set(v2 geometry.Vec2D, c draw.Cell) {
	is.GetSurface(v2).Set(v2, c)
}
