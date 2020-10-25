package render

import (
	"github.com/negrel/paon/internal/utils"
)

// Patch define a rectangle screen area to update.
type Patch struct {
	Origin utils.Point
	Frame  [][]Cell
}
