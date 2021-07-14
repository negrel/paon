package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles/value"
)

// Constraint define the constraint an flow must respect.
type Constraint struct {
	Min, Max   geometry.Rectangle
	parentSize geometry.Size
	rootSize   geometry.Size
}

// NewConstraint returns a new Constraint object with the given constraint.
func NewConstraint(min, max geometry.Rectangle, parent, root geometry.Size) Constraint {
	return Constraint{
		Min:        min,
		Max:        max,
		parentSize: parent,
		rootSize:   root,
	}
}

// SetMin sets the constraint minimum rectangle.
func (c Constraint) SetMin(min geometry.Rectangle) Constraint {
	c.Min = min
	return c
}

// SetMax sets the constraint maximum rectangle.
func (c Constraint) SetMax(max geometry.Rectangle) Constraint {
	c.Max = max
	return c
}

// Equals returns true if the given Constraint is equal to this Constraint.
func (c Constraint) Equals(other Constraint) bool {
	return c.Min.Equals(other.Min) && c.Max.Equals(other.Max) &&
		c.parentSize.Equals(other.parentSize) && c.rootSize.Equals(other.rootSize)
}

// ToCellUnit converts the given value to a Cell based value.Unit and returns it.
func (c Constraint) ToCellUnit(unit value.Unit) int {
	switch unit.UnitID {
	case value.CellUnit:
		return unit.Value

	case value.PercentageWidthUnit:
		return c.parentSize.Width() / 100 * unit.Value

	case value.PercentageHeightUnit:
		return c.parentSize.Height() / 100 * unit.Value

	case value.WindowWidthUnit:
		return c.rootSize.Width() / 100 * unit.Value

	case value.WindowHeightUnit:
		return c.rootSize.Height() / 100 * unit.Value

	default:
		panic("invalid unit")
	}
}
