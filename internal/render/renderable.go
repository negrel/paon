package render

import (
	"github.com/negrel/paon/internal/utils"
)

// Renderable define any object that can produce a render Patch.
type Renderable interface {
	// Render the object inside the given rectangle and return a Patch object.
	Render(rectangle utils.Rectangle) Patch
}
