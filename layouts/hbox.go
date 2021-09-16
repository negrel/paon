package layouts

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	"github.com/negrel/paon/pdk/layout/manager"
	"github.com/negrel/paon/pdk/math"
	"github.com/negrel/paon/pdk/widgets"
)

var _ widgets.Layout = &HBox{}

// HBox is a layout that align contained widgets horizontally.
type HBox struct {
	*widgets.BaseLayout
}

// NewHBox returns a new HBox layout.
func NewHBox() *HBox {
	hbox := &HBox{}
	hbox.BaseLayout = widgets.NewBaseLayout(
		widgets.WidgetOptions(
			widgets.Wrap(hbox),
			widgets.LayoutManager(HBoxLayout(hbox)),
			widgets.Drawer(LayoutDrawer(hbox)),
		),
	)

	return hbox
}

// HBoxLayout returns the layout.Manager used by HBox.
func HBoxLayout(l widgets.Layout) layout.Manager {
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
					// Update children layout
					for child != nil {
						childBox := child.Layout(childConstraint)
						childBox.SetOrigin(childOrigin)

						child = child.NextSibling()
					}

					break layoutLoop
				}

				// Layout the child
				childBox := child.Layout(childConstraint)
				childBox.SetOrigin(childOrigin)

				// Update child origin for next iter
				childSize := childBox.MarginBox().Size()
				childOrigin = childOrigin.Add(geometry.Pt(childSize.Width(), 0))

				// Update size of this HBox
				width += childBox.MarginBox().Width()
				height = math.Max(height, childBox.MarginBox().Height())

				childConstraint.MaxSize = geometry.NewSize(
					childConstraint.MaxSize.Width()-childBox.MarginBox().Width(),
					c.MaxSize.Height(),
				)

				child = child.NextSibling()
			}

			box := layout.NewBox(geometry.NewSize(width, height)).
				ApplyMargin(style).
				ApplyBorder(style).
				ApplyPadding(style)

			return box
		}))
	})
}

// LayoutDrawer is a generic drawer that draw the box of the drawer and then
// iterate over the widgets.Layout children and layout.
func LayoutDrawer(l widgets.Layout) draw.Drawer {
	return draw.Box(l, l, draw.DrawerFn(func(c draw.Canvas) {
		lContentBox := l.Box().ContentBox()
		layoutBox := geometry.Rectangle{
			Min: geometry.Vec2D{},
			Max: geometry.Pt(lContentBox.Width(), lContentBox.Height()),
		}

		child := l.FirstChild()

		for child != nil {
			childBox := child.Box()

			// If the child is not in the parent contant box or the child size is 0, skip the draw.
			if childBox == nil || !layoutBox.Contains(childBox.MarginBox().Min) || childBox.MarginBox().Empty() {
				break
			}

			subcanvas := draw.NewSubCanvas(c, childBox.BorderBox())

			child.Draw(subcanvas)

			child = child.NextSibling()
		}
	}))
}
