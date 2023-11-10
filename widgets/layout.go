package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

// LayoutAlgo define closures responsible of computing children position and
// it's own size. See HBox or VBox for an example.
type LayoutAlgo = func(_ layout.Constraint, _ []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size)

// BaseLayout is a basic layout widget.
// It can either be used alone or embedded within a struct.
// BaseLayout
// - Dispatch mouse events to children
// - Drawing children
type BaseLayout struct {
	Widget

	// Children positions.
	childrenRects []geometry.Rectangle
	// Layout constraint used for the latest layout.
	latestLayoutConstraint layout.Constraint
	// Layout algorithm.
	layoutAlgo func(_ layout.Constraint, _ []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size)
}

// NewBaseLayout returns a BaseLayout that embed the given widget.
func NewBaseLayout(
	widget Widget,
	algo LayoutAlgo,
) *BaseLayout {
	bl := &BaseLayout{
		Widget:        widget,
		childrenRects: []geometry.Rectangle{},
		layoutAlgo:    algo,
	}

	// Event handler that forwards mouse position to children.
	dispatchClickEvent := func(event mouse.ClickEvent) {
		// If layout is root, trigger a layout to sync childrenRects with current
		// widget tree.
		if bl.Node().Root() == bl.Node() {
			bl.Layout(bl.latestLayoutConstraint)
		}

		child := bl.Node().FirstChild()
		for _, boundingRect := range bl.childrenRects {
			if boundingRect.Contains(event.RelPosition) {
				event.RelPosition = event.RelPosition.Sub(boundingRect.TopLeft())
				child.Unwrap().(events.Target).DispatchEvent(event)
			}

			child = child.Next()
		}
	}
	dispatchMouseEvent := func(event mouse.Event) {
		// If layout is root, trigger a layout to sync childrenRects with current
		// widget tree.
		if bl.Node().Root() == bl.Node() {
			bl.Layout(bl.latestLayoutConstraint)
		}

		child := bl.Node().FirstChild()
		for _, boundingRect := range bl.childrenRects {
			if boundingRect.Contains(event.RelPosition) {
				event.RelPosition = event.RelPosition.Sub(boundingRect.TopLeft())
				child.Unwrap().(events.Target).DispatchEvent(event)
			}

			child = child.Next()
		}
	}

	// Dispatch mouse event to child.
	widget.AddEventListener(mouse.PressListener(dispatchMouseEvent))
	widget.AddEventListener(mouse.UpListener(dispatchMouseEvent))
	widget.AddEventListener(mouse.ClickListener(dispatchClickEvent))

	return bl
}

// Layout implements layout.Layout.
func (bl *BaseLayout) Layout(co layout.Constraint) (size geometry.Size) {
	bl.latestLayoutConstraint = co

	bl.childrenRects, size = bl.layoutAlgo(co, bl.childrenRects[:0])
	return size
}

// Draw implements draw.Drawer.
func (bl *BaseLayout) Draw(surface draw.Surface) {
	child := bl.Node().FirstChild()
	for _, boundingRect := range bl.childrenRects {
		childDrawer := child.Unwrap().(draw.Drawer)
		subsurface := draw.NewSubSurface(surface, boundingRect)

		childDrawer.Draw(subsurface)

		child = child.Next()
	}
}
