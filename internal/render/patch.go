package render

import (
	"github.com/negrel/paon/internal/geometry"
)

var _ Surface = &Patch{}

// Patch define a rectangle screen area to update.
type Patch struct {
	Origin geometry.Point
	Frame  [][]Cell
}

func NewPatch(bound geometry.Rectangle) *Patch {
	p := &Patch{
		Origin: geometry.Pt(0, 0),
		Frame:  make([][]Cell, bound.Height()),
	}

	w := bound.Width()
	for i := range p.Frame {
		p.Frame[i] = make([]Cell, w)
	}

	return p
}

func (p *Patch) Size() geometry.Size {
	return geometry.NewSize(p.Width(), p.Height())
}

func (p *Patch) Width() int {
	return len(p.Frame[0])
}

func (p *Patch) Height() int {
	return len(p.Frame)
}

func (p *Patch) Draw(pt geometry.Point, cell Cell) {
	p.Frame[pt.Y()][pt.X()] = cell
}
