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

// Theme define a styles.Style object composed of multiple styles.Style object.
type Theme interface {
	pdkstyles.Style

	// AddStyle adds the given style.Style of the Widget style list.
	AddStyle(pdkstyles.Style)
	// DelStyle deletes the given style.Style of the Widget style list.
	DelStyle(pdkstyles.Style)
}

var _ Theme = theme{}

type theme struct {
	pdkstyles.Style

	getParent func() Themed
	styles    styleList
}

// New returns a new Theme object.
func New(getParent func() Themed) Theme {
	return theme{
		getParent: getParent,
		Style:     pdkstyles.MakeStyle(),
		styles:    styleList{},
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

	t.styles.append(s)
}

// DelStyle implements the Theme interface.
func (t theme) DelStyle(delStyle pdkstyles.Style) {
	t.styles.remove(delStyle)
	t.DispatchEvent(makeEventThemeChange(delStyle, true))
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
	for _, theme := range t.styles.values() {
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
