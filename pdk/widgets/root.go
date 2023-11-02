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

// Root define the root of a widget tree.
type Root struct {
	*BaseWidget
}

var _ Layout = &Root{}

// NewRoot returns a new Widget that can be used as a root.
func NewRoot() *Root {
	root := &Root{}

	root.BaseWidget = newBaseWidget(
		Wrap(root),
		NodeOptions(treevents.NodeConstructor(func(data any) tree.Node {
			return tree.NewRoot(data)
		})),
		LayoutFunc(func(co layout.Constraint) geometry.Size {
			child := root.FirstChild()
			if child != nil {
				return child.Unwrap().(layout.Layout).Layout(co)
			}

			return co.MinSize
		}),
		DrawerFunc(func(s draw.Surface) {
			child := root.FirstChild()
			if child == nil {
				return
			}

			child.Unwrap().(draw.Drawer).Draw(s)
		}),
	)

	// Dispatch click event.
	root.AddEventListener(events.ClickListener(func(event events.Click) {
		child := root.FirstChild()
		if child == nil {
			return
		}

		child.Unwrap().(pdkevents.Target).DispatchEvent(event)
	}))

	return root
}
