package render

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

var _ draw.Surface = &Layer{}

// Layer define a node in a tree of drawing surface.
type Layer struct {
	treevents.Node
	BufferSurface
}

// NewLayer returns a new rendering layer that wraps the given drawing surface.
func NewLayer(node treevents.Node) Layer {
	assert.NotNil(node)

	layer := Layer{
		Node:          node,
		BufferSurface: NewBufferSurface(geometry.Size{}),
	}

	return layer
}

// AddLayer adds a new child Layer.
func (l Layer) AddLayer(layer Layer) error {
	return l.Node.AppendChild(layer.Node)
}

// SubSurface returns a draw.SubSurface of the given bounds. The returned
// subsurface wraps the layer surface.
func (l Layer) SubSurface(bounds geometry.Rectangle) draw.SubSurface {
	return draw.NewSubSurface(l.BufferSurface, bounds)
}
