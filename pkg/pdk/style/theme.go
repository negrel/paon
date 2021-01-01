package style

import "github.com/negrel/paon/pkg/pdk/style/property"

type Theme interface {
	priority() int

	Set(property.Property)
	Get(property.ID) property.Property
}

var _ Theme = theme{}

type theme struct {
	_priority int

	onPropertyChange func(old, new property.Property)
	props            map[property.ID]property.Property
}

func (t theme) priority() int {
	return t._priority
}

func (t theme) Set(prop property.Property) {
	if prop == nil {
		return
	}

	old := t.props[prop.ID()]
	if old != prop {
		t.onPropertyChange(old, prop)
		t.props[prop.ID()] = prop
	}
}

func (t theme) Get(id property.ID) property.Property {
	return t.props[id]
}
