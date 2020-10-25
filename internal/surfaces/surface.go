package surfaces

import (
	"github.com/negrel/paon/internal/render"
)

// Surface define terminal window/surface to draw on.
type Surface interface {
	Update()
	Apply(render.Patch)
}
