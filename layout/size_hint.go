package layout

import "github.com/negrel/paon/geometry"

// SizeHint define a size hint returned after a layout call.
type SizeHint struct {
	MinSize          geometry.Size
	Size             geometry.Size
	VerticalPolicy   SizePolicy
	HorizontalPolicy SizePolicy
}

func (sh SizeHint) ExpandingDirections() Orientation {
	var orientation Orientation = 0

	if (sh.HorizontalPolicy & ExpandFlag) != 0 {
		orientation |= HorizontalOrientation
	}
	if (sh.VerticalPolicy & ExpandFlag) != 0 {
		orientation |= VerticalOrientation
	}

	return orientation
}
