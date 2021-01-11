package widgets

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/pkg/pdk/events"
	pdkstyles "github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/style"
)

func delStyleListener(handler func(event eventDelStyle)) *events.Listener {
	l := events.Listener{
		Type: eventTypeDelStyle,
		Handle: func(event events.Event) {
			spe, ok := event.(eventDelStyle)

			if !ok {
				log.Warnf("click listener expected %v, but got %v", eventTypeDelStyle, event.Type())
				return
			}

			handler(spe)
		},
	}

	return &l
}

var eventTypeDelStyle = events.MakeType("delete-style")

type eventDelStyle struct {
	events.Event
	pdkstyles.Style
}

func makeEventDelStyle(style pdkstyles.Style) eventDelStyle {
	return eventDelStyle{
		Event: events.MakeEvent(eventTypeDelStyle),
		Style: style,
	}
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
	Widget
	pdkstyles.Style

	styles []pdkstyles.Style // Switch to a list
}

// MakeTheme return a new Theme object with the given shared style.Style and
// a new unique style.Style object.
func MakeTheme(widget Widget, themes ...pdkstyles.Style) Theme {
	return theme{
		Widget: widget,
		Style:  pdkstyles.MakeStyle(),
		styles: themes,
	}
}

func (t theme) parent() pdkstyles.Style {
	if parent := t.Widget.Parent(); parent != nil {
		return parent.Theme()
	}

	return nil
}

// AddStyle implements the Theme interface.
func (t theme) AddStyle(s pdkstyles.Style) {
	// Watch property change to trigger redraw/reflow if needed.
	spListener := pdkstyles.SetPropertyListener(func(event pdkstyles.EventSetProperty) {
		// Nothing changed
		if t.Get(event.Old.ID()) == event.Old {
			return
		}

		ScheduleRenderingFor(t.Widget)
	})
	s.AddEventListener(spListener)

	// Delete the property change listener when the theme will be removed
	var dsListener *events.Listener
	dsListener = delStyleListener(func(event eventDelStyle) {
		s.RemoveEventListener(spListener)
		s.RemoveEventListener(dsListener)
	})

	t.styles = append(t.styles, s)
}

// DelStyle implements the Theme interface.
func (t theme) DelStyle(delStyle pdkstyles.Style) {
	for i, s := range t.styles {
		if s == delStyle {
			t.styles = append(t.styles[:i], t.styles[i+1:]...)
			t.DispatchEvent(makeEventDelStyle(s))
		}
	}
}

// Set implements the style.Style interface.
func (t theme) Set(property property.Property) {
	t.Style.Set(property)
}

func (t theme) getFromParent(id property.ID) property.Property {
	if pStyle := t.parent(); pStyle != nil {
		return pStyle.Get(id)
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
