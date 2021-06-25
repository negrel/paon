package layout

import "github.com/negrel/debuggo/pkg/assert"

// Cache is a wrapper for Flowable object.
type Cache struct {
	Algo
	cache      BoxModel
	constraint Constraint
}

// NewCache returns a new Cache wrapper for the given Flowable.
func NewCache(algo Algo) *Cache {
	return &Cache{
		Algo:  algo,
		cache: nil,
	}
}

// Layout implements the Algo interface.
func (c *Cache) Layout(constraint Constraint) BoxModel {
	assert.NotNil(c.Algo)

	if c.cache != nil && c.constraint.Equals(constraint) {
		return c.cache
	}

	// Update cache
	c.constraint = constraint
	c.cache = c.Algo.Layout(constraint)

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
