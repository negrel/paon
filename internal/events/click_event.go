package events

import (
	"github.com/negrel/paon/internal/utils"
)

type ClickListener interface {
	OnClick(ClickEvent)
}

var _ Event = ClickEvent{}

// ClickEvent is triggered when the user click inside the rendering surface.
type ClickEvent struct {
	event
	Position utils.Point
}

// MakeResizeEvent returns a new ClickEvent object.
func MakeClickEvent(position utils.Point) ClickEvent {
	return ClickEvent{
		event:    makeEvent(ClickEventType),
		Position: position,
	}
}
