package widgets

import "github.com/negrel/paon/internal/geometry"

type cache struct {
	position      geometry.Point
	validPosition bool

	width      int
	validWidth bool

	height      int
	validHeight bool
}
