package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

// LayoutContext contains additional information passed for laying out widgets.
type LayoutContext struct {
	ParentSize geometry.Size
	RootSize   geometry.Size
	Extras     any
}

// Renderable define any object that can be rendered.
type Renderable interface {
	draw.Drawer
	layout.Layout[LayoutContext]
}
