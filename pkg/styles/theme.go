package styles

type ThemeOpt func(Theme)

type Theme interface {
	priority() int8

	GetUnitProp(PropID) *UnitProperty
	SetUnitProp(UnitProperty)
}

type theme struct {
	_priority int8

	unitProps map[PropID]*UnitProperty
}

func (t *theme) priority() int8 {
	return t._priority
}

func (t *theme) GetUnitProp(id PropID) *UnitProperty {
	return t.unitProps[id]
}

func (t *theme) SetUnitProp(prop *UnitProperty) {
	t.unitProps[prop.ID()] = prop
}
