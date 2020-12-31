package styles

func Padding(value int, unit Unit) ThemeOpt {
	return func(_ Theme) {
		PaddingX(value, unit)
		PaddingY(value, unit)
	}
}

func PaddingX(value int, unit Unit) ThemeOpt {
	return func(_ Theme) {
		PaddingLeft(value, unit)
		PaddingRight(value, unit)
	}
}

func PaddingY(value int, unit Unit) ThemeOpt {
	return func(_ Theme) {
		PaddingTop(value, unit)
		PaddingBottom(value, unit)
	}
}

func PaddingLeft(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDPaddingLeft, value, unit),
		)
	}
}

func PaddingTop(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDPaddingTop, value, unit),
		)
	}
}

func PaddingRight(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDPaddingRight, value, unit),
		)
	}
}

func PaddingBottom(value int, unit Unit) ThemeOpt {
	return func(theme Theme) {
		theme.SetUnitProp(
			makeUnitProperty(PropIDPaddingBottom, value, unit),
		)
	}
}
