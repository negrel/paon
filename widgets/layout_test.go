package widgets

import (
	"testing"

	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events/keypress"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/stretchr/testify/require"
)

func TestLayoutPropagateMouseEvents(t *testing.T) {
	surfaceSize := geometry.NewSize(100, 100)

	// Number of mouse press events propagated to child.
	childWidgetMousePress := 0
	childLayoutMousePress := 0
	greatChildWidgetMousePress := 0
	greatChildWidget2MousePress := 0

	// A layout that place widgets diagonally (top right to bottom left).
	var parent *BaseLayout
	parent = NewBaseLayout(
		func(co layout.Constraint, childrenPositions []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
			origin := geometry.NewVec2D(0, 0)

			for child := parent.Node().FirstChild(); child != nil; child = child.Next() {
				childSize := child.Unwrap().(layout.Layout).Layout(co)
				childrenPositions = append(childrenPositions,
					geometry.Rect(
						origin.X(),
						origin.Y(),
						origin.X()+childSize.Width(),
						origin.Y()+childSize.Height()),
				)

				origin = origin.Add(geometry.NewVec2D(childSize.Width(), childSize.Height()))
			}

			return childrenPositions, co.MaxSize
		},
	)

	// A 10x10 widget.
	childWidget := NewBaseWidget(
		LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.NewSize(10, 10)
		}),
		DrawerFunc(func(surface draw.Surface) {}),
		Listener(mouse.PressListener(func(event mouse.Event) {
			childWidgetMousePress++
		})),
	)

	// A childLayout that place widgets diagonally (top right to bottom left).
	var childLayout *BaseLayout
	childLayout = NewBaseLayout(
		func(co layout.Constraint, childrenPositions []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
			origin := geometry.NewVec2D(0, 0)

			for child := childLayout.Node().FirstChild(); child != nil; child = child.Next() {
				childSize := child.Unwrap().(layout.Layout).Layout(co)
				childrenPositions = append(childrenPositions,
					geometry.Rect(
						origin.X(),
						origin.Y(),
						origin.X()+childSize.Width(),
						origin.Y()+childSize.Height()),
				)

				origin = origin.Add(geometry.NewVec2D(childSize.Width(), childSize.Height()))
			}

			return childrenPositions, geometry.NewSize(origin.X(), origin.Y())
		},
		WidgetOptions(
			Listener(mouse.PressListener(func(event mouse.Event) {
				childLayoutMousePress++
			})),
		),
	)

	// A 10x10 widget.
	greatChildWidget := NewBaseWidget(
		LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.NewSize(10, 10)
		}),
		DrawerFunc(func(surface draw.Surface) {}),
		Listener(mouse.PressListener(func(event mouse.Event) {
			greatChildWidgetMousePress++
		})),
	)
	greatChildWidget2 := NewBaseWidget(
		LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.NewSize(10, 10)
		}),
		DrawerFunc(func(surface draw.Surface) {}),
		Listener(mouse.PressListener(func(event mouse.Event) {
			greatChildWidget2MousePress++
		})),
	)

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
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(50, 50), mouse.Button1, keypress.ModNone))
	require.Equal(t, 0, childWidgetMousePress, "event propagated to child widget")
	require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
	require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
	require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")

	// Dispatch a mouse event that is out of child widget X bounds.
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(50, 5), mouse.Button1, keypress.ModNone))
	require.Equal(t, 0, childWidgetMousePress, "event propagated to child widget")
	require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
	require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
	require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")

	// Dispatch a mouse event that is out of child widget Y bounds.
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(5, 50), mouse.Button1, keypress.ModNone))
	require.Equal(t, 0, childWidgetMousePress, "event propagated to child widget")
	require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
	require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
	require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")

	// Dispatch a mouse event that is within child widget bounds.
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(5, 5), mouse.Button1, keypress.ModNone))
	require.Equal(t, 1, childWidgetMousePress, "event not propagated to child widget")
	require.Equal(t, 0, childLayoutMousePress, "event propagated to child layout")
	require.Equal(t, 0, greatChildWidgetMousePress, "event propagated to great child widget")
	require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")

	// Dispatch a mouse event that is within second great child widget bounds.
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(15, 15), mouse.Button1, keypress.ModNone))
	require.Equal(t, 1, childWidgetMousePress, "event not propagated to child widget")
	require.Equal(t, 1, childLayoutMousePress, "event not propagated to child widget")
	require.Equal(t, 1, greatChildWidgetMousePress, "event propagated to great child widget")
	require.Equal(t, 0, greatChildWidget2MousePress, "event propagated to second great child widget")

	// Dispatch a mouse event that is within second great child widget bounds.
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(25, 25), mouse.Button1, keypress.ModNone))
	require.Equal(t, 1, childWidgetMousePress, "event not propagated to child widget")
	require.Equal(t, 2, childLayoutMousePress, "event not propagated to child widget")
	require.Equal(t, 1, greatChildWidgetMousePress, "event propagated to great child widget")
	require.Equal(t, 1, greatChildWidget2MousePress, "event propagated to second great child widget")

	// Remove child to ensure it doesn't reveive event anymore.
	for parent.Node().FirstChild() != nil {
		err = parent.Node().RemoveChild(parent.Node().FirstChild())
		require.NoError(t, err)
	}

	// No need to layout again.

	// Dispatch a mouse event that is within child widget bounds.
	parent.DispatchEvent(mouse.NewPress(geometry.NewVec2D(5, 5), mouse.Button1, keypress.ModNone))
	require.Equal(t, 1, childWidgetMousePress, "event propagated to unmounted child widget")
	require.Equal(t, 2, childLayoutMousePress, "event propagated to unmounted child layout")
	require.Equal(t, 1, greatChildWidgetMousePress, "event propagated to unmounted great child widget")
	require.Equal(t, 1, greatChildWidget2MousePress, "event propagated to unmounted second great child widget")
}
