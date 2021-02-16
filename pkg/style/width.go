package style

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

func Width(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "width value must be positive")

	return property.MakeUnit(property.IDWidth, unit)
}

func MinWidth(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "min-width value must be positive")

	return property.MakeUnit(property.IDMinWidth, unit)
}

func MaxWidth(unit value.Unit) property.Unit {
	assert.GreaterOrEqual(unit.Value, 0, "max-width value must be positive")

	return property.MakeUnit(property.IDMaxWidth, unit)
}
