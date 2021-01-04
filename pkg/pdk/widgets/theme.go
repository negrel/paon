package widgets

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/style"
)

type Theme interface {
	pdkstyle.Style

	AddStyle(pdkstyle.Style)
	DelStyle(pdkstyle.Style)
}

var _ Theme = theme{}

type theme struct {
	widget *widget

	widgetStyle pdkstyle.Style
	styles      []pdkstyle.Style
}

func MakeTheme(themes ...pdkstyle.Style) Theme {
	return theme{
		widgetStyle: pdkstyle.MakeStyle(),
		styles:      themes,
	}
}

func (t theme) AddStyle(s pdkstyle.Style) {
	t.styles = append(t.styles, s)
}

func (t theme) DelStyle(delStyle pdkstyle.Style) {
	for i, s := range t.styles {
		if s == delStyle {
			t.styles = append(t.styles[:i], t.styles[i+1:]...)
		}
	}
}

func (t theme) Set(property property.Property) {
	t.widgetStyle.Set(property)
}

func (t theme) getFromParent(id property.ID) property.Property {
	if wParent := t.widget.Parent(); wParent != nil {
		return wParent.Theme().Get(id)
	}

	return nil
}

func (t theme) Get(id property.ID) property.Property {
	prop := t.get(id)
	if _, isInherited := prop.(style.InheritedProp); isInherited {
		return t.getFromParent(id)
	}

	return prop
}

func (t theme) get(id property.ID) property.Property {
	if prop := t.widgetStyle.Get(id); prop != nil {
		return prop
	}

	// Find the first theme with the given property
	for _, theme := range t.styles {
		if prop := theme.Get(id); prop != nil {
			return prop
		}
	}

	return nil
}

func (t theme) Del(id property.ID) {
	t.widgetStyle.Del(id)
}
