package styles

import (
	"sync"

	"github.com/negrel/paon/pkg/pdk/styles/property"
)

type _style Style

// Theme define a composition of style.
type Theme interface {
	Style

	// AddStyle adds the given Style to the StyleList.
	AddStyle(Style)
	// DelStyle deletes the given Style from the StyleList.
	DelStyle(Style)

	// Styles returns all the Style present in this Theme.
	Styles() []Style
}

// theme is a composition of Style object.
type theme struct {
	sync.RWMutex

	_style
	shared []Style
}

// NewTheme return a new Theme object with the given internal Style.
func NewTheme(defaultStyle Style) Theme {
	shared := make([]Style, 0, 8)
	shared = append(shared, defaultStyle)

	return &theme{
		_style: NewStyle(),
		shared: shared,
	}
}

// Get implements the Style interface.
func (t *theme) Get(id property.ID) property.Property {
	t.RLock()
	defer t.RUnlock()

	if prop := t._style.Get(id); prop != nil {
		return prop
	}

	for i := len(t.shared) - 1; i >= 0; i-- {
		if prop := t.shared[i].Get(id); prop != nil {
			return prop
		}
	}

	return nil
}

// Styles implements the Theme interface.
func (t *theme) Styles() []Style {
	t.RLock()
	defer t.RUnlock()

	return t.shared
}

// AddStyle implements the Theme interface.
func (t *theme) AddStyle(s Style) {
	panic("implement me")
}

// DelStyle implements the Theme interface.
func (t *theme) DelStyle(s Style) {
	panic("implement me")
}
