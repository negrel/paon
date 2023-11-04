package mouse

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var ClickEventType = events.NewType("Click")

// ClickListener returns an events.Listener that will call the given handler
// on click events.
func ClickListener(handler func(ClickEvent)) (events.Type, events.Handler) {
	return ClickEventType, events.HandlerFunc(func(event events.Event) {
		handler(event.(ClickEvent))
	})
}

// ClickEvent define an event triggered by a primary mouse click, that is a mouse press
// followed by a mouse up (positions can differ). Unlike mouse up event, click events
// is dispatched to target at mouse press position.
type ClickEvent struct {
	Event
	MousePress Event
}

// NewClick returns a new click event.
func NewClick(pos geometry.Vec2D, buttons ButtonMask, mods keypress.ModMask, mousePress Event) ClickEvent {
	return ClickEvent{
		Event: Event{
			Event:       events.NewEvent(ClickEventType),
			AbsPosition: pos,
			RelPosition: pos,
			Buttons:     buttons,
			Modifiers:   mods,
		},
		MousePress: mousePress,
	}
}
