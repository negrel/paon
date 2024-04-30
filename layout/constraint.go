package layout

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/minmax"
)

// Constraint define layout constraint.
type Constraint struct {
	MinSize geometry.Size
	MaxSize geometry.Size
}

// ApplyOnSize applies the size constraints on the given size and return
// the constrained size.
func (c Constraint) ApplyOnSize(size geometry.Size) geometry.Size {
	return geometry.Size{
		Width:  c.ApplyOnWidth(size.Width),
		Height: c.ApplyOnHeight(size.Height),
	}
}

// ApplyOnWidth applies the width constraints on the given width and return the
// constrained width.
func (c Constraint) ApplyOnWidth(width int) int {
	width = minmax.Max(width, c.MinSize.Width)
	return minmax.Min(width, c.MaxSize.Width)
}

// ApplyOnHeight applies the height constraints on the given height and return the
// constrained height.
func (c Constraint) ApplyOnHeight(height int) int {
	height = minmax.Max(height, c.MinSize.Height)
	return minmax.Min(height, c.MaxSize.Height)
}

// ForceSize sets min and max size to the given one.
func (c Constraint) ForceSize(size geometry.Size) Constraint {
	c.MinSize = size
	c.MaxSize = size
	return c
}
