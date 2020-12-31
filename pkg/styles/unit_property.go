package styles

var _ Property = UnitProperty{}

type UnitProperty struct {
	property
	Value UnitValue
}

func makeUnitProp(id PropID, value int, unit Unit) UnitProperty {
	return UnitProperty{
		property: prop(id),
		Value:    UnitValue{value, unit},
	}
}
