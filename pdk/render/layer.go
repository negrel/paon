package render

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/tree"
)

type surface = draw.Surface

var _ draw.Surface = &Layer{}

// Layer define a node in a tree of drawing surface used for the rendering.
type Layer struct {
	tree.Node
	surface
}

func newLayer(surface draw.Surface, leaf bool) Layer {
	layer := Layer{
		surface: surface,
	}

	if leaf {
		layer.Node = tree.NewLeafNode(layer)
	} else {
		layer.Node = tree.NewNode(layer)
	}

	return layer
}

// Set implements the draw.Surface interface.
func (l Layer) Set(pos geometry.Vec2D, cell draw.Cell) {
	l.surface.Set(pos, cell)
}
