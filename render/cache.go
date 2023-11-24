package render

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

var _ Renderable = &Cache[*VoidRenderable]{}

// Cache define a Renderable wrapper that adds caching to it.
// At the moment only layout.Layout is cached as terminal screen is cleared
// on every render.
type Cache[T Renderable] struct {
	renderable T

	cachedSize       geometry.Size
	cachedConstraint layout.Constraint
}

// NewCache returns a new Cache that wraps the given Renderable.
func NewCache[T Renderable](renderable T) Cache[T] {
	return Cache[T]{
		renderable:       renderable,
		cachedSize:       geometry.Size{},
		cachedConstraint: layout.Constraint{},
	}
}

// Unwrap returns the wrapped renderable.
func (c *Cache[T]) Unwrap() T {
	return c.renderable
}

// MarkDirty implements Renderable.
func (c *Cache[T]) MarkDirty() {
	c.renderable.MarkDirty()
}

// IsDirty implements Renderable.
func (c *Cache[T]) IsDirty() bool {
	return c.renderable.IsDirty()
}

// isReflowNeeded returns whether cache is invalid.
// Cache is considered invalid if renderable is dirty or constraint differ
// from cached one.
func (c *Cache[T]) isReflowNeeded(co layout.Constraint) bool {
	return c.IsDirty() || co != c.cachedConstraint
}

// Layout implements layout.Layout.
// If cache is dirty or constraint differ from constraint of previous layout,
// Layout call is forwarded to wrapped renderable.
func (c *Cache[T]) Layout(co layout.Constraint) geometry.Size {
	if !c.isReflowNeeded(co) {
		c.cachedConstraint = co
		return c.cachedSize
	}

	c.cachedConstraint = co
	size := c.renderable.Layout(co)
	c.cachedSize = size
	return size
}

// Draw implements draw.Drawer.
func (c *Cache[T]) Draw(surface draw.Surface) {
	c.renderable.Draw(surface)
}
