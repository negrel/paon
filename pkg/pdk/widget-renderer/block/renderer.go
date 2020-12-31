package block

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/internal/render"
	"github.com/negrel/paon/pkg/styles"
)

func Layout(ctx render.Context) {
	width := computeLayerWidth(ctx.Object)
	height := computeLayerHeight(ctx.Object)

	ctx.Layer.Resize(geometry.NewSize(width, height))
}

func computeLayerWidth(object render.Object) int {
	width := computeThemeWidth(object.Theme())
	if width == -1 {
		width = computeWidgetWidth(object)
	}

	return width
}

func computeWidgetWidth(object render.Object) int {
	// Dynamic size base on widget content
	return 0
}

func computeThemeWidth(theme styles.Theme) int {
	width := -1

	if w := theme.GetUnitProp(styles.WidthID); w != nil {
		width = w.Value.CellUnit().Value
	}

	if maxW := theme.GetUnitProp(styles.MaxWidthID); maxW != nil {
		width = math.Min(width, maxW.Value.CellUnit().Value)
	}

	if minW := theme.GetUnitProp(styles.MinWidthID); minW != nil {
		width = math.Max(width, minW.Value.CellUnit().Value)
	}

	return width
}

func computeLayerHeight(object render.Object) int {
	width := computeThemeHeight(object.Theme())
	if width == -1 {
		width = computeWidgetHeight(object)
	}

	return width
}

func computeWidgetHeight(object render.Object) int {
	// Dynamic size base on widget content
	return 0
}

func computeThemeHeight(theme styles.Theme) int {
	height := 0

	if h := theme.GetUnitProp(styles.HeightID); h != nil {
		height = h.Value.CellUnit().Value
	}

	if maxH := theme.GetUnitProp(styles.MaxHeightID); maxH != nil {
		height = math.Min(height, maxH.Value.CellUnit().Value)
	}

	if minH := theme.GetUnitProp(styles.MinHeightID); minH != nil {
		height = math.Max(height, minH.Value.CellUnit().Value)
	}

	return height
}
