package widgets

import (
	"github.com/negrel/ginger/v2/render"
	"github.com/negrel/ginger/v2/widgets/node"
)

var _ Widget = &Core{}

// Core is the core element of all widgets.
// Core is intended to be embed in more advanced
// widget.
type Core struct {
	node.Leaf

	cache Cache
	Draw  func(Constraint) *render.Frame
}

// CORE return a new widget core.
func CORE() *Core {
	return &Core{
		Leaf: &node.BaseLeaf{},
	}
}

// Render implements Rendable interface.
func (c *Core) Render(co Constraint) *render.Frame {
	if co == c.cache.C {
		return c.cache.F
	}

	return c.Render(co)
}

var _ Widget = &CoreLayout{}
var _ Layout = &CoreLayout{}

// CoreLayout is the core element of layout
// widgets. CoreLayout is intended to be embed in
// more advanced layout.
type CoreLayout struct {
	*node.BaseBranch

	Children []Widget
	cache    Cache
	Draw     func(Constraint) *render.Frame
}

// Render implements Rendable interface.
func (cl *CoreLayout) Render(co Constraint) *render.Frame {
	if co == cl.cache.C {
		return cl.cache.F
	}

	return cl.Render(co)
}
