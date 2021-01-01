package draw

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
)

// Canvas define a rectangle area to draw on.
type Canvas struct {
	geometry.Rectangle
	grid map[int]map[int]*Cell
}

func NewCanvas(bound geometry.Rectangle) *Canvas {
	return &Canvas{
		Rectangle: bound,
		grid:      make(map[int]map[int]*Cell, bound.Height()),
	}
}

func (c *Canvas) Clear() {
	c.grid = make(map[int]map[int]*Cell, c.Height())
}

func (c *Canvas) Origin() geometry.Point {
	return c.Min
}

func (c *Canvas) Bounds() geometry.Rectangle {
	return c.Rectangle
}

func (c *Canvas) Get(pt geometry.Point) *Cell {
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

func (c *Canvas) ForEach(fn func(point geometry.Point, cell *Cell)) {
	for y := c.Rectangle.Min.Y(); y <= c.Rectangle.Max.Y(); y++ {
		for x := c.Rectangle.Min.X(); x <= c.Rectangle.Max.X(); x++ {
			pt := geometry.Pt(x, y)
			if c.Contain(pt) {
				fn(pt, c.Get(pt))
			}
		}
	}
}

func (c *Canvas) Resize(size geometry.Size) {
	assert.GreaterOrEqual(size.Width(), 0, "canvas can't have a negative width")
	assert.GreaterOrEqual(size.Height(), 0, "canvas can't have a negative height")

	c.Rectangle.Max = c.Rectangle.Min.Add(geometry.Point(size))
}
