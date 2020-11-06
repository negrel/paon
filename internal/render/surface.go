package render

import (
	"github.com/negrel/paon/internal/geometry"
)

type Surface interface {
	geometry.Sized
	Draw(pt geometry.Point, cell Cell)
	Mask(geometry.Rectangle)
}
