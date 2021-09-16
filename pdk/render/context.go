package render

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
)

// Context define the rendering context.
type Context struct {
	layer      *Layer
	origin     geometry.Vec2D
	Constraint layout.Constraint
}

// NewContext returns a new rendering context.
func NewContext(layer *Layer, constraint layout.Constraint) Context {
	return newContext(layer, constraint)
}

func newContext(layer *Layer, constraint layout.Constraint) Context {
	return Context{
		layer:      layer,
		Constraint: constraint,
	}
}

// Origin returns the origin coordinate that will be use for the drawing surface.
func (c Context) Origin() geometry.Vec2D {
	return c.origin
}

// Surface returns the current context drawing surface.
func (c Context) Surface(box layout.BoxModel) draw.Surface {
	return draw.NewSubSurface(c.layer, box.BorderBox().MoveTo(c.origin))
}

// Layer returns the layer embedded in this context.
func (c Context) Layer() *Layer {
	return c.layer
}
