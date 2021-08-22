package layout

import (
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

func UnitProp(style styles.Style, id property.ID) (property.Unit, bool) {
	prop := style.Get(id)
	if prop != nil {
		if unitProp, isUnitProp := prop.(property.Unit); isUnitProp {
			return unitProp, true
		}
	}

	return property.Unit{}, false
}

func boxOf(style styles.Style, props [4]property.ID) boxOffset {

	left := 0
	leftProp, ok := UnitProp(style, props[0])
	if ok {
		left = leftProp.Value.Value
	}

	top := 0
	topProp, ok := UnitProp(style, props[1])
	if ok {
		top = topProp.Value.Value
	}

	right := 0
	rightProp, ok := UnitProp(style, props[2])
	if ok {
		right = rightProp.Value.Value
	}

	bottom := 0
	bottomProp, ok := UnitProp(style, props[3])
	if ok {
		bottom = bottomProp.Value.Value
	}

	return newBoxOffset(left, top, right, bottom)
}

func marginOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.MarginLeftID(),
		property.MarginTopID(),
		property.MarginRightID(),
		property.MarginBottomID(),
	})
}

func borderOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.BorderLeftID(),
		property.BorderTopID(),
		property.BorderRightID(),
		property.BorderBottomID(),
	})
}

func paddingOf(style styles.Style) boxOffset {
	return boxOf(style, [4]property.ID{
		property.PaddingLeftID(),
		property.PaddingTopID(),
		property.PaddingRightID(),
		property.PaddingBottomID(),
	})
}
