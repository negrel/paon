package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/math"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

type HBox struct {
	*pdkwidgets.BaseLayout
}

func NewHBox(children ...pdkwidgets.Widget) *HBox {
	w := &HBox{}

	w.BaseLayout = pdkwidgets.NewBaseLayout(
		pdkwidgets.WidgetOptions(
			pdkwidgets.Wrap(w),
		),
	)

	for _, child := range children {
		w.AppendChild(child)
	}

	return w
}

// Render implements the Widget interface.
func (hb *HBox) Render(co layout.Constraint, surface draw.Surface) geometry.Size {
	size := geometry.NewSize(0, 0)
	surfaceSize := surface.Size()

	for child := hb.FirstChild(); child != nil; child = child.NextSibling() {
		// Previous child fulfilled the surface, no need to render next siblings.
		if surfaceSize.Width() < size.Width() {
			break
		}

		// Reduce subsurface to remaining space.
		subsurface := draw.NewSubSurface(surface, geometry.Rect(size.Width(), 0, surfaceSize.Width(), surfaceSize.Height()))

		childSize := child.Render(co, subsurface)
		size = geometry.NewSize(
			size.Width()+childSize.Width(),
			math.Max(size.Height(), childSize.Height()),
		)
	}

	return size
}
