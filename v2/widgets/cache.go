package widgets

import (
	"image"

	"github.com/negrel/ginger/v2/render"
)

// Cache is used to cache the rendered frame of
// the widget and improve performance. The cache
// is returned when the components is drawn.
type Cache struct {
	valid bool
	B     image.Rectangle
	F     *render.Frame
}

// NewCache return a new widgets cache instance.
func NewCache(bounds image.Rectangle) *Cache {
	return &Cache{
		valid: false,
		B:     bounds,
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
func (c *Cache) Pull(bounds image.Rectangle) *render.Frame {
	if c.valid && c.F != nil &&
		bounds.Dx() >= c.F.Patch.Width() &&
		bounds.Dy() >= c.F.Patch.Height() {

		return c.F
	}

	return nil
}

// Update the cache.
func (c *Cache) Update(bounds image.Rectangle, fr *render.Frame) {
	c.B = bounds
	c.F = fr
	c.valid = true
}
