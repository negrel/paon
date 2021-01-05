package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
)

// Window define a terminal window to apply render Canvas on.
type Window interface {
	geometry.Sized

	Update()
	Apply(Canvas)
	Clear()
	Fini()
	PollEvent() events.Event
}
