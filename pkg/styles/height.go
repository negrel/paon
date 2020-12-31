package styles

func Height(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDHeight, value, unit),
		)
	}
}

func MinHeight(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMinHeight, value, unit),
		)
	}
}

func MaxHeight(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMaxHeight, value, unit),
		)
	}
}
