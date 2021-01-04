package renderer

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
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

func computeThemeWidth(theme styles.Style) int {
	width := -1

	if w := theme.Get(property.IDWidth); w != nil {
		width = w.(property.Unit).Value.CellUnit().Value
	}

	if maxW := theme.Get(property.IDMaxWidth); maxW != nil {
		maxWidth := maxW.(property.Unit).Value.CellUnit().Value
		width = math.Min(width, maxWidth)
	}

	if minW := theme.Get(property.IDMinWidth); minW != nil {
		minWidth := minW.(property.Unit).Value.CellUnit().Value
		width = math.Max(width, minWidth)
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

func computeThemeHeight(theme styles.Style) int {
	height := -1

	if w := theme.Get(property.IDHeight); w != nil {
		height = w.(property.Unit).Value.CellUnit().Value
	}

	if maxW := theme.Get(property.IDMaxHeight); maxW != nil {
		maxHeight := maxW.(property.Unit).Value.CellUnit().Value
		height = math.Min(height, maxHeight)
	}

	if minW := theme.Get(property.IDMinHeight); minW != nil {
		minHeight := minW.(property.Unit).Value.CellUnit().Value
		height = math.Max(height, minHeight)
	}

	return height
}

func (b block) Draw(ctx *render.Context) {
	panic("implement me")
}
