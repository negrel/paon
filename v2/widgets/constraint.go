package widgets

import (
	"image"

	"github.com/negrel/ginger/v1/tree/style"
)

// Constraint are used by Drawable to render themself
// without overflowing the constraint.
// Constraint also contain style inheritance for unset
// child style properties.
type Constraint struct {
	// Bounds is the limit of the frame.
	Bounds image.Rectangle

	// Inherited is used for unset cells that inherit from parent.
	Inherited style.Theme
}
