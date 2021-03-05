package flows

import (
	"github.com/negrel/paon/internal/math"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
)

func GetUnitProp(style styles.Style, id property.ID) property.Unit {
	prop := style.Get(id)
	if prop != nil {
		if unitProp, isUnitProp := prop.(property.Unit); isUnitProp {
			return unitProp
		}
	}

	return property.Unit{}
}

func Constrain(value, min, max int) int {
	value = math.Max(value, min)
	value = math.Min(value, max)

	return value
}

func ComputeWidth(style styles.Style) int {
	width := GetUnitProp(style, property.IDWidth).Value.Value
	min := GetUnitProp(style, property.IDMinWidth).Value.Value
	max := GetUnitProp(style, property.IDMaxWidth).Value.Value

	return Constrain(width, min, max)
}

func ComputeConstrainedWidth(style styles.Style, constraint Constraint) int {
	return Constrain(ComputeWidth(style), constraint.Min.Width(), constraint.Max.Width())
}

func ComputeHeight(style styles.Style) int {
	height := GetUnitProp(style, property.IDHeight).Value.Value
	min := GetUnitProp(style, property.IDMinHeight).Value.Value
	max := GetUnitProp(style, property.IDMaxHeight).Value.Value

	return Constrain(height, min, max)
}

func ComputeConstrainedHeight(style styles.Style, constraint Constraint) int {
	return Constrain(ComputeHeight(style), constraint.Min.Height(), constraint.Max.Height())
}

func boxOf(style styles.Style, props [4]property.ID) boxOffset {
	left := GetUnitProp(style, props[0]).Value.Value
	top := GetUnitProp(style, props[1]).Value.Value
	right := GetUnitProp(style, props[2]).Value.Value
	bottom := GetUnitProp(style, props[3]).Value.Value

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
