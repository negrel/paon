package layout

import "github.com/negrel/paon/geometry"

// Cache define a layout cache.
type Cache struct {
	size       geometry.Size
	constraint Constraint
}

// NewCache returns a new layout cache.
func NewCache() Cache {
	return Cache{}
}

// IsValid returns cached size along whether it is valid or not.
func (c *Cache) IsValid(co Constraint) (geometry.Size, bool) {
	return c.size, c.constraint.RootSize == co.RootSize &&
		c.constraint.ParentSize == co.ParentSize &&
		co.ApplyOnSize(c.size) == c.constraint.ApplyOnSize(c.size)
}

// UpdateCache update cached size.
func (c *Cache) UpdateCache(co Constraint, size geometry.Size) {
	c.size = size
	c.constraint = co
}
