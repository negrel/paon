package render

import (
	"github.com/negrel/paon/internal/geometry"
)

// Patch define an area of the Screen to update.
type Patch struct {
	Pos   geometry.Point
	Cells [][]*Cell
}

func (p Patch) ForEachCell(fn func(pos geometry.Point, cell *Cell)) {
	height := len(p.Cells)
	width := len(p.Cells[0])

	for i := 0; i < height; i++ {
		y := p.Pos.Y() + i
		for j := 0; j < width; j++ {
			x := p.Pos.X() + j
			fn(geometry.Pt(x, y), p.Cells[i][j])
		}
	}
}
