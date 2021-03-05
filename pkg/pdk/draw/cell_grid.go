package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/render"
)

var _ Canvas = CellGrid{}

type CellGrid struct {
	bounds geometry.Rectangle
	grid   [][]*render.Cell
}

// MakeCellGrid return a new CellGrid Canvas.
func MakeCellGrid(bounds geometry.Rectangle) CellGrid {
	return CellGrid{
		bounds: bounds,
		grid:   make([][]*render.Cell, bounds.Height()),
	}
}

func (cg CellGrid) get(pos geometry.Point) *render.Cell {
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
func (cg CellGrid) Get(pos geometry.Point) *render.Cell {
	if !cg.bounds.Contains(pos) {
		return nil
	}

	return cg.get(pos)
}

// Bounds implements the Canvas interface.
func (cg CellGrid) Bounds() geometry.Rectangle {
	return cg.bounds
}

// SubCanvas implements the Canvas interface.
func (cg CellGrid) SubCanvas(bounds geometry.Rectangle) Canvas {
	return CellGrid{
		bounds: bounds,
		grid:   cg.grid,
	}
}

// Patch implements the Canvas interface.
func (cg CellGrid) Patch() render.Patch {
	return render.Patch{
		Pos:   cg.bounds.Min,
		Cells: cg.grid,
	}
}

func (cg CellGrid) NewContext(bounds geometry.Rectangle) Context {
	return newContext(cg, bounds)
}
