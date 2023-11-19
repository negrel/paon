package vbox

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

type Widget struct {
	widgets.StyledLayout
}

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
