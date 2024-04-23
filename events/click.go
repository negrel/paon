package events

import (
	"github.com/negrel/paon/geometry"
)

var ClickEventType = NewType("Click")

// ClickListener returns a listener that will call the given handler
// on click events.
func ClickListener(handler func(Event, ClickEventData)) (Type, Listener) {
	return ClickEventType, NewListenerFunc(func(event Event) {
		handler(event, event.Data.(ClickEventData))
	})
}

// ClickEventData define data contained in a click event.
type ClickEventData struct {
	MouseEventData
	MousePress MouseEventData
}

// NewClick returns a new click event triggered by a primary mouse click, that
// is a mouse press followed by a mouse up (positions can differ). Unlike mouse
// up event, click events is dispatched to target at mouse press position.
func NewClick(pos geometry.Vec2D, buttons ButtonMask, mods ModMask, mousePress MouseEventData) Event {
	data := ClickEventData{
		MouseEventData: MouseEventData{
			AbsPosition: pos,
			RelPosition: pos,
			Buttons:     buttons,
			Modifiers:   mods,
		},
		MousePress: mousePress,
	}

	return NewEvent(ClickEventType, data)
}

// RelativePosition implements PointerEventData.
func (ced ClickEventData) RelativePosition() geometry.Vec2D {
	return ced.RelPosition
}

// WithPositionRelativeToOrigin implements PointerEventData.
func (ced ClickEventData) WithPositionRelativeToOrigin(origin geometry.Vec2D) PointerEventData {
	ced.RelPosition = ced.RelPosition.Sub(origin)
	ced.MousePress.RelPosition = ced.RelPosition
	return ced
}
