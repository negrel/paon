package widgets

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/style"
)

// Theme is a composition of multiple style.Style object. Each widget have a unique
// style.Style and may have shared across multiple widgets) style.Style object.
type Theme interface {
	pdkstyle.Style

	// AddStyle add the given style.Style of the Widget style list.
	AddStyle(pdkstyle.Style)
	// DelStyle delete the given style.Style of the Widget style list.
	DelStyle(pdkstyle.Style)
}

var _ Theme = theme{}

type theme struct {
	widget *widget

	widgetStyle pdkstyle.Style
	styles      []pdkstyle.Style
}

// MakeTheme return a new Theme object with the given shared style.Style and
// a new unique style.Style object.
func MakeTheme(themes ...pdkstyle.Style) Theme {
	return theme{
		widgetStyle: pdkstyle.MakeStyle(),
		styles:      themes,
	}
}

// AddStyle implements the Theme interface.
func (t theme) AddStyle(s pdkstyle.Style) {
	t.styles = append(t.styles, s)
}

// DelStyle implements the Theme interface.
func (t theme) DelStyle(delStyle pdkstyle.Style) {
	for i, s := range t.styles {
		if s == delStyle {
			t.styles = append(t.styles[:i], t.styles[i+1:]...)
		}
	}
}

// Set implements the style.Style interface.
func (t theme) Set(property property.Property) {
	t.widgetStyle.Set(property)
}

func (t theme) getFromParent(id property.ID) property.Property {
	if wParent := t.widget.Parent(); wParent != nil {
		return wParent.Theme().Get(id)
	}

	return nil
}

// Get implements the style.Style interface.
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

// Del implements the style.Style interface.
func (t theme) Del(id property.ID) {
	t.widgetStyle.Del(id)
}
