package render

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
)

// Buffer defined a rectangle area to draw on.
type Buffer struct {
	Bounds geometry.Rectangle
	grid   map[int]map[int]*Cell
}

func NewCanvas(bound geometry.Rectangle) *Buffer {
	return &Buffer{
		Bounds: bound,
		grid:   make(map[int]map[int]*Cell, bound.Height()),
	}
}

func (b *Buffer) Clear() {
	b.grid = make(map[int]map[int]*Cell, b.Bounds.Height())
}

func (b *Buffer) Get(pt geometry.Point) *Cell {
	if b.Bounds.Contains(pt) {
		// Make position relative to the top-left corner
		// of the patch.
		pt = pt.Add(b.Bounds.Min)

		if row, ok := b.grid[pt.Y()]; ok {
			if cell, ok := row[pt.X()]; ok {
				return cell
			} else { // Missing cell
				c := &Cell{}
				row[pt.X()] = c

				return c
			}
		} else { // Missing row
			b.grid[pt.Y()] = make(map[int]*Cell)
			return b.Get(pt)
		}
	}

	return nil
}

func (b *Buffer) ForEach(fn func(point geometry.Point, cell *Cell)) {
	for y := b.Bounds.Min.Y(); y <= b.Bounds.Max.Y(); y++ {
		for x := b.Bounds.Min.X(); x <= b.Bounds.Max.X(); x++ {
			pt := geometry.Pt(x, y)
			if b.Bounds.Contains(pt) {
				fn(pt, b.Get(pt))
			}
		}
	}
}

func (b *Buffer) Resize(size geometry.Size) {
	assert.GreaterOrEqual(size.Width(), 0, "canvas can't have a negative width")
	assert.GreaterOrEqual(size.Height(), 0, "canvas can't have a negative height")

	b.Bounds.Max = b.Bounds.Min.Add(geometry.Point(size))
}
