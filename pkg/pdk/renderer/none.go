package renderer

import "github.com/negrel/paon/internal/render"

var _ render.Renderer = none{}

type none struct{}

func makeHidden() render.Renderer {
	return none{}
}

func (n none) Layout(ctx render.Context) {}

func (n none) Draw(ctx render.Context) {}
