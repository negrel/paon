package layouts

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/layout/manager"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/pdk/widgets"
)

var _ widgets.Layout = &HBox{}

// VBox is a layout that align contained widgets vertically.
type VBox struct {
	*widgets.BaseLayout
}

// NewVBox returns a new HBox layout.
func NewVBox() *VBox {
	vbox := &VBox{}
	vbox.BaseLayout = widgets.NewBaseLayout(
		widgets.WidgetOptions(
			widgets.Wrap(vbox),
			widgets.LayoutManager(VBoxLayout(vbox)),
			widgets.Drawer(LayoutDrawer(vbox)),
		),
	)

	return vbox
}

// VBoxLayout returns the layout.Manager used by VBox.
func VBoxLayout(l widgets.Layout) layout.Manager {
	return layout.ManagerFn(func(c layout.Constraint) *layout.Box {
		style := l.Theme()

		return manager.Block(style, c, layout.ManagerFn(func(c layout.Constraint) *layout.Box {
			width, height := c.MinSize.Width(), c.MinSize.Height()

			child := l.FirstChild()
			childConstraint := layout.Constraint{
				MinSize:    geometry.Size{},
				MaxSize:    c.MaxSize,
				ParentSize: geometry.Size{},
				RootSize:   c.RootSize,
			}
			childOrigin := geometry.Vec2D{}
		layoutLoop:
			for child != nil {
				// All the space is used
				noFreeSpace := childConstraint.MaxSize.Width() == 0 || childConstraint.MaxSize.Height() == 0
				if noFreeSpace {
					updateOrigin(child, childConstraint, childOrigin)

					break layoutLoop
				}

				// Layout the child
				childBox := child.Layout(childConstraint)
				childBox.SetOrigin(childOrigin)

				// Update child origin for next iter
				childSize := childBox.MarginBox()
				childOrigin = childOrigin.Add(geometry.Pt(0, childSize.Height()))

				// Update size of this VBox
				height += childBox.MarginBox().Height()
				width = math.Max(width, childBox.MarginBox().Width())

				// Update constraint for next iter
				childConstraint.MaxSize = geometry.NewSize(
					c.MaxSize.Width(),
					childConstraint.MaxSize.Height()-childBox.MarginBox().Height(),
				)

				child = child.NextSibling()
			}

			size := c.ApplyOnSize(geometry.NewSize(width, height))

			box := layout.NewBox(size).
				ApplyMargin(style).
				ApplyBorder(style).
				ApplyPadding(style)

			return box
		}))
	})
}

// updateOrigin to the given widgets.Widget and all its next sibling.
func updateOrigin(child widgets.Widget, constraint layout.Constraint, origin geometry.Vec2D) {
	for child != nil {
		child.Layout(constraint).
			SetOrigin(origin)

		child = child.NextSibling()
	}
}
