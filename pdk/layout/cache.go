package layout

import (
	"github.com/negrel/paon/geometry"
)

var _ BoxedObject = &Cache{}
var _ Packer = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	Packer

	isExpired  bool
	cache      PositionedBoxModel
	constraint Constraint
}

// NewCache returns a new Cache wrapper for the given Flowable.
func NewCache(p Packer) *Cache {
	return &Cache{
		Packer:    p,
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

// Pack implements the Packer interface.
func (c *Cache) Pack(co Constraint) BoxModel {
	box := c.Packer.Pack(co)
	c.Store(co, box)

	return box
}

// Store updates cache validity, Constraint and BoxModel.
// Don't forget to call SetPosition to sets the position.
// This function is called by Cache.Pack to store the result
// of the Packer.
func (c *Cache) Store(co Constraint, bm BoxModel) {
	c.constraint = co
	c.isExpired = false
	c.cache.BoxModel = bm
}
