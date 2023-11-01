package layout

import (
	"github.com/negrel/paon/geometry"
)

var _ BoxedObject = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	layout Layout

	isExpired  bool
	cache      PositionedBoxModel
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
	if c.cache.BoxModel == nil {
		return false
	}

	mbSize := c.BoxModel().MarginBox().Size()
	return !c.isExpired && co.ApplyOnSize(mbSize).Equals(mbSize)
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

// BoxModel implements the BoxedObject interface.
func (c Cache) BoxModel() BoxModel {
	return c.cache
}

// Position implements the geometry.Positioned interface.
func (c Cache) Position() geometry.Vec2D {
	return c.cache.Origin
}

// SetPosition sets the position of the BoxModel.
func (c *Cache) SetPosition(origin geometry.Vec2D) {
	c.cache.Origin = origin
}

// Layout implements the Layout interface.
func (c *Cache) Layout(co Constraint) BoxModel {
	box := c.layout.Layout(co)
	c.constraint = co
	c.isExpired = false
	c.cache.BoxModel = box

	return box
}
