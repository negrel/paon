package layout

import (
	"github.com/negrel/paon/geometry"
)

var _ Layout = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	layout Layout

	isExpired  bool
	cache      geometry.Size
	constraint Constraint
}

// NewCache returns a new Cache wrapper for the given Flowable.
func NewCache(l Layout) *Cache {
	return &Cache{
		layout:    l,
		isExpired: true,
	}
}

// IsValid returns true if the cache data is valid.
func (c Cache) IsValid(co Constraint) bool {
	return !c.isExpired && co == c.constraint
}

// IsExpired returns true if the cache is marked as expired.
func (c Cache) IsExpired() bool {
	return c.isExpired
}

// Expire marks the cache as expired.
func (c *Cache) Expire() {
	c.isExpired = true
}

// Constraint returns the cached constraint of the last layout.
func (c Cache) Constraint() Constraint {
	return c.constraint
}

// BoundingRect returns cached bounding rectangle.
func (c Cache) Size() geometry.Size {
	return c.cache
}

// Layout implements the Layout interface.
func (c *Cache) Layout(co Constraint) geometry.Size {
	if c.IsValid(co) {
		return co.ApplyOnSize(c.cache)
	}

	size := c.layout.Layout(co)
	c.constraint = co
	c.isExpired = false
	c.cache = size

	return co.ApplyOnSize(size)
}
