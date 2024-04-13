package mouse

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/geometry"
)

var ScrollEventType = events.NewType("Scroll")

func ScrollListener(handler func(ev ScrollEvent)) (events.Type, events.Handler) {
	return ScrollEventType, events.HandlerFunc(func(ev events.Event) {
		handler(ev.(ScrollEvent))
	})
}

// ScrollEvent is a mouse event triggered on mouse scroll. It supports horizontal
// and vertical scroll.
type ScrollEvent struct {
	Event
	ScrollDirection
}

type ScrollDirection int8

const (
	ScrollUp ScrollDirection = iota
	ScrollDown
	ScrollLeft
	ScrollRight
)

func NewScroll(pos geometry.Vec2D, modifiers keypress.ModMask, direction ScrollDirection) ScrollEvent {
	return ScrollEvent{
		Event: Event{
			Event:       events.NewEvent(ScrollEventType),
			AbsPosition: pos,
			RelPosition: pos,
			Buttons:     0,
			Modifiers:   modifiers,
		},
		ScrollDirection: direction,
	}
}

func (se ScrollEvent) Vertical() bool {
	return se.ScrollDirection == ScrollUp || se.ScrollDirection == ScrollDown
}

func (se ScrollEvent) Horizontal() bool {
	return se.ScrollDirection == ScrollLeft || se.ScrollDirection == ScrollRight
}
