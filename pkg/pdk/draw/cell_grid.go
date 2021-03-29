package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/render"
	"sync"
)

var _ Canvas = &CellGrid{}
var _ render.Patch = &CellGrid{}

type CellGrid struct {
	*sync.Mutex
	bounds geometry.Rectangle
	grid   [][]*render.Cell
}

// NewCellGrid return a new CellGrid Canvas.
func NewCellGrid(bounds geometry.Rectangle) *CellGrid {
	cg := &CellGrid{
		Mutex:  new(sync.Mutex),
		bounds: bounds,
		grid:   make([][]*render.Cell, bounds.Height()),
	}

	for i := 0; i < bounds.Height(); i++ {
		cg.grid[i] = make([]*render.Cell, bounds.Width())
	}

	return cg
}

func (cg *CellGrid) get(pos geometry.Point) *render.Cell {
	cg.Mutex.Lock()
	defer cg.Mutex.Unlock()

	row := cg.grid[pos.Y()]
	if row == nil {
		row = make([]*render.Cell, cg.bounds.Width())
		cg.grid[pos.Y()] = row
	}

	cell := row[pos.X()]
	if cell == nil {
		cell = &render.Cell{}
		row[pos.X()] = cell
	}

	return cell
}

// Get implements the Canvas interface.
func (cg *CellGrid) Get(pos geometry.Point) *render.Cell {
	if !cg.bounds.Contains(pos) {
		return nil
	}

	return cg.get(pos)
}

// Bounds implements the Canvas interface.
func (cg *CellGrid) Bounds() geometry.Rectangle {
	return cg.bounds
}

// SubCanvas implements the Canvas interface.
func (cg *CellGrid) SubCanvas(bounds geometry.Rectangle) Canvas {
	return &CellGrid{
		Mutex:  cg.Mutex,
		bounds: cg.bounds.Mask(bounds),
		grid:   cg.grid,
	}
}

// Patch implements the Canvas interface.
func (cg *CellGrid) Patch() render.Patch {
	return cg
}

// ForEachCell implements the render.Patch interface.
func (cg *CellGrid) ForEachCell(fn func(pos geometry.Point, cell *render.Cell)) {
	for i := cg.bounds.Min.Y(); i < cg.bounds.Max.Y(); i++ {
		for j := cg.bounds.Min.X(); j < cg.bounds.Max.X(); j++ {
			pos := geometry.Pt(j, i)
			fn(pos, cg.get(pos))
		}
	}
}

func (cg *CellGrid) NewContext(bounds geometry.Rectangle) Context {
	return newContext(cg, bounds)
}
