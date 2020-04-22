package widgets

import (
	"image"

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
func Text(content string, theme style.Theme) Widget {
	t := &_text{
		Core:    NewCore("text"),
		content: []rune(content),
		Theme:   theme,
	}

	t.Rendering = t.rendering

	return t
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Rendering implements Widget interface.
func (t *_text) rendering(bounds image.Rectangle) *render.Frame {
	len := len(t.content)
	width := len

	// If text overflow bounds
	if maxWidth := bounds.Dx(); maxWidth < width {
		width = maxWidth
		t.cache.Invalid()
	}

	// Frame to return
	frame := render.NewFrame(width, 1)

	for i := 0; i < width; i++ {
		frame.Patch.M[0][i] = &render.Cell{
			Char:  t.content[i],
			Theme: t.Theme,
		}
	}

	return frame
}
