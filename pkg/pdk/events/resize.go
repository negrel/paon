package events

import (
	"github.com/negrel/paon/internal/geometry"
)

var _ Event = Resize{}

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
