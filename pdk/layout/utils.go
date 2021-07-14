package layout

import (
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

// UnitProp converts the property contained in the given styles.Style that
// have the given property.ID and returns it.
// A true boolean is also returned if the property is a property.Unit.
func UnitProp(style styles.Style, id property.ID) (property.Unit, bool) {
	prop := style.Get(id)
	if prop != nil {
		if unitProp, isUnitProp := prop.(property.Unit); isUnitProp {
			return unitProp, true
		}
	}

	return property.Unit{}, false
}

// MinMaxWidth returns min and max width that respect the Constraint and
// the styles.Styles min/max width properties.
func MinMaxWidth(style styles.Style, constraint Constraint) (int, int) {
	minWidth := constraint.Min.Width()
	maxWidth := constraint.Max.Width()

	minWidthProp, ok := UnitProp(style, property.MinWidthID())
	if ok {
		minWidth = math.Max(minWidth, constraint.ToCellUnit(minWidthProp.Value))
	}
	maxWidthProp, ok := UnitProp(style, property.MaxWidthID())
	if ok {
		maxWidth = math.Min(maxWidth, constraint.ToCellUnit(maxWidthProp.Value))
	}

	return minWidth, maxWidth
}

// MinMaxHeight returns min and max height that respect the given Constraint and
// the styles.Styles min/max height properties.
func MinMaxHeight(style styles.Style, constraint Constraint) (int, int) {
	minHeight := constraint.Min.Height()
	maxHeight := constraint.Max.Height()

	minHeightProp, ok := UnitProp(style, property.MinHeightID())
	if ok {
		minHeight = math.Max(minHeight, constraint.ToCellUnit(minHeightProp.Value))
	}
	maxHeightProp, ok := UnitProp(style, property.MaxHeightID())
	if ok {
		maxHeight = math.Min(maxHeight, constraint.ToCellUnit(maxHeightProp.Value))
	}

	return minHeight, maxHeight
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
