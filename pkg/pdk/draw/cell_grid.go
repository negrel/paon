package draw

import "github.com/negrel/paon/internal/geometry"

var _ Canvas = CellGrid{}

type CellGrid struct {
	bounds geometry.Rectangle
	grid   [][]Cell
}

func (cg CellGrid) Get(pos geometry.Point) Cell {
	if !cg.bounds.Contains(pos) {
		return Cell{}
	}

	return cg.grid[pos.Y()][pos.X()]
}

func (cg CellGrid) Bounds() geometry.Rectangle {
	return cg.bounds
}

func (cg CellGrid) Set(pos geometry.Point, cell Cell) {
	if !cg.bounds.Contains(pos) {
		return
	}

	cg.grid[pos.Y()][pos.X()] = cell
}
