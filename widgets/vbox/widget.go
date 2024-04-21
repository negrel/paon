package vbox

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/minmax"
	"github.com/negrel/paon/render"
	"github.com/negrel/paon/widgets"
)

type Option func(*Widget)

// WithChildren return an option that appends the given children.
func WithChildren(children ...widgets.Widget) Option {
	return func(w *Widget) {
		for _, child := range children {
			err := w.Node().AppendChild(child.Node())
			if err != nil {
				panic(err)
			}
		}
	}
}

// WithStyle return an option that sets button widget style.
func WithStyle(style widgets.Style) Option {
	return func(w *Widget) {
		w.style.InnerStyle = style
	}
}

type Widget struct {
	*widgets.BaseLayout
	style widgets.InheritStyle
}

func New(options ...Option) *Widget {
	w := &Widget{}
	w.style = widgets.InheritStyle{
		NodeAccessor: w,
		InnerStyle:   widgets.Style{},
	}

	w.BaseLayout = widgets.NewBaseLayout(
		widgets.NewPanicWidget(w),
		widgets.NewStyledLayoutRenderable(
			&w.style,
			widgets.NewBaseLayoutRenderable(
				w,
				widgets.LayoutChildrenFunc(
					func(co layout.Constraint, childrenLayout *widgets.ChildrenLayout) geometry.Size {
						size := geometry.Size{}
						freeSpace := co.MaxSize

						for child := w.Node().FirstChild(); child != nil; child = child.Next() {
							// Previous child fulfilled the surface, no need to render next siblings.
							if freeSpace.Width <= 0 {
								break
							}

							// Compute child size.
							childLayout := child.Unwrap().(render.RenderableAccessor).Renderable()

							childSize := childLayout.Layout(layout.Constraint{
								MinSize:    geometry.Size{},
								MaxSize:    freeSpace,
								ParentSize: co.ParentSize,
								RootSize:   co.RootSize,
							})

							// Store child rectangle.
							childrenLayout.Append(widgets.ChildLayout{
								Node: child,
								Bounds: geometry.Rectangle{
									Origin:   geometry.Vec2D{Y: size.Height},
									RectSize: childSize,
								},
							})

							// Update freespace.
							freeSpace.Height -= childSize.Height

							// Update VBox size.
							size = geometry.Size{
								Width:  minmax.Max(size.Width, childSize.Width),
								Height: size.Height + childSize.Height,
							}
						}

						return size
					},
				),
			),
		),
	)

	for _, applyOption := range options {
		applyOption(w)
	}

	return w
}
