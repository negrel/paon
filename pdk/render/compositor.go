package render

import (
	"math"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
)

const maxLayer = math.MaxUint8

type compositeCell struct {
	position geometry.Vec3D
	cell     draw.Cell
}

// Compositor define a rendering compositor
type Compositor struct {
	dst       draw.Surface
	rootLayer Layer

	dirtyMap   []int
	compositeQ chan compositeCell
}

// NewCompositor returns a new Compositor for the given surface.
func NewCompositor(destination draw.Surface) Compositor {
	return Compositor{
		dst:        destination,
		rootLayer:  newLayer(NewBufferSurface(destination.Bounds()), false),
		compositeQ: make(chan compositeCell, 4096),
	}
}
