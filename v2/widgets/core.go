package widgets

import (
	"image"

	"github.com/negrel/ginger/v2/render"
)

var _ Widget = &Core{}

// Core is the core element of all widgets.
// Core is intended to be embed in more advanced
// widget.
type Core struct {
	name   string
	parent Layout
	owner  Layout
	cache  *Cache

	Rendering func(Constraint) *render.Frame
}

// NewCore return a new core layout.
func NewCore(name string) *Core {
	return &Core{
		name: name,
		cache: &Cache{
			valid: false,
			C: Constraint{
				Bounds: image.Rect(0, 0, 0, 0),
			},
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
func (c *Core) Render(co Constraint) *render.Frame {
	if c.Attached() {
		// Pull cached frame.
		cachedFrame := c.cache.Pull(co)

		if cachedFrame != nil {
			return cachedFrame
		}

		frame := c.Rendering(co)

		c.cache.F = frame
		c.cache.C = co
		c.cache.valid = true

		return frame
	}

	return nil
}
