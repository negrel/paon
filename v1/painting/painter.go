package painting

import (
	"image"
	"time"
)

// Painter paint layer to the screen.
type Painter struct {
	Channel chan *Frame
	Paint   func(RawCell)
	Refresh func()
}

// Start the painter.
func (p *Painter) Start() {
	for {
		select {
		case frame := <-p.Channel:
			xOffset := frame.Position.X
			yOffset := frame.Position.Y

			for i := 0; i < frame.Patch.Height(); i++ {
				for j := 0; j < frame.Patch.Width(); j++ {
					// Computing global position
					pt := image.Pt(yOffset+j, xOffset+i)
					cell := frame.Patch.M[i][j].Compute(pt)

					p.Paint(cell)
				}
			}

		// Every 16 ms update the screen.
		// 60 fps.
		case <-time.After(time.Millisecond * 16):
			p.Refresh()
		}
	}
}
