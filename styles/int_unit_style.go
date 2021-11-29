package styles

import (
	"unsafe"

	"github.com/negrel/paon/pdk/id"
	"github.com/negrel/paon/pdk/id/store"
	"github.com/negrel/paon/styles/property"
)

// IntUnitStyle define objects containing property.IntUnit style properties.
type IntUnitStyle interface {
	IntUnit(property.IntUnitID) *property.IntUnit
	SetIntUnit(property.IntUnitID, *property.IntUnit)
}

var _ IntUnitStyle = intUnitStyle{}

type intUnitStyle struct {
	units store.PtrSlice
}

// NewIntUnitStyle returns a new UnitStyle instance.
func NewIntUnitStyle() IntUnitStyle {
	return newIntUnitStyle()
}

func newIntUnitStyle() intUnitStyle {
	return intUnitStyle{
		units: store.NewPtrSlice(int(property.IntUnitIDCount() + 1)),
	}
}

// IntUnit implements the IntUnitStyle interface.
func (ius intUnitStyle) IntUnit(i property.IntUnitID) *property.IntUnit {
	return (*property.IntUnit)(ius.units.Get(id.ID(i)))
}

// SetIntUnit implements the IntUnitStyle interface.
func (ius intUnitStyle) SetIntUnit(i property.IntUnitID, u *property.IntUnit) {
	ius.units.Set(id.ID(i), unsafe.Pointer(u))
}
