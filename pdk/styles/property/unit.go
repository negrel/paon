package property

import (
	val "github.com/negrel/paon/pdk/styles/value"
)

var _ Property = Unit{}

type Unit struct {
	Prop
	Value val.Unit
}

func NewUnit(id ID, unit val.Unit) Unit {
	return Unit{
		Prop:  NewProp(id),
		Value: unit,
	}
}
