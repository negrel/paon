package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/tree"
	treevents "github.com/negrel/paon/pdk/tree/events"
)

// Root define the root of a widget tree.
type Root struct {
	*BaseLayout
}

var _ Layout = &Root{}

// NewRoot returns a new Widget that can be used as a root.
func NewRoot() *Root {
	root := &Root{}

	root.BaseLayout = NewBaseLayout(
		WidgetOptions(
			Wrap(root),
			NodeOptions(treevents.NodeConstructor(func(data any) tree.Node {
				return tree.NewRoot(data)
			})),
			DrawerFunc(func(s draw.Surface) {
				child := root.FirstChild()
				if child == nil {
					return
				}

				child.Unwrap().(draw.Drawer).Draw(s)
			}),
		),
		LayoutAlgo(func(co layout.Constraint, childrenRects []geometry.Rectangle) ([]geometry.Rectangle, geometry.Size) {
			child := root.FirstChild()
			if child != nil {
				childSize := child.Unwrap().(layout.Layout).Layout(co)
				childrenRects = append(childrenRects, geometry.Rect(0, 0, childSize.Width(), childSize.Height()))
			}

			return childrenRects, co.MinSize
		}),
	)

	return root
}
