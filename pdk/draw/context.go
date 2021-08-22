package draw

import (
	"github.com/negrel/paon/internal/geometry"
)

// Context define a drawing context on a Canvas.
type Context struct {
	canvas Canvas
}

// NewContext returns a new Context with the given canvas and bounds.
func NewContext(canvas Canvas) Context {
	return newContext(canvas)
}

func newContext(canvas Canvas) Context {
	return Context{
		canvas: canvas,
	}
}

// Canvas returns the canvas that is tied to this Context.
func (c Context) Canvas() Canvas {
	return c.canvas
}

// FillRectangle draws a geometry.Rectangle that is filled according to the current fill color.
// Note that this methods overwrite text present on the Canvas.
func (c Context) FillRectangle(rectangle geometry.Rectangle, cs CellStyle) {
	for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
		for j := rectangle.Min.Y(); j < rectangle.Max.Y(); j++ {
			pos := geometry.Pt(i, j)

			c.canvas.Set(pos, Cell{
				Content: '\000',
				Style:   cs,
			})
		}
	}
}

// FillTextH draws the given text horizontally from the given origin on this Canvas.
// If the text overflow the Canvas of the context, the overflowing chars are dropped.
func (c Context) FillTextH(origin geometry.Point, cs CellStyle, text string) {
	c.FillRunesH(origin, cs, []rune(text)...)
}

// FillRunesH is equivalent to FillTextH but receive a list of rune instead of a string.
func (c Context) FillRunesH(origin geometry.Point, cs CellStyle, runes ...rune) {
	for i, char := range runes {
		pos := geometry.Pt(origin.X()+i, origin.Y())

		c.canvas.Set(pos, Cell{
			Content: char,
			Style:   cs,
		})
	}
}

// FillTextV draws the given text vertically from the given origin on this Canvas.
// If the text overflow this canvas, the overflowing chars are dropped.
func (c Context) FillTextV(origin geometry.Point, cs CellStyle, text string) {
	c.FillRunesV(origin, cs, []rune(text)...)
}

// FillRunesV is equivalent to FillTextV but receive a list of rune instead of a string.
func (c *Context) FillRunesV(origin geometry.Point, cs CellStyle, runes ...rune) {
	for i, char := range runes {
		pos := geometry.Pt(origin.X(), origin.Y()+i)

		c.canvas.Set(pos, Cell{
			Content: char,
			Style:   cs,
		})
	}
}

// FillLine draws a line between the two given geometry.Point.
func (c *Context) FillLine(from, to geometry.Point) {
	panic("implement me")
}
