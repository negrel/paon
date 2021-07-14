package layout

import "github.com/negrel/debuggo/pkg/assert"

var _ Boxed = &Cache{}
var _ Manager = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	Manager
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

	if c.cache != nil && c.constraint.Equals(constraint) {
		return c.cache
	}

	// Update cache
	c.constraint = constraint
	c.cache = c.Manager.Layout(constraint)

	return c.cache
}

// Invalidate invalidates the cache data.
func (c *Cache) Invalidate() {
	c.cache = nil
}

// IsValid returns true if the cache data is valid.
func (c *Cache) IsValid() bool {
	return c.cache != nil
}

// Constraint returns the cached constraint of the last layout.
func (c *Cache) Constraint() Constraint {
	return c.constraint
}

// Box return the cached BoxModel of the last flow.
func (c *Cache) Box() BoxModel {
	return c.cache
}
