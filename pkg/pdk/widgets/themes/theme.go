package themes

import (
	"github.com/negrel/paon/pkg/pdk/events"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/style"
)

type Themed interface {
	Theme() Theme
}

// Theme is a composition of multiple style.Style object. Each widget have a unique
// style.Style and may have shared across multiple widgets) style.Style object.
type Theme interface {
	pdkstyles.Style

	// AddStyle add the given style.Style of the Widget style list.
	AddStyle(pdkstyles.Style)
	// DelStyle delete the given style.Style of the Widget style list.
	DelStyle(pdkstyles.Style)
}

var _ Theme = theme{}

type theme struct {
	pdkstyles.Style

	getParent func() Themed
	styles    []pdkstyles.Style // Switch to a list
}

// Make return a new Theme object with the given shared style.Style and
// a new unique style.Style object.
func Make(getParent func() Themed) Theme {
	return theme{
		getParent: getParent,
		Style:     pdkstyles.MakeStyle(),
		styles:    make([]pdkstyles.Style, 8),
	}
}

// AddStyle implements the Theme interface.
func (t theme) AddStyle(s pdkstyles.Style) {
	// Watch property change to trigger redraw/reflow if needed.
	spListener := pdkstyles.PropertyChangeListener(func(event pdkstyles.EventPropertyChange) {
		// Nothing changed
		if t.Get(event.Old.ID()) == event.Old {
			return
		}
	})
	s.AddEventListener(spListener)

	// Delete the property change listener when the theme will be removed
	var dsListener *events.Listener
	dsListener = ThemeChangeListener(func(event EventThemeChange) {
		if event.DeletedStyle {
			s.RemoveEventListener(spListener)
			s.RemoveEventListener(dsListener)
		}
	})

	t.styles = append(t.styles, s)
}

// DelStyle implements the Theme interface.
func (t theme) DelStyle(delStyle pdkstyles.Style) {
	for i, s := range t.styles {
		if s == delStyle {
			t.styles = append(t.styles[:i], t.styles[i+1:]...)
			t.DispatchEvent(makeEventThemeChange(s, true))
		}
	}
}

// Set implements the style.Style interface.
func (t theme) Set(property property.Property) {
	t.Style.Set(property)
}

func (t theme) getFromParent(id property.ID) property.Property {
	if parent := t.getParent(); parent != nil {
		if pTheme := parent.Theme(); pTheme != nil {
			return pTheme.Get(id)
		}
	}

	return nil
}

func (t theme) get(id property.ID) property.Property {
	if prop := t.Style.Get(id); prop != nil {
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

// Get implements the style.Style interface.
func (t theme) Get(id property.ID) property.Property {
	prop := t.get(id)
	if _, isInherited := prop.(style.InheritedProp); isInherited {
		return t.getFromParent(id)
	}

	return prop
}

// Del implements the style.Style interface.
func (t theme) Del(id property.ID) {
	t.Style.Del(id)
}
