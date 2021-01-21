package renderer

import (
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

func computeThemeWidth(style styles.Style) int {
	width := -1

	if w := style.Get(property.IDWidth); w != nil {
		width = w.(property.Unit).Value.CellUnit().Value
	}

	if maxW := style.Get(property.IDMaxWidth); maxW != nil {
		maxWidth := maxW.(property.Unit).Value.CellUnit().Value
		width = math.Min(width, maxWidth)
	}

	if minW := style.Get(property.IDMinWidth); minW != nil {
		minWidth := minW.(property.Unit).Value.CellUnit().Value
		width = math.Max(width, minWidth)
	}

	return width
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
