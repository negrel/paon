package styles

import (
	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
)

// Themed define a generic interface for objects
// containing a Theme object.
type Themed interface {
	Theme() Theme
}

// Theme define a set of Style. These styles are read-only and can't be modified
// through the Theme object. Getting a property value returns the value of the
// Style with the biggest weight containing the property.
//
// Theme also implements the Style interface because it embed one read-write Style object.
// It has a higher weight than any other Style objects.
type Theme interface {
	Style

	// AddStyle adds the given Style to the StyleList.
	AddStyle(WeightedStyle)
	// DelStyle deletes the given Style from the StyleList.
	DelStyle(WeightedStyle)

	// Styles returns all the Style present in this Theme.
	Styles() []WeightedStyle
}

// theme is a composition of Style object.
type theme struct {
	Style

	cache  Style
	shared styleSlice
	sorted bool
}

// NewTheme return a new Theme object with the given internal Style.
func NewTheme(defaultStyle WeightedStyle) Theme {
	return newTheme(defaultStyle)
}

var noOpTarget = events.NewNoOpTarget()

func newTheme(defaultStyle WeightedStyle) *theme {
	shared := make([]WeightedStyle, 0, 8)
	if defaultStyle != nil {
		shared = append(shared, defaultStyle)
	}
	
	return &theme{
		cache:  New(noOpTarget),
		Style:  New(events.NewTarget()),
		shared: shared,
	}
}

// Int implements the IntStyle interface.
func (t *theme) Int(id property.IntID) *property.Int {
	if prop := t.cache.Int(id); prop != nil {
		return prop
	}

	prop := t.getInt(id)
	t.cache.SetInt(id, prop)

	return prop
}

func (t *theme) getInt(id property.IntID) *property.Int {
	if prop := t.Style.Int(id); prop != nil {
		return prop
	}

	for i := len(t.shared) - 1; i >= 0; i-- {
		if prop := t.shared[i].Int(id); prop != nil {
			return prop
		}
	}

	return nil
}

// Color implements the ColorStyle interface.
func (t *theme) Color(id property.ColorID) *property.Color {
	if prop := t.cache.Color(id); prop != nil {
		return prop
	}

	prop := t.getColor(id)
	t.cache.SetColor(id, prop)

	return prop
}

func (t *theme) getColor(id property.ColorID) *property.Color {
	if prop := t.Style.Color(id); prop != nil {
		return prop
	}

	for i := len(t.shared) - 1; i >= 0; i-- {
		if prop := t.shared[i].Color(id); prop != nil {
			return prop
		}
	}

	return nil
}

// Iface implements the IfaceStyle interface.
func (t *theme) Iface(id property.IfaceID) interface{} {
	if prop := t.cache.Iface(id); prop != nil {
		return prop
	}

	prop := t.getIface(id)
	t.cache.SetIface(id, prop)

	return prop
}

func (t *theme) getIface(id property.IfaceID) interface{} {
	if prop := t.Style.Iface(id); prop != nil {
		return prop
	}

	for i := len(t.shared) - 1; i >= 0; i-- {
		if prop := t.shared[i].Iface(id); prop != nil {
			return prop
		}
	}

	return nil
}

// IntUnit implements the IntUnitStyle interface.
func (t *theme) IntUnit(id property.IntUnitID) *property.IntUnit {
	if prop := t.cache.IntUnit(id); prop != nil {
		return prop
	}

	prop := t.getIntUnit(id)
	t.cache.SetIntUnit(id, prop)

	return prop
}

func (t *theme) getIntUnit(id property.IntUnitID) *property.IntUnit {
	if prop := t.Style.IntUnit(id); prop != nil {
		return prop
	}

	for i := len(t.shared) - 1; i >= 0; i-- {
		if prop := t.shared[i].IntUnit(id); prop != nil {
			return prop
		}
	}

	return nil
}

// Styles implements the Theme interface.
func (t *theme) Styles() []WeightedStyle {
	return t.shared
}

// AddStyle implements the Theme interface.
func (t *theme) AddStyle(s WeightedStyle) {
	t.sorted = false
	t.shared = append(t.shared, s)

	// Empty cache
	t.cache = New(nil)
}

// DelStyle implements the Theme interface.
func (t *theme) DelStyle(s WeightedStyle) {
	t.sorted = false
	for i, style := range t.shared {
		if style == s {
			t.shared[len(t.shared)-1], t.shared[i] = t.shared[i], t.shared[len(t.shared)-1]
			t.shared = t.shared[:len(t.shared)-1]

			// Empty cache
			t.cache = New(nil)
			return
		}
	}
}
