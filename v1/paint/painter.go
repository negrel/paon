package paint

import (
	"log"
	"time"

	"github.com/negrel/ginger/v1/style"
)

// Painter paint layer to the screen.
type Painter struct {
	Channel <-chan style.Frame
	Paint   func(x, y int, mainc rune, combc []rune, style style.Colors)
	Refresh func()
}

// Start the painter.
func (p *Painter) Start() {
	for {
		select {
		case frame := <-p.Channel:
			log.Printf("Painter receive : %+v ", frame.R.Size())
			log.Printf("Painter receive : %+v ", frame.G)
			for i := 0; i < frame.R.Dy(); i++ {
				for j := 0; j < frame.R.Dx(); j++ {
					rect := frame.R
					cell := RawCell{
						X:     rect.Min.X + j,
						Y:     rect.Min.Y + i,
						Mainc: frame.G[i][j].Char,
						Style: frame.G[i][j].Colors,
					}
					p.Paint(cell.X, cell.Y, cell.Mainc, []rune{}, cell.Style)
				}
			}

		// Every 16 ms update the screen.
		// 60 fps.
		case <-time.After(time.Millisecond * 16):
			p.Refresh()
		}
	}
}
