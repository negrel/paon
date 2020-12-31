package styles

type ThemeOpt func(Theme)

type Theme interface {
	priority() int8

	GetUnitProp(PropID) *UnitProperty
	SetUnitProp(UnitProperty)

	GetColorProp(PropID) *ColorProperty
	SetColorProp(ColorProperty)
}

var _ Theme = &theme{}

type theme struct {
	_priority int8

	unitProps  map[PropID]*UnitProperty
	colorProps map[PropID]*ColorProperty
}

func (t *theme) priority() int8 {
	return t._priority
}

func (t *theme) GetUnitProp(id PropID) *UnitProperty {
	return t.unitProps[id]
}

func (t *theme) SetUnitProp(prop UnitProperty) {
	t.unitProps[prop.ID()] = &prop
}

func (t *theme) GetColorProp(id PropID) *ColorProperty {
	return t.colorProps[id]
}

func (t *theme) SetColorProp(prop ColorProperty) {
	t.colorProps[prop.ID()] = &prop
}
