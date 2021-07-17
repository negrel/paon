package events

import (
	"fmt"

	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/events"
	pdkevents "github.com/negrel/paon/pdk/events"
)

var resizeType = pdkevents.NewType("resize")

// ResizeType returns the type for Resize events.
func ResizeType() pdkevents.Type {
	return resizeType
}

// Resize define an event that is triggered when the terminal is resized.
type Resize struct {
	events.Event
	Old, New geometry.Size
}

// NewResize returns a new Resize event.
func NewResize(old, new geometry.Size) Resize {
	return Resize{
		Event: pdkevents.NewEvent(resizeType),
		Old:   old,
		New:   new,
	}
}

// String implements the fmt.Stringer interface.
func (r Resize) String() string {
	return fmt.Sprintf("%s{Old: %v, New: %v}", r.Event.Type(), r.Old, r.New)
}

// IsWider returns true if the new terminal size is wider.
func (r Resize) IsWider() bool {
	return r.Old.Size().Width() < r.New.Size().Width()
}

// IsGreater returns true if the new terminal size is greater.
func (r Resize) IsGreater() bool {
	return r.Old.Size().Height() < r.New.Size().Height()
}
