package render

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

// Renderable define any object that can be rendered.
type Renderable interface {
	layout.Layout
	draw.Drawer

	// IsDirty return whether renderable needs to be rendered.
	IsDirty() bool

	// MarkDirty marks the renderable as dirty. It a tree of renderable,
	// implementation may also mark all ancestor of a node as dirty.
	MarkDirty()
}

// RenderableAccessor define any object owning a Renderable.
type RenderableAccessor interface {
	Renderable() Renderable
}

// VoidRenderable define a minimal renderable with a size of zero. It handles
// efficiently dirty state management and is ideal for embedding into more complex
// renderable.
type VoidRenderable struct {
	target events.Target
	dirty  bool
}

// NewVoidRenderable returns a new VoidRenderable.
func NewVoidRenderable(target events.Target) VoidRenderable {
	return VoidRenderable{
		dirty: true,
	}
}

// Renderable implements RenderableAccessor.
func (vr *VoidRenderable) Renderable() Renderable {
	return vr
}

// Draw implements Renderable.
// VoidRenderable doesn't draw anything but change it's internal state to clean
// (not dirty).
func (vr *VoidRenderable) Draw(draw.Surface) {
	vr.dirty = false
}

// Layout implements Renderable.
// VoidRenderable always return a size of 0.
func (*VoidRenderable) Layout(layout.Constraint) geometry.Size {
	return geometry.Size{}
}

// IsDirty implements Renderable.
func (vr *VoidRenderable) IsDirty() bool {
	return vr.dirty
}

// MarkDirty implements Renderable.
func (vr *VoidRenderable) MarkDirty() {
	if !vr.dirty {
		vr.dirty = true
	}
}

type LayoutLayout = layout.Layout

// ComposedRenderable define a Renderable composed of a VoidRenderable,
// a layout.Layout and a draw.Drawer.
type ComposedRenderable struct {
	VoidRenderable
	LayoutLayout
	draw.Drawer
}

func NewComposedRenderable(layout layout.Layout, drawer draw.Drawer) *ComposedRenderable {
	return &ComposedRenderable{
		VoidRenderable: NewVoidRenderable(),
		LayoutLayout:   layout,
		Drawer:         drawer,
	}
}

// Renderable implements render.RenderableAccessor.
func (cr *ComposedRenderable) Renderable() Renderable {
	return cr
}

// Layout implements Renderable.
func (cr *ComposedRenderable) Layout(co layout.Constraint) geometry.Size {
	return cr.LayoutLayout.Layout(co)
}

// Draw implements Renderable.
func (cr *ComposedRenderable) Draw(surface draw.Surface) {
	cr.VoidRenderable.Draw(surface)
	cr.Drawer.Draw(surface)
}
