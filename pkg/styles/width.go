package styles

import "github.com/negrel/debuggo/pkg/assert"

func Width(value int, unit Unit) ThemeOpt {
	assert.GreaterOrEqual(value, 0, "width must be a positive number")

	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProp(PropIDWidth, value, unit),
		)
	}
}

func MinWidth(value int, unit Unit) ThemeOpt {
	assert.GreaterOrEqual(value, 0, "min-width must be a positive number")

	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProp(PropIDMinWidth, value, unit),
		)
	}
}

func MaxWidth(value int, unit Unit) ThemeOpt {
	assert.GreaterOrEqual(value, 0, "max-width must be a positive number")

	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProp(PropIDMaxWidth, value, unit),
		)
	}
}
