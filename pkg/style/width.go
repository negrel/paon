package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/style/property"
	"github.com/negrel/paon/pkg/pdk/style/value"
)

func Width(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "width value must have a positive value")

	return property.MakeUnit(property.IDWidth, unit)
}

func MinWidth(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "min-width value must have a positive value")

	return property.MakeUnit(property.IDMinWidth, unit)
}

func MaxWidth(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "max-width value must have a positive value")

	return property.MakeUnit(property.IDMaxWidth, unit)
}
