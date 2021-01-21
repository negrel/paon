package events

import (
	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/geometry"
)

var _ Event = Resize{}

// ResizeListener convert the given event handler as a generic Listener.
func ResizeListener(handler func(Resize)) *Listener {
	l := Listener{
		Type: TypeMouseMove,
		Handle: func(event Event) {
			re, ok := event.(Resize)

			if !ok {
				log.Warnf("\"%v\" listener expected, but got %v", TypeMouseMove, event.Type())
				return
			}

			handler(re)
		},
	}

	return &l
}

// Resize is triggered when the user resize rendering surface.
type Resize struct {
	Event
	geometry.Size

	IsWider, IsGreater bool
}

// MakeResize return a new Resize events.Event.
func MakeResize(newSize, oldSize geometry.Size) Resize {
	return Resize{
		Event:     MakeEvent(TypeResize),
		Size:      newSize,
		IsWider:   newSize.Width() > oldSize.Width(),
		IsGreater: newSize.Height() > oldSize.Height(),
	}
}
