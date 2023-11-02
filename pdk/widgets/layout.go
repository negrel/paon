package widgets

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	pdkevents "github.com/negrel/paon/pdk/events"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/tree"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

// Layout is an extension of the Widget interface that adds the support for
// children widgets. Howerver, it is strongly recommended to create custom
// layouts using the BaseLayout implementation.
type Layout interface {
	Widget
}

var _ Layout = &BaseLayout{}

// BaseLayout define a basic implementation of the Layout interface.
// BaseLayout can either be used alone (see NewBaseLayout for the required options)
// or in composite struct.
// BaseLayout is an extension of BaseWidget.
// BaseLayout takes care of the following things for you:
// - Dispatch LifeCycleEvents to children
type BaseLayout struct {
	*BaseWidget

	layoutAlgo func(_ layout.Constraint, _ []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size)
}

// NewBaseLayout returns a new BaseLayout object configured with the given
// options.
// When embedding this layout in a composite struct, wrap the struct instance
// using Wrap widget options.
func NewBaseLayout(options ...LayoutOption) *BaseLayout {
	layout := newBaseLayout(options...)

	return layout
}

func newBaseLayout(options ...LayoutOption) *BaseLayout {
	l := &BaseLayout{}

	// Children positions.
	childrenRects := []geometry.Rectangle{}
	// Layout constraint used for the latest layout.
	latestLayoutConstraint := layout.Constraint{}

	layoutConf := &baseLayoutOption{
		BaseLayout:        l,
		widgetConstructor: NewBaseWidget,
		widgetOptions: []WidgetOption{
			Wrap(l),
			NodeOptions(treevents.NodeConstructor(tree.NewNode)),
			LayoutFunc(func(co layout.Constraint) (size geometry.Size) {
				latestLayoutConstraint = co

				childrenRects, size = l.layoutAlgo(co, childrenRects[:0])
				return size
			}),
			DrawerFunc(func(surface draw.Surface) {
				child := l.FirstChild()
				for _, boundingRect := range childrenRects {
					childDrawer := child.Unwrap().(draw.Drawer)
					subsurface := draw.NewSubSurface(surface, boundingRect)

					childDrawer.Draw(subsurface)

					child = child.Next()
				}
			}),
			// Dispatch click event to child.
			Listener(events.ClickListener(func(event events.Click) {
				// If layout is root, trigger a layout to sync childrenRects with current
				// widget tree.
				if l.Root().IsSame(l) {
					l.Layout(latestLayoutConstraint)
				}

				child := l.FirstChild()
				for _, boundingRect := range childrenRects {
					if boundingRect.Contains(event.Position) {
						child.Unwrap().(pdkevents.Target).DispatchEvent(event)
					}

					child = child.Next()
				}
			})),
		},
	}

	for _, option := range options {
		option(layoutConf)
	}

	l.BaseWidget = layoutConf.widgetConstructor(layoutConf.widgetOptions...)

	return l
}
