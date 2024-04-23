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

func (s Size) WithHeight(h int) Size {
	s.Height = h
	return s
}

func (s Size) WithWidth(w int) Size {
	s.Width = w
	return s
}
