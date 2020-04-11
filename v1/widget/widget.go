package widget

import (
	"image"

	"github.com/negrel/ginger/v1/color"
	"github.com/negrel/ginger/v1/painting"
)

// Constraint give the grid/matrix on which the
// children must draw.
type Constraint struct {
	bounds image.Rectangle

	// Default colors for unused space
	C color.Style
}

// Widget are node in the render tree that have a style
// and some state.
type Widget interface {
	// Draw is use to paint the widget in the given frame
	// and return a rectangle defining the frame it used.
	Draw(image.Rectangle) *painting.Frame

	// AdoptedBy set the parent widget.
	AdoptedBy(Layout)

	// Abandonned method remove the widget parent and make it
	// inactive.
	Abandonned()

	// Parent return the widget parent.
	Parent() Layout
}
