package events

import (
	"github.com/negrel/paon/internal/geometry"
)

var _ Event = ResizeEvent{}

// ResizeEvent is triggered when the user resize rendering surface.
type ResizeEvent struct {
	event
	geometry.Size

	IsWider, IsGreater bool
}

// MakeResizeEvent return a new ResizeEvent object.
func MakeResizeEvent(newSize, oldSize geometry.Size) ResizeEvent {
	return ResizeEvent{
		event:     makeEvent(ResizeEventType),
		Size:      newSize,
		IsWider:   newSize.Width() > oldSize.Width(),
		IsGreater: newSize.Height() > oldSize.Height(),
	}
}
