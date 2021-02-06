package renderer

import "github.com/negrel/paon/internal/render"

var _ render.Renderer = flex{}

type flex struct{}

func makeFlex() render.Renderer {
	return flex{}
}

func (f flex) Layout(ctx render.Context) {}

func (f flex) Draw(ctx render.Context) {}
