package draw

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/geometry"
)

// Screen define a terminal window to apply render Canvas on.
type Screen interface {
	geometry.Sized

	Update()
	Apply(Canvas)
	Clear()
	Fini()
	PollEvent() events.Event
}
