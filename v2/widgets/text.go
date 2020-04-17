package widgets

import (
	"github.com/negrel/ginger/v2/render"
	"github.com/negrel/ginger/v2/style"
)

var _ Widget = &_text{}

type _text struct {
	*Core

	content []rune
	Theme   style.Theme
}

// Text return a static text widget.
func Text(content string) Widget {
	t := &_text{
		Core:    CORE(),
		content: []rune(content),
	}
	t.Draw = t.draw

	return t
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

func (t *_text) draw(c Constraint) *render.Frame {
	len := len(t.content)
	width := len

	// If text overflow bounds
	if maxWidth := c.Bounds.Dx(); maxWidth < width {
		width = maxWidth
	}

	// Frame to return
	frame := render.NewFrame(c.Bounds.Min, width, 1)

	for i := 0; i < width; i++ {
		frame.Patch.M[0][i] = &render.Cell{
			Char:  t.content[i],
			Theme: &t.Theme,
		}
	}

	// Not enough space / overflow
	if width := frame.Patch.Width(); width < len {
		frame.Patch.M[0][width-1] = &render.CellOverflow
	}

	return frame
}
