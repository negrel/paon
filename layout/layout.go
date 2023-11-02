package layout

import "github.com/negrel/paon/geometry"

// Layout is a generic interface for object
// that can produce BoxModel using the given Constraint.
type Layout interface {
	Layout(Constraint) geometry.Size
}

type LayoutFunc func(Constraint) geometry.Size

// Layout implements the Layout interface.
func (lf LayoutFunc) Layout(co Constraint) geometry.Size {
	return lf(co)
}
