package widget

import (
	"image"

	"github.com/negrel/ginger/v1/style"
)

// Constraint represent the size & position constrait
// for children widget frames.
type Constraint struct {
	// Space that child can use to draw itself
	R image.Rectangle
	// Default colors for unused space
	c style.Colors
}

// Widget are node in the render tree that have a style
// and some state.
type Widget interface {
	// Draw is use to paint the widget
	Draw(c Constraint) *style.Frame

	// AdoptedBy set the parent widget.
	AdoptedBy(Widget)

	// Parent return the widget parent.
	Parent() Widget

	// Reflow method redraw entirely the component.
	reflow()

	// Repaint method send to the parent an update frame to paint.
	repaint(style.Frame)
}
