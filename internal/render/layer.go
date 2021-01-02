package render

import (
	"github.com/negrel/paon/internal/draw"
	"github.com/negrel/paon/internal/geometry"
)

type Layer struct {
	*draw.Canvas
}

func NewLayer() *Layer {
	return &Layer{
		Canvas: draw.NewCanvas(geometry.Rectangle{}),
	}
}
