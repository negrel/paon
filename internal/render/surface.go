package render

import (
	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/utils"
)

// Surface define terminal window/surface to draw on.
type Surface interface {
	Update()
	Apply(Patch)
	Size() utils.Size
	Width() int
	Height() int
	Clear()
	Fini()
	PollEvent() events.Event
}
