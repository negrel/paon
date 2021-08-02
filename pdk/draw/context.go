package draw

import (
	stdcontext "context"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles/value"
)

// Context define a drawing context on a Canvas.
type Context struct {
	canvas    Canvas
	bounds    geometry.Rectangle
	fillColor value.Color
	ops       *[]func(Canvas)
	ctx       stdcontext.Context
}

// NewContext returns a new Context with the given canvas and bounds.
func NewContext(ctx stdcontext.Context, canvas Canvas, bounds geometry.Rectangle, opsBuffer *[]func(Canvas)) *Context {
	return newContext(ctx, canvas, bounds, opsBuffer)
}

func newContext(ctx stdcontext.Context, canvas Canvas, bounds geometry.Rectangle, opsBuffer *[]func(Canvas)) *Context {
	return &Context{
		canvas: canvas,
		bounds: bounds,
		ops:    opsBuffer,
		ctx:    ctx,
	}
}

// SubContext returns a new Context sharing the same Canvas, context.Context and
// drawing operation buffer.
func (c *Context) SubContext(bounds geometry.Rectangle) *Context {
	return &Context{
		canvas: c.canvas,
		bounds: c.bounds.Mask(bounds),
		ops:    c.ops,
		ctx:    c.ctx,
	}
}

func (c *Context) addOp(op func(c Canvas)) {
	*c.ops = append(*c.ops, op)
}

// Canvas returns the canvas that is tied to this Context.
func (c *Context) Canvas() Canvas {
	return c.canvas
}

// Bounds returns the bounds of the underlying Canvas of this Context.
func (c *Context) Bounds() geometry.Rectangle {
	return c.bounds
}

// SetFillColor sets the fill color for fill methods.
func (c *Context) SetFillColor(color value.Color) {
	c.fillColor = color
}

// FillColor returns the current fill color.
func (c *Context) FillColor() value.Color {
	return c.fillColor
}

// FillRectangle draws a geometry.Rectangle that is filled according to the current fill color.
// Note that this methods overwrite text present on the Canvas.
func (c *Context) FillRectangle(rectangle geometry.Rectangle) {
	rectangle = c.bounds.Mask(rectangle)
	fillColor := c.fillColor

	c.addOp(func(canvas Canvas) {
		for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
			for j := rectangle.Min.Y(); j < rectangle.Max.Y(); j++ {
				pos := geometry.Pt(i, j)

				cell := Cell{}
				cell.Style.Background = fillColor
				cell.Style.Foreground = fillColor
				cell.Content = 0

				canvas.Set(pos, cell)
			}
		}
	})
}

// FillTextH draws the given text horizontally from the given origin on this Canvas.
// If the text overflow the Canvas of the context, the overflowing chars are dropped.
func (c *Context) FillTextH(origin geometry.Point, text string) {
	c.FillRunesH(origin, []rune(text)...)
}

// FillRunesH is equivalent to FillTextH but receive a list of rune instead of a string.
func (c *Context) FillRunesH(origin geometry.Point, runes ...rune) {
	origin = origin.Add(c.bounds.Min)
	rectangle := geometry.Rect(origin.X(), origin.Y(), origin.X()+len(runes), origin.Y()+1)
	rectangle = c.bounds.Mask(rectangle)
	if rectangle.Empty() {
		return
	}

	fillColor := c.fillColor
	c.addOp(func(canvas Canvas) {
		for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
			pos := geometry.Pt(i, origin.Y())
			cell := canvas.Get(pos)

			cell.Content = runes[i-rectangle.Min.X()]
			cell.Style.Foreground = fillColor

			canvas.Set(pos, cell)
		}
	})
}

// FillTextV draws the given text vertically from the given origin on this Canvas.
// If the text overflow this canvas, the overflowing chars are dropped.
func (c *Context) FillTextV(origin geometry.Point, text string) {
	c.FillRunesV(origin, []rune(text)...)
}

// FillRunesV is equivalent to FillTextV but receive a list of rune instead of a string.
func (c *Context) FillRunesV(origin geometry.Point, runes ...rune) {
	origin = origin.Add(c.bounds.Min)
	rectangle := geometry.Rect(origin.X(), origin.Y(), origin.X()+1, origin.Y()+len(runes))
	rectangle = c.bounds.Mask(rectangle)
	if rectangle.Empty() {
		return
	}

	fillColor := c.fillColor
	c.addOp(func(canvas Canvas) {
		for i := rectangle.Min.Y(); i < rectangle.Max.Y(); i++ {
			pos := geometry.Pt(origin.X(), i)
			cell := canvas.Get(pos)

			cell.Content = runes[i-rectangle.Min.Y()]
			cell.Style.Foreground = fillColor

			canvas.Set(pos, cell)
		}
	})
}

// FillLine draws a line between the two given geometry.Point.
func (c *Context) FillLine(from, to geometry.Point) {
	panic("implement me")
}

// Commit implements the Context interface.
func (c *Context) Commit() {
	for i, op := range *c.ops {
		select {
		case <-c.ctx.Done():
			*c.ops = (*c.ops)[i:]
			return

		default:
			op(c.canvas)
		}
	}

	*c.ops = (*c.ops)[:0]
}
