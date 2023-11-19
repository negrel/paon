package hbox

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/math"
	"github.com/negrel/paon/widgets"
)

type Option func(*Widget)

// WithChildren return an option that appends the given children.
func WithChildren(children ...widgets.Widget) Option {
	return func(w *Widget) {
		for _, child := range children {
			w.Node().AppendChild(child.Node())
		}
	}
}

// WithStyle return an option that sets button widget style.
func WithStyle(style widgets.Style) Option {
	return func(w *Widget) {
		w.StyledLayout.SetStyle(widgets.InheritStyle{
			Widget:     w,
			InnerStyle: style,
		})
	}
}

// Widget define a widget that displays its children in an horizontal array.
// To cause a child to expand to fill the available horizontal space, wrap the
// child in an Expanded widget.
type Widget struct {
	widgets.StyledLayout
}

// New returns a new hbox widget configured with the given options.
func New(options ...Option) *Widget {
	w := &Widget{}

	w.StyledLayout = widgets.NewStyledLayout(
		widgets.NewBaseWidget(w),
		nil,
		func(co layout.Constraint, childrenRects []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
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
		},
	)

	for _, applyOption := range options {
		applyOption(w)
	}

	if w.StyledLayout.Style() == nil {
		w.StyledLayout.SetStyle(widgets.InheritStyle{
			Widget:     w,
			InnerStyle: widgets.Style{},
		})
	}

	return w
}
