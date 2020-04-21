package widgets

import (
	"image"

	"github.com/negrel/ginger/v2/render"
)

var _ Widget = &Core{}

// Core is the core element of all widgets.
// Core is intended to be embedded in more advanced
// widget.
//
// Core focus on the widget tree structure
// and optimization.
type Core struct {
	name   string
	parent Layout
	owner  Layout
	cache  *Cache

	Rendering render.Func
}

// NewCore return a new core layout.
func NewCore(name string) *Core {
	return &Core{
		name: name,
		cache: &Cache{
			valid: false,
			B:     image.Rect(0, 0, 0, 0),
		},
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Name implements Widget interface.
func (c *Core) Name() string {
	return c.name
}

// Attached implements Widget interface.
func (c *Core) Attached() bool {
	return c.owner != nil
}

// Attach implements Widget interface.
func (c *Core) Attach(owner Layout) {
	c.owner = owner
}

// Detach implements Widget interface.
func (c *Core) Detach() {
	c.owner = nil
}

// Owner implements Widget interface.
func (c *Core) Owner() Layout {
	return c.owner
}

// Parent implements Widget interface.
func (c *Core) Parent() Layout {
	return c.parent
}

func (c *Core) setParent(parent Layout) {
	c.parent = parent
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Render implements Rendable interface.
func (c *Core) Render(bounds image.Rectangle) *render.Frame {
	if c.Attached() {
		// Pull cached frame.
		if cachedFrame, valid := c.cache.Pull(bounds); valid {
			return cachedFrame
		}

		// Set valid state before rendering so the rendering function
		// can set it to unvalid if the frame is cropped to fit
		// the bounds.
		c.cache.valid = true
		frame := c.Rendering(bounds)

		// Updating the cache.
		c.cache.B = bounds
		c.cache.F = frame

		return frame
	}

	return nil
}
