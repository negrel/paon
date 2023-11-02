package widgets

import (
	"github.com/negrel/paon/geometry"
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
		pdkwidgets.LayoutAlgo(func(co layout.Constraint, childrenRects []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
			size := geometry.NewSize(0, 0)
			freeSpace := co.MaxSize

			for child := w.FirstChild(); child != nil; child = child.Next() {
				// Previous child fulfilled the surface, no need to render next siblings.
				if freeSpace.Height() <= size.Height() {
					break
				}

				// Compute child size.
				childLayout := child.Unwrap().(layout.Layout)
				childSize := childLayout.Layout(layout.Constraint{
					MinSize:    geometry.NewSize(0, 0),
					MaxSize:    freeSpace,
					ParentSize: co.ParentSize,
					RootSize:   co.RootSize,
				})

				// Store child rectangle.
				childrenRects = append(childrenRects, geometry.Rect(size.Width(), 0, size.Width()+childSize.Width(), childSize.Height()))

				// Update freespace.
				freeSpace = geometry.NewSize(freeSpace.Width()-childSize.Width(), freeSpace.Height())

				// Update HBox size.
				size = geometry.NewSize(
					size.Width()+childSize.Width(),
					math.Max(size.Height(), childSize.Height()),
				)
			}

			return childrenRects, co.ApplyOnSize(size)
		}),
	)

	for _, child := range children {
		w.AppendChild(child.Node())
	}

	return w
}
