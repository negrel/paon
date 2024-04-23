package events

import (
	"github.com/negrel/paon/geometry"
)

var ScrollEventType = NewType("Scroll")

func ScrollListener(handler func(Event, ScrollEventData)) (Type, Listener) {
	return ScrollEventType, NewListenerFunc(func(ev Event) {
		handler(ev, ev.Data.(ScrollEventData))
	})
}

// ScrollEventData is a mouse event triggered on mouse scroll. It supports horizontal
// and vertical scroll.
type ScrollEventData struct {
	MouseEventData
	ScrollDirection
}

type ScrollDirection int8

const (
	ScrollUp ScrollDirection = iota
	ScrollDown
	ScrollLeft
	ScrollRight
)

// NewScroll returns a new scroll event.
func NewScroll(pos geometry.Vec2D, modifiers ModMask, direction ScrollDirection) Event {
	data := ScrollEventData{
		MouseEventData: MouseEventData{
			AbsPosition: pos,
			RelPosition: pos,
			Buttons:     0,
			Modifiers:   modifiers,
		},
		ScrollDirection: direction,
	}

	return NewEvent(ScrollEventType, data)
}

// Vertical returns true if scroll is vertical.
func (se ScrollEventData) Vertical() bool {
	return se.ScrollDirection == ScrollUp || se.ScrollDirection == ScrollDown
}

// Horizontal returns true if scroll is horizontal.
func (se ScrollEventData) Horizontal() bool {
	return se.ScrollDirection == ScrollLeft || se.ScrollDirection == ScrollRight
}
