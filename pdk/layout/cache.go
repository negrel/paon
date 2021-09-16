package layout

var _ BoxedObject = &Cache{}
var _ Object = &Cache{}

// Cache is a wrapper for Flowable object.
type Cache struct {
	Object
	isExpired  bool
	cache      BoxModel
	constraint Constraint
}

// NewCache returns a new Cache wrapper for the given Flowable.
func NewCache(obj Object) *Cache {
	return &Cache{
		Object: obj,
	}
}

// Layout implements the Object interface.
func (c *Cache) Layout(constraint Constraint) BoxModel {
	c.constraint = constraint

	// the cache is still valid if the new constraint has the same size
	// than the cached constraint and the distance between the Min and Max
	// rectangle remains the same.
	if c.IsValid(constraint) {
		return c.cache
	} else {
		box := c.Object.Layout(constraint)
		c.cache = box
		c.isExpired = true

		return box
	}
}

// IsValid returns true if the cache data is valid.
func (c *Cache) IsValid(co Constraint) bool {
	mbSize := c.Box().MarginBox().Size()
	return c.isExpired && co.ApplyOnSize(mbSize).Equals(mbSize)
}

// IsExpired returns true if the cache is marked as expired.
func (c *Cache) IsExpired() bool {
	return c.isExpired
}

// Expire marks the cache as expired.
func (c *Cache) Expire() {
	c.isExpired = true
}

// Constraint returns the cached constraint of the last layout.
func (c *Cache) Constraint() Constraint {
	return c.constraint
}

// Box returns the cached BoxModel of the last flow.
func (c *Cache) Box() BoxModel {
	return c.cache
}
