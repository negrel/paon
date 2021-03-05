package flows

import "github.com/negrel/paon/internal/geometry"

// Constraint define the constraint an Flow must respect.
type Constraint struct {
	Min, Max geometry.Rectangle
}

// Equals returns true if the given Constraint is equal to this Constraint.
func (c Constraint) Equals(other Constraint) bool {
	return c.Min.Equals(other.Min) && c.Max.Equals(other.Max)
}
