package render

import (
	"github.com/negrel/paon/internal/geometry"
)

// Patch define an area of the Screen to update.
type Patch struct {
	Pos   geometry.Point
	Cells [][]*Cell
}
