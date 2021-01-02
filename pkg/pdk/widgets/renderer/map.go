package renderer

import (
	"github.com/negrel/paon/internal/render"
)

var Map = map[int]func() render.Renderer{}
