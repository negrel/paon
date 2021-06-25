package property

import (
	val "github.com/negrel/paon/styles/value"
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
