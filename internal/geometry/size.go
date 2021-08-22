package geometry

var _ Sized = Size{}

// Size define the width and the height of an
// object.
type Size struct {
	width  int
	height int
}

// NewSize return a new Size using the given width and height.
func NewSize(width int, height int) Size {
	return Size{
		width:  width,
		height: height,
	}
}

// Size implements the Sized interface.
func (s Size) Size() Size {
	return s
}

// Height of the object.
func (s Size) Height() int {
	return s.height
}

// Width of the object.
func (s Size) Width() int {
	return s.width
}

// Equals returns true if the given Size is equal to this Size.
func (s Size) Equals(other Size) bool {
	return s.width == other.width && s.height == other.height
}
