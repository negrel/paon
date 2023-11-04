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
	childMousePressEventHandled := 0

	// Widget tree root.
	root := NewRoot()

	// A 10x10 widget.
	childWidget := NewBaseWidget(
		LayoutFunc(func(co layout.Constraint) geometry.Size {
			return geometry.NewSize(10, 10)
		}),
		DrawerFunc(func(surface draw.Surface) {}),
		Listener(mouse.PressListener(func(event mouse.Event) {
			childMousePressEventHandled++
		})),
	)

	// A parent that position widget to bottom right corner.
	var parent *BaseLayout
	parent = NewBaseLayout(
		LayoutAlgo(func(co layout.Constraint, childrenPositions []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
			for child := parent.FirstChild(); child != nil; child = child.Next() {
				childSize := child.(layout.Layout).Layout(co)
				childrenPositions = append(childrenPositions,
					geometry.Rect(
						surfaceSize.Width()-childSize.Width(),
						surfaceSize.Height()-childSize.Height(),
						surfaceSize.Width(),
						surfaceSize.Height()),
				)
			}

			return childrenPositions, co.MaxSize
		}),
	)

	// Add widgets to tree.
	err := root.AppendChild(parent)
	require.NoError(t, err)

	err = parent.AppendChild(childWidget)
	require.NoError(t, err)

	// Trigger a first layout so widgets are positionned.
	root.Layout(layout.Constraint{
		MinSize:    geometry.Size{},
		MaxSize:    surfaceSize,
		ParentSize: surfaceSize,
		RootSize:   surfaceSize,
	})

	// Dispatch a mouse event that is out of child widget bounds.
	root.DispatchEvent(mouse.NewPress(geometry.NewVec2D(10, 10), mouse.Button1, keypress.ModNone))
	require.Equal(t, 0, childMousePressEventHandled, "event propagated to child widget")

	// Dispatch a mouse event that is within child widget bounds.
	root.DispatchEvent(mouse.NewPress(geometry.NewVec2D(95, 95), mouse.Button1, keypress.ModNone))
	require.Equal(t, 1, childMousePressEventHandled, "event not propagated to child widget")

	// Remove child to ensure it doesn't reveive event anymore.
	err = parent.RemoveChild(childWidget.Node())
	require.NoError(t, err)

	// No need to layout again.

	// Dispatch a mouse event that is within child widget bounds.
	root.DispatchEvent(mouse.NewPress(geometry.NewVec2D(95, 95), mouse.Button1, keypress.ModNone))
	require.Equal(t, 1, childMousePressEventHandled, "event propagated to unmounted child widget")
}
