package widgets

// func newTestRoot() *Root {
// 	root, _ := NewRoot(nil)
// 	return root
// }

// func newMountedLayout() Layout {
// 	root := newTestRoot()
// 	layout := newLayout()

// 	root.SetChild(layout)

// 	return layout
// }

// func newLayout() Layout {
// 	layout := NewBaseLayout(
// 		WidgetOptions(
// 			LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
// 				return layout.NewBox(geometry.Rectangle{})
// 			})),
// 			Drawer(draw.DrawerFn(func(c draw.Canvas) {})),
// 		),
// 	)

// 	return layout
// }

// func newWidget(options ...WidgetOption) Widget {
// 	options = append(
// 		[]WidgetOption{
// 			Drawer(draw.DrawerFn(func(c draw.Canvas) {})),
// 			LayoutManager(layout.ManagerFn(func(c layout.Constraint) layout.BoxModel {
// 				return layout.NewBox(geometry.Rectangle{})
// 			})),
// 		}, options...)

// 	return NewBaseWidget(options...)
// }
