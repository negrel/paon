package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/styles/value"
)

// Constraint define the constraint an flow must respect.
type Constraint struct {
	Min, Max   geometry.Rectangle
	ParentSize geometry.Size
	RootSize   geometry.Size
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
		c.ParentSize.Equals(other.ParentSize) && c.RootSize.Equals(other.RootSize)
}

// EqualsSize returns true if the given Constraint has the same min and
// max size.
func (c Constraint) EqualsSize(other Constraint) bool {
	return c.Min.Size().Equals(other.Min.Size()) && c.Max.Size().Equals(other.Max.Size())
}

// ToCellUnit converts the given value to a Cell based value.Unit and returns it.
func (c Constraint) ToCellUnit(unit value.Unit) int {
	switch unit.UnitID {
	case value.CellUnit:
		return unit.Value

	case value.PercentageWidthUnit:
		return c.ParentSize.Width() / 100 * unit.Value

	case value.PercentageHeightUnit:
		return c.ParentSize.Height() / 100 * unit.Value

	case value.WindowWidthUnit:
		return c.RootSize.Width() / 100 * unit.Value

	case value.WindowHeightUnit:
		return c.RootSize.Height() / 100 * unit.Value

	default:
		panic("invalid unit")
	}
}
