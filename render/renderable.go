package render

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/layout"
)

// Renderable define any object that can be rendered.
type Renderable interface {
	layout.Layout
	draw.Drawer
}
