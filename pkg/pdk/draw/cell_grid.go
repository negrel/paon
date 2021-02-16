package draw

import "github.com/negrel/paon/internal/geometry"

var _ Canvas = CellGrid{}

type CellGrid struct {
	bounds geometry.Rectangle
	grid   [][]*Cell
}

func (cg CellGrid) get(pos geometry.Point) *Cell {
	row := cg.grid[pos.Y()]
	if row == nil {
		row = make([]*Cell, cg.bounds.Width())
		cg.grid[pos.Y()] = row
	}

	cell := row[pos.X()]
	if cell == nil {
		cell = &Cell{}
		row[pos.X()] = cell
	}

	return cell
}

// Get implements the Canvas interface.
func (cg CellGrid) Get(pos geometry.Point) *Cell {
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

// Draw implements the Canvas interface.
func (cg CellGrid) Draw(drawer Drawer) {
	drawer.Draw(cg)
}
