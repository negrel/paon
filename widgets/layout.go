package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/events/mouse"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
)

// BaseLayout define a generic layout widget.
// BaseLayout can either be used alone (see NewBaseLayout for the required options)
// or in composite struct.
// BaseLayout is an extension of BaseWidget.
// BaseLayout takes care of the following things for you:
// - Dispatch mouse events to children
type BaseLayout struct {
	*BaseWidget
}

// NewBaseLayout returns a new BaseLayout object configured with the given
// options.
// When embedding this layout in a composite struct, wrap the struct instance
// using Wrap widget options.
func NewBaseLayout(
	algo func(_ layout.Constraint, _ []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size),
	options ...LayoutOption,
) *BaseLayout {
	layout := newBaseLayout(algo, options...)

	return layout
}

func newBaseLayout(
	algo func(_ layout.Constraint, _ []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size),
	options ...LayoutOption,
) *BaseLayout {
	l := &BaseLayout{}

	// Children positions.
	childrenRects := []geometry.Rectangle{}
	// Layout constraint used for the latest layout.
	latestLayoutConstraint := layout.Constraint{}

	// Event handler that forwards mouse position to children.
	dispatchPositionnedEvent := func(relpos *geometry.Vec2D, event events.Event) {
		// If layout is root, trigger a layout to sync childrenRects with current
		// widget tree.
		if l.Node().Root() == l.Node() {
			l.Layout(latestLayoutConstraint)
		}

		child := l.Node().FirstChild()
		for _, boundingRect := range childrenRects {
			if boundingRect.Contains(*relpos) {
				*relpos = relpos.Sub(boundingRect.TopLeft())
				child.Unwrap().(events.Target).DispatchEvent(event)
			}

			child = child.Next()
		}
	}
	dispatchMouseEvent := func(event mouse.Event) {
		dispatchPositionnedEvent(&event.RelPosition, event)
	}

	layoutConf := &baseLayoutOption{
		BaseLayout:        l,
		widgetConstructor: NewBaseWidget,
		widgetOptions: []WidgetOption{
			Wrap(l),
			LayoutFunc(func(co layout.Constraint) (size geometry.Size) {
				latestLayoutConstraint = co

				childrenRects, size = algo(co, childrenRects[:0])
				return size
			}),
			DrawerFunc(func(surface draw.Surface) {
				child := l.Node().FirstChild()
				for _, boundingRect := range childrenRects {
					childDrawer := child.Unwrap().(draw.Drawer)
					subsurface := draw.NewSubSurface(surface, boundingRect)

					childDrawer.Draw(subsurface)

					child = child.Next()
				}
			}),
			// Dispatch mouse event to child.
			Listener(mouse.PressListener(dispatchMouseEvent)),
			Listener(mouse.UpListener(dispatchMouseEvent)),
			Listener(mouse.ClickListener(func(event mouse.ClickEvent) {
				dispatchPositionnedEvent(&event.MousePress.RelPosition, event)
			})),
		},
	}

	for _, option := range options {
		option(layoutConf)
	}

	l.BaseWidget = layoutConf.widgetConstructor(layoutConf.widgetOptions...)

	return l
}
