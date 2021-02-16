package layout

import (
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

func unitProp(style styles.Style, id property.ID) property.Unit {
	prop := style.Get(id)
	if prop != nil {
		if unitProp, isUnitProp := prop.(property.Unit); isUnitProp {
			return unitProp
		}
	}

	return property.Unit{}
}

func constrain(value, min, max int) int {
	value = math.Max(value, min)
	value = math.Min(value, max)

	return value
}

func computeWidth(style styles.Style) int {
	width := unitProp(style, property.IDWidth).Value.Value
	min := unitProp(style, property.IDMinWidth).Value.Value
	max := unitProp(style, property.IDMaxWidth).Value.Value

	return constrain(width, min, max)
}

func computeConstrainedWidth(style styles.Style, constraint Constraint) int {
	return constrain(computeWidth(style), constraint.Min.Width(), constraint.Max.Width())
}

func computeHeight(style styles.Style) int {
	height := unitProp(style, property.IDHeight).Value.Value
	min := unitProp(style, property.IDMinHeight).Value.Value
	max := unitProp(style, property.IDMaxHeight).Value.Value

	return constrain(height, min, max)
}

func computeConstrainedHeight(style styles.Style, constraint Constraint) int {
	return constrain(computeHeight(style), constraint.Min.Height(), constraint.Max.Height())
}

func boxOf(style styles.Style, props [4]property.ID) boxOffset {
	left := unitProp(style, props[0]).Value.Value
	top := unitProp(style, props[1]).Value.Value
	right := unitProp(style, props[2]).Value.Value
	bottom := unitProp(style, props[3]).Value.Value

	return makeBoxOffset(left, top, right, bottom)
}

func marginOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.IDMarginLeft,
		property.IDMarginTop,
		property.IDMarginRight,
		property.IDMarginBottom,
	})
}

func borderOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.IDBorderLeft,
		property.IDBorderTop,
		property.IDBorderRight,
		property.IDBorderBottom,
	})
}

func paddingOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.IDPaddingLeft,
		property.IDPaddingTop,
		property.IDPaddingRight,
		property.IDPaddingBottom,
	})
}
