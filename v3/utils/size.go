package utils

// Size define the width and the height of an
// object.
type Size Point

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Width of the object.
func (s *Size) Width() int {
	return s.X
}

// Height of the object.
func (s *Size) Height() int {
	return s.Y
}

// IsFinite return wether or not
// width and height are finite value.
func (s Size) IsFinite() bool {
	return s.Width() != Infinite && s.Height() != Infinite
}

// IsInfinite return wether or not
// width or height are infinite value.
func (s Size) IsInfinite() bool {
	return s.Width() == Infinite || s.Height() == Infinite
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// BottomCenter return the offset to the center of the
// bottom edge of the rectangle described by the given
// offset (which is interpreted as the top-left corner)
// and this size.
func (s Size) BottomCenter(origin Point) Point {
	return Pt(
		origin.X+s.Width()/2,
		origin.Y+s.Height(),
	)
}

// BottomLeft return the offset to the intersection of
// the bottom and left edges of the rectangle described
// by the given offset which is interpreted as the top-left
// corner) and this size.
func (s Size) BottomLeft(origin Point) Point {
	return Pt(
		origin.X,
		origin.Y+s.Height(),
	)
}

// BottomRight The offset to the intersection of the bottom
// and right edges of the rectangle described by the given
// offset (which is interpreted as the top-left corner) and
// this size.
func (s Size) BottomRight(origin Point) Point {
	return Pt(
		origin.X+s.Width(),
		origin.Y+s.Height(),
	)
}

// Center return the offset to the point halfway between
// the left and right and the top and bottom edges of the
// rectangle described by the given offset (which is
// interpreted as the top-left corner) and this size.
func (s Size) Center(origin Point) Point {
	return Pt(
		(origin.X + s.Width()/2),
		(origin.Y + s.Height()/2),
	)
}

// CenterLeft return the offset to the center of the left
// edge of the rectangle described by the given offset
// (which is interpreted as the top-left corner) and this
// size.
func (s Size) CenterLeft(origin Point) Point {
	return Pt(
		origin.X,
		origin.Y+s.Height()/2,
	)
}

// CenterRight return the offset to the center of the right
// edge of the rectangle described by the given offset (which
// is interpreted as the top-left corner) and this size.
func (s Size) CenterRight(origin Point) Point {
	return Pt(
		origin.X+s.Width(),
		origin.Y+s.Height()/2,
	)
}

// Contains return whether the point specified by the given
// point (which is assumed to be relative to the top left
// of the size) lies between the left and right and the
// top and bottom edges of a rectangle of this size.
func (s Size) Contains(point Point) bool {
	return point.X >= 0.0 && point.X < s.Width() && point.Y >= 0.0 && point.Y < s.Height()
}

// TopCenter the offset to the center of the top edge of
// the rectangle described by the given offset (which is
// interpreted as the top-left corner) and this size.
func (s Size) TopCenter(origin Point) Point {
	return Pt(
		origin.X+s.Width()/2,
		origin.Y,
	)
}

// TopLeft the offset to the center of the top edge of
// the rectangle described by the given offset (which is
// interpreted as the top-left corner) and this size.
func (s Size) TopLeft(origin Point) Point {
	return Pt(
		origin.X,
		origin.Y,
	)
}

// TopRight the offset to the center of the top edge of
// the rectangle described by the given offset (which is
// interpreted as the top-left corner) and this size.
func (s Size) TopRight(origin Point) Point {
	return Pt(
		origin.X+s.Width(),
		origin.Y,
	)
}
