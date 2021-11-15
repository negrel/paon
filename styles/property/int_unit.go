package property

import (
	"github.com/negrel/paon/pdk/id"
)

// IntUnitID define a unique ID accross all Unit properties.
type IntUnitID id.ID

var (
	intUnitRegistry = id.Registry{}
	intUnitMap      = id.NewStringMap()
)

// NewIntUnitID returns a new unique Unit property ID.
func NewIntUnitID(name string) IntUnitID {
	id := intUnitRegistry.New()
	intUnitMap.Set(id, name)

	return IntUnitID(id)
}

// String implements the fmt.Stringer interface.
func (iui IntUnitID) String() string {
	return intUnitMap.Get(id.ID(iui))
}

// IntUnitIDCount returns the number of UnitID generated.
func IntUnitIDCount() uint32 {
	return uint32(intUnitRegistry.Last())
}

// IntUnit defines an Int value with a unit.
type IntUnit struct {
	value int
	unit  Unit
}

// NewIntUnit returns a new IntUnit instance with the given value and unit.
func NewIntUnit(value int, unit Unit) IntUnit {
	return IntUnit{
		value: value,
		unit:  unit,
	}
}

// Value returns the stored integer value.
func (iu *IntUnit) Value() int {
	return iu.value
}

// Unit returns the Unit of this IntUnit.
func (iu *IntUnit) Unit() Unit {
	return iu.unit
}

var (
	_IDWidth         = NewIntUnitID("width")
	_IDMinWidth      = NewIntUnitID("min-width")
	_IDMaxWidth      = NewIntUnitID("max-width")
	_IDHeight        = NewIntUnitID("height")
	_IDMinHeight     = NewIntUnitID("min-height")
	_IDMaxHeight     = NewIntUnitID("max-height")
	_IDMarginLeft    = NewIntUnitID("margin-left")
	_IDMarginTop     = NewIntUnitID("margin-top")
	_IDMarginRight   = NewIntUnitID("margin-right")
	_IDMarginBottom  = NewIntUnitID("margin-bottom")
	_IDBorderLeft    = NewIntUnitID("border-left")
	_IDBorderTop     = NewIntUnitID("border-top")
	_IDBorderRight   = NewIntUnitID("border-right")
	_IDBorderBottom  = NewIntUnitID("border-bottom")
	_IDPaddingLeft   = NewIntUnitID("padding-left")
	_IDPaddingTop    = NewIntUnitID("padding-top")
	_IDPaddingRight  = NewIntUnitID("padding-right")
	_IDPaddingBottom = NewIntUnitID("padding-bottom")
)

// Width returns the IntUnitID of the "width" property.
func Width() IntUnitID {
	return _IDWidth
}

// MinWidth returns the IntUnitID of the "min-width" property.
func MinWidth() IntUnitID {
	return _IDMinWidth
}

// MaxWidth returns the IntUnitID of the "max-width" property.
func MaxWidth() IntUnitID {
	return _IDMaxWidth
}

// Height returns the IntUnitID of the "height" property.
func Height() IntUnitID {
	return _IDHeight
}

// MinHeight returns the IntUnitID of the "min-height" property.
func MinHeight() IntUnitID {
	return _IDMinHeight
}

// MaxHeight returns the IntUnitID of the "max-height" property.
func MaxHeight() IntUnitID {
	return _IDMaxHeight
}

// MarginLeft returns the IntUnitID of the "margin-left" property.
func MarginLeft() IntUnitID {
	return _IDMarginLeft
}

// MarginTop returns the IntUnitID of the "margin-top" property.
func MarginTop() IntUnitID {
	return _IDMarginTop
}

// MarginRight returns the IntUnitID of the "margin-right" property.
func MarginRight() IntUnitID {
	return _IDMarginRight
}

// MarginBottom returns the IntUnitID of the "margin-bottom" property.
func MarginBottom() IntUnitID {
	return _IDMarginBottom
}

// BorderLeft returns the IntUnitID of the "border-left" property.
func BorderLeft() IntUnitID {
	return _IDBorderLeft
}

// BorderTop returns the IntUnitID of the "border-top" property.
func BorderTop() IntUnitID {
	return _IDBorderTop
}

// BorderRight returns the IntUnitID of the "border-right" property.
func BorderRight() IntUnitID {
	return _IDBorderRight
}

// BorderBottom returns the IntUnitID of the "border-bottom" property.
func BorderBottom() IntUnitID {
	return _IDBorderBottom
}

// PaddingLeft returns the IntUnitID of the "padding-left" property.
func PaddingLeft() IntUnitID {
	return _IDPaddingLeft
}

// PaddingTop returns the IntUnitID of the "padding-top" property.
func PaddingTop() IntUnitID {
	return _IDPaddingTop
}

// PaddingRight returns the IntUnitID of the "padding-right" property.
func PaddingRight() IntUnitID {
	return _IDPaddingRight
}

// PaddingBottom returns the IntUnitID of the "padding-bottom" property.
func PaddingBottom() IntUnitID {
	return _IDPaddingBottom
}
