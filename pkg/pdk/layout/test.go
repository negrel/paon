package layout

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pkg/pdk/styles"
	"github.com/negrel/paon/pkg/pdk/styles/property"
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

type mockStylised struct {
	style styles.Style
}

func makeMockStylised(style styles.Style) mockStylised {
	return mockStylised{
		style: style,
	}
}

func (ms mockStylised) Style() styles.Style {
	return ms.style
}

func applyBoxOffset(s styles.Style, offsets boxOffset, props [4]property.ID) {
	s.Set(property.MakeUnit(props[0], value.Unit{Value: offsets[0], UnitID: value.CellUnit}))
	s.Set(property.MakeUnit(props[1], value.Unit{Value: offsets[1], UnitID: value.CellUnit}))
	s.Set(property.MakeUnit(props[2], value.Unit{Value: offsets[2], UnitID: value.CellUnit}))
	s.Set(property.MakeUnit(props[3], value.Unit{Value: offsets[3], UnitID: value.CellUnit}))
}

func applyMargin(s styles.Style, margin boxOffset) {
	applyBoxOffset(s, margin, [4]property.ID{
		property.IDMarginLeft, property.IDMarginTop, property.IDMarginRight, property.IDMarginBottom,
	})
}

func applyBorder(s styles.Style, border boxOffset) {
	applyBoxOffset(s, border, [4]property.ID{
		property.IDBorderLeft, property.IDBorderTop, property.IDBorderRight, property.IDBorderBottom,
	})
}

func applyPadding(s styles.Style, padding boxOffset) {
	applyBoxOffset(s, padding, [4]property.ID{
		property.IDPaddingLeft, property.IDPaddingTop, property.IDPaddingRight, property.IDPaddingBottom,
	})
}

func makeConstraint(minWidth, minHeight, maxWidth, maxHeight int) Constraint {
	return Constraint{
		Min: geometry.Rect(0, 0, minWidth, minHeight),
		Max: geometry.Rect(0, 0, maxWidth, maxHeight),
	}
}
