package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
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
