package renderer

import "github.com/negrel/paon/internal/render"

var _ render.Renderer = hidden{}

type hidden struct{}

func makeHidden() render.Renderer {
	return hidden{}
}

func (h hidden) Layout(ctx render.Context) {}

func (h hidden) Draw(ctx render.Context) {}
