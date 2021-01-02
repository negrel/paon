package style

import "github.com/negrel/paon/pkg/pdk/style/property"

type Theme interface {
	Set(property.Property)
	Get(property.ID) property.Property
	Del(property.ID)
}

var _ Theme = theme{}

type theme struct {
	props map[property.ID]property.Property
}

func MakeTheme() Theme {
	return theme{
		props: make(map[property.ID]property.Property, 8),
	}
}

func (t theme) Del(id property.ID) {
	delete(t.props, id)
}

func (t theme) Set(prop property.Property) {
	if prop == nil {
		return
	}

	t.props[prop.ID()] = prop
}

func (t theme) Get(id property.ID) property.Property {
	return t.props[id]
}
