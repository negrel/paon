package render

import (
	"github.com/negrel/paon/pdk/layout"
)

// Flag define option flag for renderint context
type Flag uint8

// All bit flag values
const (
	NoFlag   = 0
	DrawFlag = 1 << iota
	LayoutFlag
	FullRenderFlag = DrawFlag & LayoutFlag
)

// Context define the rendering context.
type Context struct {
	Flags Flag
	Layer
	layout.Layout
	layout.Constraint
}
