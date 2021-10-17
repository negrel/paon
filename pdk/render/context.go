package render

import (
	"github.com/negrel/paon/pdk/layout"
)

// Flag define option flag for renderint context
type Flag uint8

const (
	NoFlag = 1 << iota
	DrawFlag
	LayoutFlag
	FullRenderFlag
)

// Context define the rendering context.
type Context struct {
	Flags Flag
	Layer
	layout.Layout
	layout.Constraint
}
