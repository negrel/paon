package property

import (
	val "github.com/negrel/paon/pkg/pdk/styles/value"
)

var _ Property = Unit{}

type Unit struct {
	Prop
	val.Unit
}

func MakeUnit(id ID, unit val.Unit) Unit {
	return Unit{
		Prop: MakeProp(id),
		Unit: unit,
	}
}
