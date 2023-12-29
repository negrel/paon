package geometry

var _ Sized = Size{}

// Size define the width and the height of an
// object.
type Size struct {
	Width  int
	Height int
}

// Size implements the Sized interface.
func (s Size) Size() Size {
	return s
}
