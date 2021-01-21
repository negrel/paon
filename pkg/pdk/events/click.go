package events

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
)

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
func ClickListener(handler func(Click)) *Listener {
	l := Listener{
		Type: TypeClick,
		Handle: func(event Event) {
			ce, ok := event.(Click)

			if !ok {
				log.Warnf("\"%v\" listener expected, but got %v", TypeClick, event.Type())
				return
			}

			handler(ce)
		},
	}

	return &l
}

var _ Event = Click{}

// Click is triggered when the user click inside the rendering surface.
type Click struct {
	Event
	Position  geometry.Point
	ClickType ClickType
}

// MakeClick returns a new Click events.Event.
func MakeClick(position geometry.Point, clickType ClickType) Click {
	return Click{
		Event:     MakeEvent(TypeClick),
		Position:  position,
		ClickType: clickType,
	}
}

func (c Click) Is(clickType ClickType) bool {
	return (c.ClickType & clickType) != 0
}
