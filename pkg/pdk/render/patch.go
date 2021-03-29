package render

import (
	"github.com/negrel/paon/internal/geometry"
)

// Patch
type Patch interface {
	ForEachCell(fn func(pos geometry.Point, cell *Cell))
}
