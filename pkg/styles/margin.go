package styles

func Margin(value int, unit Unit) ThemeOpt {
	return func(_ Theme) {
		MarginX(value, unit)
		MarginY(value, unit)
	}
}

func MarginX(value int, unit Unit) ThemeOpt {
	return func(_ Theme) {
		MarginLeft(value, unit)
		MarginRight(value, unit)
	}
}

func MarginY(value int, unit Unit) ThemeOpt {
	return func(_ Theme) {
		MarginTop(value, unit)
		MarginBottom(value, unit)
	}
}

func MarginLeft(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMarginLeft, value, unit),
		)
	}
}

func MarginTop(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMarginTop, value, unit),
		)
	}
}

func MarginRight(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMarginRight, value, unit),
		)
	}
}

func MarginBottom(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDMarginBottom, value, unit),
		)
	}
}
