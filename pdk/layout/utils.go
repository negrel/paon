package layout

import (
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

func getUnitProp(style styles.Style, id property.ID) (property.Unit, bool) {
	prop := style.Get(id)
	if prop != nil {
		if unitProp, isUnitProp := prop.(property.Unit); isUnitProp {
			return unitProp, true
		}
	}

	return property.Unit{}, false
}

func computeMinMaxWidth(style styles.Style, constraint Constraint) (int, int) {
	minWidth := constraint.Min.Width()
	maxWidth := constraint.Max.Width()

	minWidthProp, ok := getUnitProp(style, property.MinWidth())
	if ok {
		minWidth = math.Max(minWidth, constraint.ToCellUnit(minWidthProp.Value))
	}
	maxWidthProp, ok := getUnitProp(style, property.MaxWidth())
	if ok {
		maxWidth = math.Min(maxWidth, constraint.ToCellUnit(maxWidthProp.Value))
	}

	return minWidth, maxWidth
}

func computeMinMaxHeight(style styles.Style, constraint Constraint) (int, int) {
	minHeight := constraint.Min.Height()
	maxHeight := constraint.Max.Height()

	minHeightProp, ok := getUnitProp(style, property.MinHeight())
	if ok {
		minHeight = math.Max(minHeight, constraint.ToCellUnit(minHeightProp.Value))
	}
	maxHeightProp, ok := getUnitProp(style, property.MaxHeight())
	if ok {
		maxHeight = math.Min(maxHeight, constraint.ToCellUnit(maxHeightProp.Value))
	}

	return minHeight, maxHeight
}

func boxOf(style styles.Style, props [4]property.ID) boxOffset {

	left := 0
	leftProp, ok := getUnitProp(style, props[0])
	if ok {
		left = leftProp.Value.Value
	}

	top := 0
	topProp, ok := getUnitProp(style, props[1])
	if ok {
		top = topProp.Value.Value
	}

	right := 0
	rightProp, ok := getUnitProp(style, props[2])
	if ok {
		right = rightProp.Value.Value
	}

	bottom := 0
	bottomProp, ok := getUnitProp(style, props[3])
	if ok {
		bottom = bottomProp.Value.Value
	}

	return makeBoxOffset(left, top, right, bottom)
}

func marginOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.MarginLeft(),
		property.MarginTop(),
		property.MarginRight(),
		property.MarginBottom(),
	})
}

func borderOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.BorderLeft(),
		property.BorderTop(),
		property.BorderRight(),
		property.BorderBottom(),
	})
}

func paddingOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.PaddingLeft(),
		property.PaddingTop(),
		property.PaddingRight(),
		property.PaddingBottom(),
	})
}
