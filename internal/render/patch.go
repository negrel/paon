package render

import (
	"image"

	"github.com/negrel/paon/internal/render/cells"
)

// Patch define a rectangle screen area to update.
type Patch struct {
	Origin image.Point
	Frame  [][]cells.Cell
}
