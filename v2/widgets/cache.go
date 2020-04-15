package widgets

import "github.com/negrel/ginger/v2/rendering"

// Cache is used to cache the rendered frame of
// the widget and improve performance. The cache
// is returned when the components is drawn.
type Cache struct {
	C Constraint
	F *rendering.Frame
}
