package widget

import (
	"image"

	"github.com/negrel/ginger/v1/style"
)

// Text is a basic widget to display text.
type Text struct {
	*Base

	f       *style.Frame
	Content string
	Colors  style.Colors
}

/*****************************************************
 ********************* INTERFACE *********************
 *****************************************************/
// ANCHOR Getters & setter

// Widget

// Draw implements Widget interface.
func (t *Text) Draw(c Constraint) *style.Frame {
	r := []rune(t.Content)
	width := len(r)

	// Respect width constraint
	if width > c.R.Dx() {
		width = c.R.Dx()
	}

	row := make(style.Row, len(r))

	for i := 0; i < width; i++ {
		row[i] = &style.Cell{
			Char:   r[i],
			Colors: t.Colors,
		}
	}

	t.f = &style.Frame{
		R: image.Rectangle{
			Min: c.R.Min,
			Max: c.R.Min.Add(image.Pt(width, 1)),
		},
		G: style.Grid{
			row,
		},
	}

	return t.f
}
