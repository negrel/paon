package widgets

import (
	"github.com/negrel/ginger/v2/rendering"
	"github.com/negrel/ginger/v2/widgets/node"
)

var _ Widget = &Core{}

// Core is the core element of all widgets.
// Core is intended to be embed in more advanced
// widget.
type Core struct {
	node.Leaf

	cache  Cache
	Render func(Constraint) *rendering.Frame
}

// CORE return a new widget core.
func CORE() *Core {
	return &Core{
		Leaf: &node.BaseLeaf{},
	}
}

// Draw implements Drawable interface.
func (c *Core) Draw() *rendering.Frame {
	return c.cache.F
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
	Render   func(Constraint) *rendering.Frame
}
