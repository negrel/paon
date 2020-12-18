package draw

import (
	"github.com/negrel/paon/internal/geometry"
)

// Patch define a rectangle area to draw on.
type Patch struct {
	geometry.Rectangle
	grid map[int]map[int]*Cell
}

func MakePatch(bound geometry.Rectangle) Patch {
	p := Patch{
		Rectangle: bound,
		grid:      make(map[int]map[int]*Cell, bound.Height()),
	}

	return p
}

func (c Patch) Origin() geometry.Point {
	return c.Min
}

func (c Patch) Bounds() geometry.Rectangle {
	return c.Rectangle
}

func (c Patch) Get(pt geometry.Point) *Cell {
	if c.Contain(pt) {
		// Make position relative to the top-left corner
		// of the patch.
		pt = pt.Add(c.Rectangle.Min)

		if row, ok := c.grid[pt.Y()]; ok {
			if cell, ok := row[pt.X()]; ok {
				return cell
			} else { // Missing cell
				c := &Cell{}
				row[pt.X()] = c

				return c
			}
		} else { // Missing row
			c.grid[pt.Y()] = make(map[int]*Cell)
			return c.Get(pt)
		}
	}

	return nil
}

func (c Patch) ForEach(fn func(cell *Cell)) {
	for y := c.Rectangle.Min.Y(); y <= c.Rectangle.Max.Y(); y++ {
		for x := c.Rectangle.Min.X(); x <= c.Rectangle.Max.X(); x++ {
			fn(c.Get(geometry.Pt(x, y)))
		}
	}
}
