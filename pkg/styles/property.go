package styles

type Property interface {
	ID() PropID
	IsInherited() bool
}

type property struct {
	id        PropID
	inherited bool
}

func prop(id PropID) property {
	return property{
		id:        id,
		inherited: false,
	}
}

func (p property) ID() PropID {
	return p.id
}

func (p property) IsInherited() bool {
	return p.inherited
}

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
