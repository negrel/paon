package styles

import "github.com/negrel/paon/styles/property"

// IntUnitStyle define objects containing property.IntUnit style properties.
type IntUnitStyle interface {
	IntUnit(property.IntUnitID) *property.IntUnit
	SetIntUnit(property.IntUnitID, *property.IntUnit)
}

var _ IntUnitStyle = intUnitStyle{}

type intUnitStyle struct {
	units []*property.IntUnit
}

// NewIntUnitStyle returns a new UnitStyle instance.
func NewIntUnitStyle() IntUnitStyle {
	return newIntUnitStyle()
}

func newIntUnitStyle() intUnitStyle {
	return intUnitStyle{
		units: make([]*property.IntUnit, property.IntUnitIDCount()+1),
	}
}

// IntUnit implements the IntUnitStyle interface.
func (us intUnitStyle) IntUnit(id property.IntUnitID) *property.IntUnit {
	return us.units[uint32(id)]
}

// SetIntUnit implements the IntUnitStyle interface.
func (us intUnitStyle) SetIntUnit(id property.IntUnitID, u *property.IntUnit) {
	us.units[uint32(id)] = u
}
