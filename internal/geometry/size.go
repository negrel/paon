package geometry

var _ Sized = Size{}

// Size define the width and the height of an
// object.
type Size Point

func NewSize(width int, height int) Size {
	return Size{
		x: width,
		y: height,
	}
}

func (s Size) Size() Size {
	return s
}

// Height of the object.
func (s Size) Height() int {
	return s.y
}

// Width of the object.
func (s Size) Width() int {
	return s.x
}

// Equal return whether the given size is equal.
func (s Size) Equal(other Size) bool {
	return s.x == other.x && s.y == other.y
}
