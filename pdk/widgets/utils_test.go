package widgets

import (
	draw "github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
)

func newMountedLayout() Layout {
	root := NewRoot()
	layout := newLayout()

	root.SetChild(layout)

	return layout
}

func newLayout() Layout {
	layout := NewBaseLayout(
		WidgetOptions(
			LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
				return layout.NewBox(geometry.Rectangle{})
			})),
			Drawer(draw.DrawerFn(func(c draw.Context) {})),
		),
	)

	return layout
}

func newWidget(options ...WidgetOption) Widget {
	options = append(
		[]WidgetOption{
			Drawer(draw.DrawerFn(func(c draw.Context) {})),
			LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
				return layout.NewBox(geometry.Rectangle{})
			})),
		}, options...)

	return NewBaseWidget(options...)
}
