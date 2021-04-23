package events

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
)

var _TypeResize = events.MakeType("resize")

func TypeResize() events.Type {
	return _TypeResize
}

var _ events.Event = Resize{}

// ResizeListener convert the given event handler as a generic Listener.
func ResizeListener(handler func(Resize)) *events.Listener {
	l := events.Listener{
		Type: _TypeResize,
		Handle: func(event events.Event) {
			assert.IsType(event, MakeResize(geometry.Size{}, geometry.Size{}))
			handler(event.(Resize))
		},
	}

	return &l
}

// Resize is triggered when the user resize rendering surface.
type Resize struct {
	events.Event
	geometry.Size
	OldSize geometry.Size

	IsWider, IsGreater bool
}

// MakeResize return a new Resize events.Event.
func MakeResize(newSize, oldSize geometry.Size) Resize {
	return Resize{
		Event:     events.MakeEvent(_TypeResize),
		Size:      newSize,
		OldSize:   oldSize,
		IsWider:   newSize.Width() > oldSize.Width(),
		IsGreater: newSize.Height() > oldSize.Height(),
	}
}
