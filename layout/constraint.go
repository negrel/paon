package layout

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/minmax"
)

// Constraint define layout constraint.
type Constraint struct {
	MinSize, MaxSize geometry.Size
	ParentSize       geometry.Size
	RootSize         geometry.Size
}

// Equals returns true if the given Constraint is equal to this Constraint.
func (c Constraint) Equals(other Constraint) bool {
	return c.MinSize.Equals(other.MinSize) && c.MaxSize.Equals(other.MaxSize) &&
		c.ParentSize.Equals(other.ParentSize) && c.RootSize.Equals(other.RootSize)
}

// ApplyOnSize applies the size constraints on the given size and return
// the constrained size.
func (c Constraint) ApplyOnSize(size geometry.Size) geometry.Size {
	return geometry.NewSize(
		c.ApplyOnWidth(size.Width()),
		c.ApplyOnHeight(size.Height()),
	)
}

// ApplyOnWidth applies the width constraints on the given width and return the
// constrained width.
func (c Constraint) ApplyOnWidth(width int) int {
	width = minmax.Max(width, c.MinSize.Width())
	return minmax.Min(width, c.MaxSize.Width())
}

// ApplyOnHeight applies the height constraints on the given height and return the
// constrained height.
func (c Constraint) ApplyOnHeight(height int) int {
	height = minmax.Max(height, c.MinSize.Height())
	return minmax.Min(height, c.MaxSize.Height())
}
