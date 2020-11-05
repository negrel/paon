package render

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/geometry"
)

// Screen define a terminal window to apply render Patch on.
type Screen interface {
	geometry.Sized

	Update()
	Apply(Patch)
	Clear()
	Fini()
	PollEvent() events.Event
}
