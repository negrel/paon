package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
)

type CellGrid map[int]map[int]*render.Cell

type Canvas struct {
	Bounds   geometry.Rectangle
	cellGrid CellGrid
}

// MakeCanvas returns a new Canvas object.
func MakeCanvas(bounds geometry.Rectangle) Canvas {
	return Canvas{
		Bounds:   bounds,
		cellGrid: make(CellGrid),
	}
}

// SubCanvas returns a Canvas sharing the same render.Cell grid
// but with different bounds.
func (c Canvas) SubCanvas(bounds geometry.Rectangle) Canvas {
	return Canvas{
		Bounds:   bounds,
		cellGrid: c.cellGrid,
	}
}

// Clean cleans the render.Cell of the canvas.
func (c Canvas) Clean() {
	c.cellGrid = make(CellGrid)
}

// Get returns a pointer to the render.Cell at the given position.
// Nil value is only returned if the given position is not within the Canvas bounds.
func (c Canvas) Get(point geometry.Point) *render.Cell {
	if !c.Bounds.Contains(point) {
		return nil
	}

	if row, ok := c.cellGrid[point.Y()]; ok {
		if cell, ok := row[point.X()]; ok {
			return cell
		} else {
			c := &render.Cell{}
			row[point.X()] = c

			return c
		}
	} else {
		c.cellGrid[point.Y()] = make(map[int]*render.Cell)
		return c.Get(point)
	}
}
