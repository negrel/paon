package widget

import (
	"image"
	"log"

	"github.com/negrel/ginger/v1/color"
	"github.com/negrel/ginger/v1/painting"
)

var _ Widget = &Text{}

// Text is a basic widget to display text.
type Text struct {
	*Base

	Content    string
	Foreground int32
	Background int32
}

// Draw implements Widget interface.
func (t *Text) Draw(bounds image.Rectangle) *painting.Frame {
	r := []rune(t.Content)
	width := len(r)

	log.Println("TEXT BOUNDS : ", bounds)

	// If text overflow bounds
	if cWidth := bounds.Dx(); cWidth < width {
		width = cWidth
	}

	// Frame to return
	frame := painting.NewFrame(bounds.Min, width, 1)

	for i := 0; i < width; i++ {
		frame.Patch.M[0][i] = &painting.Cell{
			Char: r[i],
			Style: color.Style{
				Foreground: t.Foreground,
				Background: t.Background,
			},
		}
	}

	log.Println("TEXT POSITION:", frame.Position)

	return frame
}
