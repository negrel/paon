package layout

import (
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/styles/property"
)

func marginOf(style styles.Style) boxOffset {
	return newBoxOffset(
		style.IntUnit(property.MarginLeft()).Value(),
		style.IntUnit(property.MarginTop()).Value(),
		style.IntUnit(property.MarginRight()).Value(),
		style.IntUnit(property.MarginBottom()).Value(),
	)
}

func borderOf(style styles.Style) boxOffset {
	return newBoxOffset(
		style.IntUnit(property.BorderLeft()).Value(),
		style.IntUnit(property.BorderTop()).Value(),
		style.IntUnit(property.BorderRight()).Value(),
		style.IntUnit(property.BorderBottom()).Value(),
	)
}

func paddingOf(style styles.Style) boxOffset {
	return newBoxOffset(
		style.IntUnit(property.PaddingLeft()).Value(),
		style.IntUnit(property.PaddingTop()).Value(),
		style.IntUnit(property.PaddingRight()).Value(),
		style.IntUnit(property.PaddingBottom()).Value(),
	)
}
