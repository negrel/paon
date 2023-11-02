package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/layout"
)

type baseLayoutOption struct {
	*BaseLayout

	widgetConstructor func(...WidgetOption) *BaseWidget
	widgetOptions     []WidgetOption
}

// LayoutOption define an option for BaseLayout.
type LayoutOption func(*baseLayoutOption)

// WidgetOptions returns a LayoutOption that will define the options to use
// by the internal Widget object.
func WidgetOptions(options ...WidgetOption) LayoutOption {
	return func(blo *baseLayoutOption) {
		blo.widgetOptions = append(blo.widgetOptions, options...)
	}
}

func widgetConstructor(constructor func(...WidgetOption) *BaseWidget) LayoutOption {
	return func(blo *baseLayoutOption) {
		blo.widgetConstructor = constructor
	}
}

// LayoutAlgo returns a LayoutOptions that define the layout algorithm.
func LayoutAlgo(algo func(co layout.Constraint, childrenPositions []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size)) LayoutOption {
	return func(blo *baseLayoutOption) {
		blo.layoutAlgo = algo
	}
}
