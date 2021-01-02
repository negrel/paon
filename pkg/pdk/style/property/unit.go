package property

import (
	val "github.com/negrel/paon/pkg/pdk/style/value"
)

var _ Property = Unit{}

type Unit struct {
	Prop
	unit val.Unit
}

func MakeUnit(id ID, unit val.Unit) Unit {
	return Unit{
		Prop: MakeProp(id),
		unit: unit,
	}
}

func (up Unit) Value() val.Unit {
	return up.unit
}

func (up *Unit) SetValue(value int, unit val.UnitID) {
	up.unit.Value = value
	up.unit.ID = unit
}
