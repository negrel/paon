package renderer

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/widgets"
)

var _ render.Renderer = block{}

type block struct{}

func makeBlock() render.Renderer {
	return block{}
}

func (b block) Layout(ctx *render.Context) {
	width := computeLayerWidth(ctx.Object)
	height := computeLayerHeight(ctx.Object)

	ctx.Layer.Resize(geometry.NewSize(width, height))
}

func computeLayerWidth(object render.Object) int {
	width := computeThemeWidth(object.Style())
	if width == -1 {
		width = 0
	}

	return width
}

func computeLayerHeight(object render.Object) int {
	height := computeThemeHeight(object.Style())
	if height == -1 {
		height = computeWidgetHeight(object.(widgets.Widget))
	}

	return height
}

func computeWidgetHeight(object render.Object) int {
	// Dynamic size base on widget content
	return 0
}

func (b block) Draw(ctx *render.Context) {
	panic("implement me")
}
