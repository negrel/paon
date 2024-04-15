package widgets

import (
	"testing"

	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/render"
	"github.com/stretchr/testify/require"
)

type testLayout struct {
	*BaseLayout
}

func TestBaseLayout(t *testing.T) {
	t.Run("Renderable/BaseLayoutRenderable", func(t *testing.T) {
		baseLayout := &testLayout{}
		renderable := NewBaseLayoutRenderable(baseLayout, LayoutChildrenFunc(
			func(co layout.Constraint, childrenPositions *ChildrenLayout) geometry.Size {
				return geometry.Size{}
			}),
		)
		baseLayout.BaseLayout = NewBaseLayout(NewPanicWidget(baseLayout), renderable)

		require.Equal(t, renderable, baseLayout.Renderable())
	})

	t.Run("Style/FromPanicWidget", func(t *testing.T) {
		baseLayout := &testLayout{}
		renderable := NewBaseLayoutRenderable(baseLayout, LayoutChildrenFunc(
			func(co layout.Constraint, childrenPositions *ChildrenLayout) geometry.Size {
				return geometry.Size{}
			}),
		)
		baseLayout.BaseLayout = NewBaseLayout(NewPanicWidget(baseLayout), renderable)

		require.Panics(t, func() { baseLayout.Style() })
	})
	t.Run("Style/StyledLayoutRenderable", func(t *testing.T) {
		baseLayout := &testLayout{}
		renderable := NewStyledLayoutRenderable(Style{}.Padding(1), NewBaseLayoutRenderable(baseLayout, LayoutChildrenFunc(
			func(co layout.Constraint, childrenPositions *ChildrenLayout) geometry.Size {
				return geometry.Size{}
			}),
		))
		baseLayout.BaseLayout = NewBaseLayout(NewPanicWidget(baseLayout), renderable)

		require.Equal(t, Style{}.Padding(1), baseLayout.Style())
	})
}

