package helpers

// Rectangle define a rectangle with two points
// for the minimum (top-left corner) and the
// maximum (bottom-right corner).
type Rectangle struct {
	Min, Max Point
}

// Rect return always a well-formed rectangle with the
// given dimension.
func Rect(x0, y0, x1, y1 int) Rectangle {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle{
		Min: Pt(x0, y0),
		Max: Pt(x1, y1),
	}
}

// RectFromCenter return a rectangle with the given
// size and the given center.
func RectFromCenter(center Point, size Size) Rectangle {
	return Rectangle{
		Min: Pt(
			center.X-size.Width()/2,
			center.Y-size.Height()/2,
		),
		Max: Pt(
			center.X+size.Width()/2,
			center.Y+size.Height()/2,
		),
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Bottom return offset of the bottom edge from the
// y axis.
func (r Rectangle) Bottom() int {
	return r.Max.Y
}

// BottomCenter return the offset to the center of the
// bottom center of this rectangle.
func (r Rectangle) BottomCenter() Point {
	return Pt(r.Width()/2, r.Max.Y)
}

// BottomLeft return the offset to the bottom left
// corner of the bottom center of this rectangle.
func (r Rectangle) BottomLeft() Point {
	return Pt(r.Max.Y, r.Max.X)
}

// BottomRight return the offset to the bottom right
// corner of the bottom center of this rectangle.
func (r Rectangle) BottomRight() Point {
	return Pt(r.Right(), r.Bottom())
}

// Center return the center of the rectangle.
func (r Rectangle) Center() Point {
	return Pt(r.Min.X+r.Width()/2, r.Min.Y+r.Height()/2)
}

// CenterLeft The offset to the center of the left edge
// of this rectangle.
func (r Rectangle) CenterLeft() Point {
	return Pt(r.Min.X, r.Min.Y+r.Height()/2)
}

// CenterRight The offset to the center of the roght edge
// of this rectangle.
func (r Rectangle) CenterRight() Point {
	return Pt(r.Max.X, r.Min.Y+r.Height()/2)
}

// Height return the height of the rectangle
func (r Rectangle) Height() int {
	return r.Max.Y - r.Min.Y
}

// Left return the offset of the left edge of this
// rectangle from the x axis.
func (r Rectangle) Left() int {
	return r.Min.X
}

// Size return the rectangle width and height in a
// Size object.
func (r Rectangle) Size() Size {
	return Size{
		X: r.Width(),
		Y: r.Height(),
	}
}

// Right return the offset of the right edge of
// this rectangle from the x axis.
func (r Rectangle) Right() int {
	return r.Max.X
}

// Top return offset of the top edge from the
// y axis.
func (r Rectangle) Top() int {
	return r.Min.Y
}

// TopCenter return the offset to the center of the
// top center of this rectangle.
func (r Rectangle) TopCenter() Point {
	return Pt(r.Width()/2, r.Min.Y)
}

// TopLeft return the offset to the bottom left
// corner of the top center of this rectangle.
func (r Rectangle) TopLeft() Point {
	return Pt(r.Left(), r.Top())
}

// TopRight return the offset to the bottom right
// corner of the bottom center of this rectangle.
func (r Rectangle) TopRight() Point {
	return r.Max
}

// Width return the widht of the rectangle
func (r Rectangle) Width() int {
	return r.Max.X - r.Min.X
}
