package screens

import (
	"github.com/negrel/paon/internal/render"
)

// Screen define terminal window/surface to draw on.
type Screen interface {
	Update()
	Apply(render.Patch)
}
