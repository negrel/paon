package hbox

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/math"
	"github.com/negrel/paon/widgets"
)

// Widget define a widget that displays its children in an horizontal array.
// To cause a child to expand to fill the available horizontal space, wrap the
// child in an Expanded widget.
type Widget struct {
	*widgets.BaseLayout
}

// New returns a new hbox widget configured with the given options.
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
				if freeSpace.Height() <= 0 {
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

			return childrenRects, size
		}),
	)

	for _, child := range children {
		w.Node().AppendChild(child.Node())
	}

	return w
}
