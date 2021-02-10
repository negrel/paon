package render

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/events"
)

// Screen define a terminal window to apply render Buffer on.
type Screen interface {
	geometry.Sized

	Update()
	Apply(Buffer)
	Clear()
	Fini()
	PollEvent(chan<- events.Event)
}
