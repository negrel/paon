package render

import "github.com/negrel/ginger/v3/utils"

// Constraints is the constraint object passed
// down the tree for rendering widgets.
type Constraints struct {
	// Min size of the Render Object.
	Min utils.Size
	// Max size of the Render Object.
	Max utils.Size
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// HasBoundedHeight return wether there is an
// upper bound on the maximum height.
func (c Constraints) HasBoundedHeight() bool {
	return c.Max.Height() != utils.Infinite
}

// HasBoundedWidth return wether there is an
// upper bound on the maximum width.
func (c Constraints) HasBoundedWidth() bool {
	return c.Max.Width() != utils.Infinite
}

// HasInfiniteHeight return wether the height
// constraint is infinite.
func (c Constraints) HasInfiniteHeight() bool {
	return !c.HasBoundedHeight()
}

// HasInfiniteWidth return wether the width
// constraint is infinite.
func (c Constraints) HasInfiniteWidth() bool {
	return !c.HasBoundedWidth()
}

// HasTightWidth return wether there is exactly one
// size that satisfies the constraint.
func (c Constraints) HasTightWidth() bool {
	return c.Max.Width() == c.Min.Width()
}

// HasTightHeight return wether there is exactly one
// size that satisfies the constraint.
func (c Constraints) HasTightHeight() bool {
	return c.Max.Height() == c.Min.Height()
}

// IsTight retuen whether there is exactly one size
// that satisfies the constraints.
func (c Constraints) IsTight() bool {
	return c.Min.Width() == c.Max.Width() &&
		c.Min.Height() == c.Max.Height()
}
