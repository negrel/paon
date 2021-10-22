package render

import (
	"math"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
)

const maxLayer = math.MaxUint8

type compositeCell struct {
	position geometry.Vec3D
	cell     draw.Cell
}

// Compositor define a rendering compositor
type Compositor struct {
	dst       draw.Surface
	rootLayer Layer
	done      chan struct{}
}

// NewCompositor returns a new Compositor for the given surface.
func NewCompositor(destination draw.Surface) *Compositor {
	c := &Compositor{
		dst: destination,
	}

	return c
}

// Compose starts an infinite compositing loop.
func (c *Compositor) Compose(tick <-chan struct{}) {
	c.done = make(chan struct{})

	for {
		select {
		case <-c.done:
			return

		case <-tick:
			size := c.rootLayer.Size()
			for x := 0; x < size.Width(); x++ {
				for y := 0; y < size.Height(); y++ {
					v := geometry.NewVec2D(x, y)
					c.dst.Set(v, c.rootLayer.Get(v))
				}
			}
		}
	}
}

// Stop stops the compoiting loop.
func (c *Compositor) Stop() {
	if c.done == nil {
		return
	}

	c.done <- struct{}{}
	close(c.done)
	c.done = nil
	log.Debug("Compositor stopped")
}
