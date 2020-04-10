package layout

import (
	"image"

	"github.com/negrel/ginger/v1/painting"
	"github.com/negrel/ginger/v1/widget"
)

var _ Layout = &Container{}
var _ widget.Widget = &Container{}

// Container is a convenient widget that combines common
// painting, positioning, and sizing widgets.
type Container struct {
	*BaseSingleChild

	// Background define the background color of the container
	Background int32

	// Margin is the space that surround the container and
	// the child
	margin int

	// Padding is the space that surround the child
	padding int
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

// Draw implements Widget interface.
func (c *Container) Draw(bounds image.Rectangle) *painting.Frame {

	return nil
}
