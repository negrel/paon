package styles

func Width(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDWidth, value, unit),
		)
	}
}

func MinWidth(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMinWidth, value, unit),
		)
	}
}

func MaxWidth(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMaxWidth, value, unit),
		)
	}
}
