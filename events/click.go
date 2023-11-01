package events

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/geometry"
	pdkevents "github.com/negrel/paon/pdk/events"
)

var clickType = pdkevents.NewType("Click")

// ClickType returns the type for Click events.
func ClickType() pdkevents.Type {
	return clickType
}

// ClickListener returns an events.Listener that will call the given handler
// on Click events.
func ClickListener(handler func(Click)) (pdkevents.Type, pdkevents.Handler) {
	return ClickType(), pdkevents.HandlerFunc(func(event pdkevents.Event) {
		assert.IsType(event, Click{})
		handler(event.(Click))
	})
}

// Click define an event that is triggered when the terminal is Clickd.
type Click struct {
	pdkevents.Event
	Position geometry.Vec2D
}

// NewClick returns a new Click event.
func NewClick(pos geometry.Vec2D) Click {
	return Click{
		Event:    pdkevents.NewEvent(clickType),
		Position: pos,
	}
}

// String implements the fmt.Stringer interface.
func (r Click) String() string {
	return fmt.Sprintf("%s{Position: %+v}", r.Event.Type(), r.Position)
}