func TestBaseLayoutPropagateMouseEvents(t *testing.T) {
	surfaceSize := geometry.Size{Width: 100, Height: 100}

	// Number of mouse press events propagated to child.
	childWidgetMousePress := 0
	childLayoutMousePress := 0
	greatChildWidgetMousePress := 0
	greatChildWidget2MousePress := 0

	// A layout that place widgets diagonally (top right to bottom left).
	parent := &testLayout{}
	parent.BaseLayout = NewBaseLayout(
		NewPanicWidget(parent),
		NewBaseLayoutRenderable(
			parent,
			LayoutChildrenFunc(
				func(co layout.Constraint, childrenLayout *ChildrenLayout) geometry.Size {
					origin := geometry.Vec2D{X: 0, Y: 0}

					for child := parent.Node().FirstChild(); child != nil; child = child.Next() {
						childSize := child.Unwrap().(render.RenderableAccessor).Renderable().Layout(co)
						childrenLayout.layouts = append(childrenLayout.layouts,
							ChildLayout{
								Node: child,
								Bounds: geometry.Rectangle{
									Origin:   origin,
									RectSize: childSize,
								},
							},
						)

						origin = origin.Add(geometry.Vec2D{X: childSize.Width, Y: childSize.Height})
					}

					return co.MaxSize
				},
			),
		),
	)

	// A 10x10 widget.
	childWidget := NewComposedWidget(
		Style{},
		nil,
	)
	childWidget.RenderableAccessor = render.NewComposedRenderable(
		childWidget,
		layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.Size{Width: 10, Height: 10}
		}),
		draw.DrawerFunc(func(surface draw.Surface) {}),
	)

	childWidget.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		childWidgetMousePress++
	}))

	// A childLayout that place widgets diagonally (top right to bottom left).
	childLayout := &testLayout{}
	childLayout.BaseLayout = NewBaseLayout(
		NewPanicWidget(childLayout),
		NewBaseLayoutRenderable(
			childLayout,
			LayoutChildrenFunc(
				func(co layout.Constraint, childrenLayout *ChildrenLayout) geometry.Size {
					origin := geometry.Vec2D{X: 0, Y: 0}

					for child := childLayout.Node().FirstChild(); child != nil; child = child.Next() {
						childSize := child.Unwrap().(render.RenderableAccessor).Renderable().Layout(co)
						childrenLayout.layouts = append(childrenLayout.layouts,
							ChildLayout{
								Node: child,
								Bounds: geometry.Rectangle{
									Origin:   origin,
									RectSize: childSize,
								},
							},
						)

						origin = origin.Add(geometry.Vec2D{X: childSize.Width, Y: childSize.Height})
					}

					return geometry.Size{Width: origin.X, Height: origin.Y}
				},
			),
		),
	)
	childLayout.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		childLayoutMousePress++
	}))

	// A two 10x10 widget.
	greatChildWidget := NewComposedWidget(
		Style{},
		nil,
	)
	greatChildWidget.RenderableAccessor = render.NewComposedRenderable(
		greatChildWidget,
		layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.Size{Width: 10, Height: 10}
		}),
		draw.DrawerFunc(func(surface draw.Surface) {}),
	)
	greatChildWidget.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		greatChildWidgetMousePress++
	}))

	greatChildWidget2 := NewComposedWidget(
		Style{},
		nil,
	)
	greatChildWidget2.RenderableAccessor = render.NewComposedRenderable(
		greatChildWidget2,
		layout.LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.Size{Width: 10, Height: 10}
		}),
		draw.DrawerFunc(func(surface draw.Surface) {}),
	)
	greatChildWidget2.AddEventListener(mouse.PressListener(func(event mouse.Event) {
		greatChildWidget2MousePress++
	}))

	// Add widgets to tree.
	// parent __  childWidget
	//         \_ childLayout __  greatChildWidget
	//                         \_ greatChildWidget2
	err := parent.Node().AppendChild(childWidget.Node())
	require.NoError(t, err)

	err = parent.Node().AppendChild(childLayout.Node())
	require.NoError(t, err)

	err = childLayout.Node().AppendChild(greatChildWidget.Node())
	require.NoError(t, err)

	err = childLayout.Node().AppendChild(greatChildWidget2.Node())
	require.NoError(t, err)

	// Trigger a first layout so widgets are positionned.
	parent.Layout(layout.Constraint{
		MinSize:    geometry.Size{},
		MaxSize:    surfaceSize,
		ParentSize: surfaceSize,
		RootSize:   surfaceSize,
	})

	// Dispatch a mouse event that is out of child widget bounds.
	t.Run("OutOfChildBounds", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 50, Y: 50}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 0, childWidgetMousePress, "event propagated to child widget")
		require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
		require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
		require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")
	})

	// Dispatch a mouse event that is out of child widget X bounds.
	t.Run("InChildYBounds", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 50, Y: 5}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 0, childWidgetMousePress, "event propagated to child widget")
		require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
		require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
		require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")
	})

	// Dispatch a mouse event that is out of child widget Y bounds.
	t.Run("InChildXBounds", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 5, Y: 50}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 0, childWidgetMousePress, "event propagated to child widget")
		require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
		require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
		require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")
	})

	// Dispatch a mouse event that is within child widget bounds.
	t.Run("InChildBounds", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 5, Y: 5}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 1, childWidgetMousePress, "event not propagated to child widget")
		require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
		require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
		require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")
	})

	// Dispatch a mouse event that is within second great child widget bounds.
	t.Run("InSecondChildBounds", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 15, Y: 15}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 1, childWidgetMousePress, "event not propagated to child widget")
		require.Equal(t, 1, childLayoutMousePress, "event not propagated to child widget")
		require.Equal(t, 1, greatChildWidgetMousePress, "event propagated to great child widget")
		require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")
	})

	// Dispatch a mouse event that is within second great child widget bounds.
	t.Run("InSecondChildBounds", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 25, Y: 25}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 1, childWidgetMousePress, "event not propagated to child widget")
		require.Equal(t, 2, childLayoutMousePress, "event not propagated to child widget")
		require.Equal(t, 1, greatChildWidgetMousePress, "event propagated to great child widget")
		require.Equal(t, 1, greatChildWidget2MousePress, "event propagated to second great child widget")
	})

	// Remove child to ensure it doesn't reveive event anymore.
	for parent.Node().FirstChild() != nil {
		err = parent.Node().RemoveChild(parent.Node().FirstChild())
		require.NoError(t, err)
	}

	// Layout to update children bounding rect.
	parent.Layout(layout.Constraint{
		MinSize:    geometry.Size{},
		MaxSize:    surfaceSize,
		ParentSize: surfaceSize,
		RootSize:   surfaceSize,
	})

	// Dispatch a mouse event that is within child widget bounds.
	t.Run("UnmountedChildDoNotReceiveEvents", func(t *testing.T) {
		parent.DispatchEvent(mouse.NewPress(geometry.Vec2D{X: 5, Y: 5}, mouse.Button1, keypress.ModNone))
		require.Equal(t, 1, childWidgetMousePress, "event propagated to unmounted child widget")
		require.Equal(t, 2, childLayoutMousePress, "event propagated to unmounted child layout")
		require.Equal(t, 1, greatChildWidgetMousePress, "event propagated to unmounted great child widget")
		require.Equal(t, 1, greatChildWidget2MousePress, "event propagated to unmounted second great child widget")
	})
}
