package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func Height(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "height value must have a positive value")

	return property.MakeUnit(property.IDHeight, unit)
}

func MinHeight(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "min-height value must have a positive value")

	return property.MakeUnit(property.IDMinHeight, unit)
}

func MaxHeight(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "max-height value must have a positive value")

	return property.MakeUnit(property.IDMaxHeight, unit)
}
