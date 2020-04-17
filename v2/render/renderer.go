package render

import (
	"image"
	"time"
)

// Renderer render the given frame to RawCell.
type Renderer struct {
	Input   <-chan *Frame
	Paint   func(*RawCell)
	Refresh func()
	isFresh bool
}

// Start the painter.
func (p *Renderer) Start() {
	for {
		select {
		case frame := <-p.Input:
			xOffset := frame.Position.X
			yOffset := frame.Position.Y

			for i := 0; i < frame.Patch.Height(); i++ {
				for j := 0; j < frame.Patch.Width(); j++ {
					// Computing global position
					pt := image.Pt(yOffset+j, xOffset+i)
					cell := frame.Patch.M[i][j].Compute(pt)

					p.Paint(cell)
					p.isFresh = false
				}
			}

		// Every 16 ms update the screen.
		// 60 fps.
		case <-time.After(time.Millisecond * 16):
			if !p.isFresh {
				p.Refresh()
				p.isFresh = false
			}
		}
	}
}
