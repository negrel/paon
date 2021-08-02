package layout

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
)

var unsetSize = geometry.NewSize(-1, -1)

var _ Boxed = &Cache{}
var _ Manager = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	Manager
	valid      bool
	cache      BoxModel
	constraint Constraint
}

// NewCache returns a new Cache wrapper for the given Flowable.
func NewCache(man Manager) *Cache {
	return &Cache{
		Manager: man,
		cache:   nil,
	}
}

// Layout implements the Algo interface.
func (c *Cache) Layout(constraint Constraint) BoxModel {
	assert.NotNil(c.Manager)

	// the cache is still valid if the new constraint has the same size
	// than the cached constraint and the distance between the Min and Max
	// rectangle remains the same.
	if c.valid && c.constraint.EqualsSize(constraint) &&
		c.constraint.Min.Min.Sub(c.constraint.Max.Min) == constraint.Min.Min.Sub(constraint.Max.Min) {
		box := c.cache

		if !c.constraint.Min.Equals(constraint.Min) || !c.constraint.Max.Equals(constraint.Max) {
			box = &movedBox{
				BoxModel: box,
				offset:   c.constraint.Min.Min.Sub(constraint.Min.Min),
			}
		}

		return box
	}

	// Update cache
	c.constraint = constraint
	c.cache = c.Manager.Layout(constraint)
	c.valid = true

	return c.cache
}

// IsValid returns true if the cache data is valid.
func (c *Cache) IsValid() bool {
	return c.valid
}

// Invalidate marks the cache as invalid.
func (c *Cache) Invalidate() {
	c.valid = false
}

// Constraint returns the cached constraint of the last layout.
func (c *Cache) Constraint() Constraint {
	return c.constraint
}

// Box return the cached BoxModel of the last flow.
func (c *Cache) Box() BoxModel {
	return c.cache
}
