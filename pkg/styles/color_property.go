package styles

var _ Property = ColorProperty{}

type ColorProperty struct {
	property
	Value ColorValue
}

func makeColorProp(id PropID, value ColorValue) ColorProperty {
	return ColorProperty{
		property: prop(id),
		Value:    value,
	}
}
