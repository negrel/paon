package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

// Context define a drawing context on a Canvas.
type Context interface {
	// Bounds returns the bounds of the underlying Canvas of this Context.
	Bounds() geometry.Rectangle

	// SetFillColor sets the fill color for fill methods.
	SetFillColor(color value.Color)

	// FillColor returns the current fill color.
	FillColor() value.Color

	// FillRectangle draws a geometry.Rectangle that is filled according to the current fill color.
	FillRectangle(rectangle geometry.Rectangle)

	// FillTextH draws the given text horizontally from the given origin on this Canvas.
	// If the text overflow the Canvas of the context, the overflowing chars are dropped.
	FillTextH(origin geometry.Point, text string)

	// FillTextV draws the given text vertically from the given origin on this Canvas.
	// If the text overflow this canvas, the overflowing chars are dropped.
	FillTextV(origin geometry.Point, text string)

	// DrawLine draws a line between the two given geometry.Point.
	FillLine(from, to geometry.Point)

	// Commit applies all the change of the context to the Canvas.
	Commit()
}

var _ Context = &context{}

type context struct {
	canvas    Canvas
	bounds    geometry.Rectangle
	fillColor value.Color
	ops       []func(Canvas)
}

func newContext(canvas Canvas, bounds geometry.Rectangle) *context {
	return &context{
		canvas: canvas,
		ops:    make([]func(Canvas), 0, 8),
	}
}

// Bounds implements the Context interface.
func (c *context) Bounds() geometry.Rectangle {
	return c.bounds
}

// SetFillColor implements the Context interface.
func (c *context) SetFillColor(color value.Color) {
	c.fillColor = color
}

// FillColor implements the Context interface.
func (c *context) FillColor() value.Color {
	return c.fillColor
}

// FillRectangle implements the Context interface.
func (c *context) FillRectangle(rectangle geometry.Rectangle) {
	rectangle = c.bounds.Intersect(rectangle)
	fillColor := c.fillColor

	c.ops = append(c.ops, func(canvas Canvas) {
		for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
			for j := rectangle.Min.Y(); j < rectangle.Max.Y(); j++ {
				canvas.Get(geometry.Pt(i, j)).Style.Background = fillColor
			}
		}
	})
}

// FillTextH implements the Context interface.
func (c *context) FillTextH(origin geometry.Point, text string) {
	origin = origin.Add(c.bounds.Min)
	rectangle := geometry.Rect(origin.X(), origin.Y(), origin.X()+len(text), origin.Y()+1)
	if rectangle.Empty() {
		return
	}

	fillColor := c.fillColor
	c.ops = append(c.ops, func(canvas Canvas) {
		for i := rectangle.Min.X(); i < rectangle.Max.X(); i++ {
			cell := c.canvas.Get(geometry.Pt(i, origin.Y()))
			if cell != nil {
				cell.Content = rune(text[i-origin.X()])
				cell.Style.Foreground = fillColor
			}
		}
	})
}

// FillTextV implements the Context interface.
func (c *context) FillTextV(origin geometry.Point, text string) {
	origin = origin.Add(c.bounds.Min)
	rectangle := geometry.Rect(origin.X(), origin.Y(), origin.X()+1, origin.Y()+len(text))
	if rectangle.Empty() {
		return
	}

	fillColor := c.fillColor
	c.ops = append(c.ops, func(canvas Canvas) {
		for i := rectangle.Min.Y(); i < rectangle.Max.Y(); i++ {
			cell := c.canvas.Get(geometry.Pt(origin.X(), i))
			if cell != nil {
				cell.Content = rune(text[i-origin.Y()])
				cell.Style.Foreground = fillColor
			}
		}
	})
}

// FillLine implements the Context interface.
func (c *context) FillLine(from, to geometry.Point) {
	panic("implement me")
}

// Commit implements the Context interface.
func (c *context) Commit() {
	for _, op := range c.ops {
		op(c.canvas)
	}
	c.ops = make([]func(Canvas), 0, 8)
}
