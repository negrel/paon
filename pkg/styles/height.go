package styles

import "github.com/negrel/debuggo/pkg/assert"

func Height(value int, unit Unit) ThemeOpt {
	assert.GreaterOrEqual(value, 0, "height must be a positive number")

	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDHeight, value, unit),
		)
	}
}

func MinHeight(value int, unit Unit) ThemeOpt {
	assert.GreaterOrEqual(value, 0, "min-height must be a positive number")

	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMinHeight, value, unit),
		)
	}
}

func MaxHeight(value int, unit Unit) ThemeOpt {
	assert.GreaterOrEqual(value, 0, "max-height must be a positive number")

	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMaxHeight, value, unit),
		)
	}
}
