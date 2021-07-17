package styles

import (
	"sort"
	"sync"

	"github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/styles/property"
)

// Themed define a generic interface for objects
// containing a Theme object.
type Themed interface {
	Theme() Theme
}

// Theme define a set of Style. These styles are read-only and can't be modified
// throught the Theme object. Theme also implements the Style interface because
// it embed a read-write Style object.
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
	sync.RWMutex
	events.Target

	style  Style
	shared styleSlice
	sorted bool
}

// NewTheme return a new Theme object with the given internal Style.
func NewTheme(defaultStyle WeightedStyle) Theme {
	shared := make([]WeightedStyle, 0, 8)
	shared = append(shared, defaultStyle)

	return &theme{
		style:  New(),
		shared: shared,
	}
}

func (t *theme) sortShared() {
	t.Lock()
	defer t.Unlock()

	sort.Sort(t.shared)
}

// Get implements the Style interface.
func (t *theme) Get(id property.ID) property.Property {
	if !t.sorted {
		t.sortShared()
	}

	t.RLock()
	defer t.RUnlock()

	if prop := t.style.Get(id); prop != nil {
		return prop
	}

	for i := len(t.shared) - 1; i >= 0; i-- {
		if prop := t.shared[i].Get(id); prop != nil {
			return prop
		}
	}

	return nil
}

// Set implements the Style interface.
func (t *theme) Set(prop property.Property) {
	t.style.Set(prop)
}

// Del implements the Style interface.
func (t *theme) Del(prop property.ID) {
	t.style.Del(prop)
}

// Styles implements the Theme interface.
func (t *theme) Styles() []WeightedStyle {
	t.RLock()
	defer t.RUnlock()

	return t.shared
}

// AddStyle implements the Theme interface.
func (t *theme) AddStyle(s WeightedStyle) {
	t.shared = append(t.shared, s)
	sort.Sort(t.shared)
}

// DelStyle implements the Theme interface.
func (t *theme) DelStyle(s WeightedStyle) {
	for i, style := range t.shared {
		if style == s {
			t.shared[len(t.shared)-1], t.shared[i] = t.shared[i], t.shared[len(t.shared)-1]
			t.shared = t.shared[:len(t.shared)-1]
		}
	}
	sort.Sort(t.shared)
}
