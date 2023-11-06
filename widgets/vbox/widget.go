package vbox

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/math"
	"github.com/negrel/paon/widgets"
)

type Widget struct {
	*widgets.BaseLayout
}

func New(children ...widgets.Widget) *Widget {
	w := &Widget{}

	w.BaseLayout = widgets.NewBaseLayout(
		widgets.WidgetOptions(
			widgets.Wrap(w),
		),
		widgets.LayoutAlgo(func(co layout.Constraint, childrenRects []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
			size := geometry.NewSize(0, 0)
			freeSpace := co.MaxSize

			for child := w.Node().FirstChild(); child != nil; child = child.Next() {
				// Previous child fulfilled the surface, no need to render next siblings.
				if freeSpace.Width() <= 0 {
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
				childrenRects = append(childrenRects, geometry.Rect(0, size.Height(), childSize.Width(), size.Height()+childSize.Height()))

				// Update freespace.
				freeSpace = geometry.NewSize(freeSpace.Width(), freeSpace.Height()-childSize.Height())

				// Update VBox size.
				size = geometry.NewSize(
					math.Max(size.Width(), childSize.Width()),
					size.Height()+childSize.Height(),
				)
			}

			return childrenRects, size
		}),
	)

	for _, child := range children {
		w.Node().AppendChild(child.Node())
	}

	return w
}
