package render

import (
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
)

// Renderable define any object that can be rendered.
type Renderable interface {
	// Render performs the rendering of this object in the given context..
	Render(Context)
}

var _ Renderable = Object{}

// Object is a wrapper around a Renderable object.
type Object struct {
	Renderable

	dirty  bool
	drawer draw.Drawer
	cache  *layout.Cache
}

// Dirty marks the object as dirty.
func (o Object) Dirty() {
	o.dirty = true
}

// IsDirty returns true if the object is marked as dirty?
func (o Object) IsDirty() bool {
	return o.dirty
}

// Render implements the Renderable interface.
func (o Object) Render(ctx Context) {
	if !o.dirty {
		return
	}

	box := o.cache.Layout(ctx.Constraint)
	srf := ctx.Surface(box)

	o.drawer.Draw(srf)
}
