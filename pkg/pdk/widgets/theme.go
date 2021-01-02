package widgets

import (
	pdkstyle "github.com/negrel/paon/pkg/pdk/style"
	"github.com/negrel/paon/pkg/pdk/style/property"
	"github.com/negrel/paon/pkg/style"
)

var _ pdkstyle.Theme = theme{}

// Widget theme is composed of multiple style.Theme object.
type theme struct {
	widget *widget

	widgetTheme pdkstyle.Theme
	themes      []pdkstyle.Theme
}

func MakeTheme(themes ...pdkstyle.Theme) pdkstyle.Theme {
	return theme{
		widgetTheme: pdkstyle.MakeTheme(),
		themes:      themes,
	}
}

func (t theme) Set(property property.Property) {
	t.widgetTheme.Set(property)
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
	if prop := t.widgetTheme.Get(id); prop != nil {
		return prop
	}

	// Find the first theme with the given property
	for _, theme := range t.themes {
		if prop := theme.Get(id); prop != nil {
			return prop
		}
	}

	return nil
}

func (t theme) Del(id property.ID) {
	t.widgetTheme.Del(id)
}
