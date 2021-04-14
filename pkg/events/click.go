package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
)

var _TypeClick = events.MakeType("click")

func TypeClick() events.Type {
	return _TypeClick
}

type ClickType int16

const (
	Click1 ClickType = 1 << iota // Usually the left (primary) mouse button.
	Click2                       // Usually the right (secondary) mouse button.
	Click3                       // Usually the middle mouse button.
	Click4                       // Often a side button (thumb/next).
	Click5                       // Often a side button (thumb/prev).
	Click6
	Click7
	Click8
	ClickPrimary   = Click1
	ClickSecondary = Click2
	ClickMiddle    = Click3
)

// ClickListener convert the given event handler as a generic Listener.
func ClickListener(handler func(Click)) *events.Listener {
	l := events.Listener{
		Type: _TypeClick,
		Handle: func(event events.Event) {
			assert.IsType(event, MakeClick(geometry.Point{}, 0))
			handler(event.(Click))
		},
	}

	return &l
}

var _ events.Event = Click{}

// Click is triggered when the user click inside the rendering surface.
type Click struct {
	events.Event
	Position  geometry.Point
	ClickType ClickType
}

// MakeClick returns a new Click events.Event.
func MakeClick(position geometry.Point, clickType ClickType) Click {
	return Click{
		Event:     events.MakeEvent(_TypeClick),
		Position:  position,
		ClickType: clickType,
	}
}

func (c Click) Is(clickType ClickType) bool {
	return (c.ClickType & clickType) != 0
}
