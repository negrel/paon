package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func PaddingLeft(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "Padding-left value must have a positive value")

	return property.MakeUnit(property.PaddingLeft(), unit)
}

func PaddingTop(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "Padding-top value must have a positive value")

	return property.MakeUnit(property.PaddingTop(), unit)
}

func PaddingRight(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "Padding-right value must have a positive value")

	return property.MakeUnit(property.PaddingRight(), unit)
}

func PaddingBottom(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "Padding-bottom value must have a positive value")

	return property.MakeUnit(property.PaddingBottom(), unit)
}
