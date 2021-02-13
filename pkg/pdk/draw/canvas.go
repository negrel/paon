package draw

import "github.com/negrel/paon/internal/geometry"

type Canvas interface {
	Bounds() geometry.Rectangle
	Get(geometry.Point) Cell
	Set(geometry.Point, Cell)
}
