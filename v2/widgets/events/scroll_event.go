package events

import (
	"time"
)

// ScrollListener define object that can listen to scroll events.
type ScrollListener interface {
	OnScroll(*ScrollEvent)
}

// Direction define the scroll direction.
type Direction int

// String implements fmt.Stringer interface.
func (d Direction) String() string {
	switch d {
	case ScrollUp:
		return "up"
	case ScrollDown:
		return "down"
	case ScrollRight:
		return "right"
	case ScrollLeft:
		return "left"
	case ScrollUpLeft:
		return "up-left"
	case ScrollUpRight:
		return "up-right"
	case ScrollDownLeft:
		return "down-left"
	case ScrollDownRight:
		return "down-right"

	default:
		return ""
	}
}

// Possible scroll direction.
const (
	ScrollUp Direction = iota
	ScrollDown
	ScrollRight
	ScrollLeft
	ScrollUpLeft
	ScrollUpRight
	ScrollDownLeft
	ScrollDownRight
)

// ScrollEvent are triggered when the user scroll
// in the terminal window.
type ScrollEvent struct {
	*event

	dir Direction
}

// NewScrollEvent return a new ScrollEvent instance.
func NewScrollEvent(timeStamp time.Time, dir Direction) *ScrollEvent {
	return &ScrollEvent{
		event: &event{
			evType:    ScrollEventType,
			timeStamp: timeStamp,
		},
		dir: dir,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Direction return the scroll direction.
func (se *ScrollEvent) Direction() Direction {
	return se.dir
}
