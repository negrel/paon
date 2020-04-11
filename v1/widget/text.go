package widget

import (
	"image"

	"github.com/negrel/ginger/v1/color"
	"github.com/negrel/ginger/v1/painting"
)

var _ Widget = &Text{}

// Text is a basic widget to display text.
type Text struct {
	*Core

	Content    string
	Foreground int32
	Background int32
}

// Draw implements Widget interface.
func (t *Text) Draw(bounds image.Rectangle) *painting.Frame {
	r := []rune(t.Content)
	width := len(r)

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

	// Not enough space
	if width := frame.Patch.Width(); width < len(r) {
		frame.Patch.M[0][width-1] = &painting.CellOverflow
	}

	return frame
}
