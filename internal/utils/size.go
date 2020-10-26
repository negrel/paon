package utils

// Size define the width and the height of an
// object.
type Size Point

// Height of the object.
func (s *Size) Height() int {
	return s.Y
}

// Width of the object.
func (s *Size) Width() int {
	return s.X
}

// Equal return whether the given size is equal.
func (s Size) Equal(other Size) bool {
	return s.X == other.X && s.Y == other.Y
}
