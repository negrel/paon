package widget

import (
	"image"

	"github.com/negrel/ginger/v1/painting"
)

var _ Widget = &Core{}

// Core is the core element of all widgets.
type Core struct {
	active bool
	parent Layout
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget

// Abandonned implements Widget interface.
func (c *Core) Abandonned() {
	c.parent = nil
	c.active = false
}

// AdoptedBy implements Widget interface.
func (c *Core) AdoptedBy(l Layout) {
	c.parent = l
	c.active = true
}

// Draw implements Widget interface.
func (c *Core) Draw(bounds image.Rectangle) *painting.Frame {
	return nil
}

// Parent implements Widget interface.
func (c *Core) Parent() Layout {
	return c.parent
}
