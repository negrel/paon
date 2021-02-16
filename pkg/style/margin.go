package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func MarginLeft(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "margin-left value must have a positive value")

	return property.MakeUnit(property.IDMarginLeft, unit)
}

func MarginTop(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "margin-top value must have a positive value")

	return property.MakeUnit(property.IDMarginTop, unit)
}

func MarginRight(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "margin-right value must have a positive value")

	return property.MakeUnit(property.IDMarginRight, unit)
}

func MarginBottom(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "margin-bottom value must have a positive value")

	return property.MakeUnit(property.IDMarginBottom, unit)
}
