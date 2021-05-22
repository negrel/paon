package property

import (
	val "github.com/negrel/paon/pdk/styles/value"
)

var _ Property = Unit{}

type Unit struct {
	Prop
	Value val.Unit
}

func MakeUnit(id ID, unit val.Unit) Unit {
	return Unit{
		Prop:  MakeProp(id),
		Value: unit,
	}
}
