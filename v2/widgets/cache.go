package widgets

import "github.com/negrel/ginger/v2/render"

// Cache is used to cache the rendered frame of
// the widget and improve performance. The cache
// is returned when the components is drawn.
type Cache struct {
	valid bool
	C     Constraint
	F     *render.Frame
}

// NewCache return a new widgets cache instance.
func NewCache(co Constraint) *Cache {
	return &Cache{
		valid: false,
		C:     co,
		F:     nil,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// IsValid return wether or not the cache is valid.
func (c *Cache) IsValid() bool {
	return c.valid
}

// Invalid set the cache to an invalide state.
func (c *Cache) Invalid() {
	c.valid = false
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Pull the cached rendered frame.
func (c *Cache) Pull(co Constraint) *render.Frame {
	if c.valid && c.F != nil &&
		co.Bounds.Dx() >= c.F.Patch.Width() &&
		co.Bounds.Dy() >= c.F.Patch.Height() {

		return c.F
	}

	return nil
}

// Update the cache.
func (c *Cache) Update(co Constraint, fr *render.Frame) {
	c.C = co
	c.F = fr
	c.valid = true
}
