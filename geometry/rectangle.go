package geometry

var _ Sized = &Rectangle{}

// Rectangle define a rectangle with two 2D vectors
// for the minimum (top-left corner) and the
// maximum (bottom-right corner).
type Rectangle struct {
	// Top Left corner of rectangle.
	Origin Vec2D

	// RectSize define size of rectangle.
	RectSize Size
}

// Size implements Sized.
func (r Rectangle) Size() Size {
	return r.RectSize
}

// Top returns origin position on Y axis.
func (r Rectangle) Top() int {
	return r.Origin.Y
}

// Left returns origin position on X axis.
func (r Rectangle) Left() int {
	return r.Origin.X
}

// Right returns position of the right edge on X axis.
func (r Rectangle) Right() int {
	return r.Origin.X + r.RectSize.Width
}

// Bottom returns position of the bottom edge on Y axis.
func (r Rectangle) Bottom() int {
	return r.Origin.Y + r.RectSize.Height
}

// TopLeft returns position of top left corner.
func (r Rectangle) TopLeft() Vec2D {
	return r.Origin
}

// TopRight returns position of top right corner.
func (r Rectangle) TopRight() Vec2D {
	return Vec2D{r.Right(), r.Top()}
}

// BottomRight returns position of bottom right corner.
func (r Rectangle) BottomRight() Vec2D {
	return Vec2D{r.Right(), r.Bottom()}
}

// BottomLeft returns position of bottom left corner.
func (r Rectangle) BottomLeft() Vec2D {
	return Vec2D{r.Left(), r.Bottom()}
}

// Area computes surface area of the rectangle.
func (r Rectangle) Area() int {
	return r.RectSize.Width * r.RectSize.Height
}

// Contains returns true if it contains the given vector.
func (r Rectangle) Contains(p Vec2D) bool {
	return p.X >= r.TopLeft().X && p.X < r.BottomRight().X &&
		p.Y >= r.TopLeft().Y && p.Y < r.BottomRight().Y
}

// GrowLeft returns a new rectangle with left edge moved by n unit to the left.
func (r Rectangle) GrowLeft(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin.Sub(Vec2D{n, 0}),
		RectSize: Size{
			Width:  r.RectSize.Width + n,
			Height: r.RectSize.Height,
		},
	}
}

// ShrinkLeft returns a new rectangle with left edge moved by n unit to the right.
func (r Rectangle) ShrinkLeft(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin.Add(Vec2D{n, 0}),
		RectSize: Size{
			Width:  r.RectSize.Width - n,
			Height: r.RectSize.Height,
		},
	}
}

// GrowRight returns a new rectangle with right edge moved by n unit to the right.
func (r Rectangle) GrowRight(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin,
		RectSize: Size{
			Width:  r.RectSize.Width + n,
			Height: r.RectSize.Height,
		},
	}
}

// ShrinkRight returns a new rectangle with right edge moved by n unit to the left.
func (r Rectangle) ShrinkRight(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin,
		RectSize: Size{
			Width:  r.RectSize.Width - n,
			Height: r.RectSize.Height,
		},
	}
}

// GrowTop returns a new rectangle with top edge moved by n unit to the top.
func (r Rectangle) GrowTop(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin.Sub(Vec2D{0, n}),
		RectSize: Size{
			Width:  r.RectSize.Width,
			Height: r.RectSize.Height + n,
		},
	}
}

// ShrinkTop returns a new rectangle with top edge moved by n unit to the bottom.
func (r Rectangle) ShrinkTop(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin.Add(Vec2D{0, n}),
		RectSize: Size{
			Width:  r.RectSize.Width,
			Height: r.RectSize.Height - n,
		},
	}
}

// GrowBottom returns a new rectangle with bottom edge moved by n unit to the bottom.
func (r Rectangle) GrowBottom(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin,
		RectSize: Size{
			Width:  r.RectSize.Width,
			Height: r.RectSize.Height + n,
		},
	}
}

// ShrinkBottom returns a new rectangle with bottom edge moved by n unit to the top.
func (r Rectangle) ShrinkBottom(n int) Rectangle {
	return Rectangle{
		Origin: r.Origin,
		RectSize: Size{
			Width:  r.RectSize.Width,
			Height: r.RectSize.Height - n,
		},
	}
}

// MoveBy returns a new Rectangle moved by the given offset.
func (r Rectangle) MoveBy(n Vec2D) Rectangle {
	return Rectangle{
		Origin:   r.Origin.Add(n),
		RectSize: r.RectSize,
	}
}
