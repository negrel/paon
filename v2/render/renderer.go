package render

import (
	"image"
	"time"
)

// Renderer handle the render process of each frame received
// by the renderer input channel.
type Renderer struct {
	Input   <-chan *Frame
	Paint   func(*RawCell)
	Refresh func()
}

// Start the painter.
func (p *Renderer) Start() {
	var needRefresh bool = true

	for {
		select {
		case frame := <-p.Input:

			for i := 0; i < frame.Patch.Height(); i++ {
				xOffset := frame.Position.X + i

				for j := 0; j < frame.Patch.Width(); j++ {
					yOffset := frame.Position.Y + j

					pt := image.Pt(yOffset, xOffset)
					cell := frame.Patch.M[i][j].Compute(pt)

					p.Paint(cell)
					needRefresh = true
				}
			}

			// Every 16 ms update the screen.
			// 60 fps.
		case <-time.After(time.Millisecond * 16):
			if needRefresh {
				p.Refresh()
				needRefresh = false
			}
		}
	}
}
