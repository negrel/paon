package layout

// Orientation is used to signify an object orientation. Note that,
// vertical and horizontal orientation aren't mutually exclusive.
type Orientation uint8

const (
	NoOrientation Orientation = iota
	HorizontalOrientation
	VerticalOrientation
)

// Horizontal returns whether horizontal bit is set to true.
func (o Orientation) Horizontal() bool {
	return o&HorizontalOrientation != 0
}

// Vertical returns whether vertical bit is set to true.
func (o Orientation) Vertical() bool {
	return o&VerticalOrientation != 0
}
