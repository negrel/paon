package flows

import "github.com/negrel/debuggo/pkg/assert"

var _ Flowable = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	Flowable
	cache      BoxModel
	constraint Constraint
}

// NewCache returns a new Cache wrapper for the given Flowable.
func NewCache(flowable Flowable) *Cache {
	return &Cache{
		Flowable: flowable,
		cache:    nil,
	}
}

// Flow implements the Flowable interface.
func (c *Cache) Flow(constraint Constraint) BoxModel {
	assert.NotNil(c.Flowable)

	// Check in the cache if valid
	if c.cache != nil && c.constraint == constraint {
		return c.cache
	}

	// Update cache
	c.constraint = constraint
	c.cache = c.Flowable.Flow(constraint)

	return c.cache
}

// Invalidate invalidates the cache data.
func (c *Cache) Invalidate() {
	c.cache = nil
}

// Box return the cached BoxModel of the last flow.
func (c *Cache) Box() BoxModel {
	return c.cache
}
