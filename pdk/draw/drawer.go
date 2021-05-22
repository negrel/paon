package draw

type Drawer interface {
	Draw(Context)
}

// DrawerFn define a function that implements the Drawable interface.
type DrawerFn func(Context)

// Draw implements the Drawable interface.
func (fn DrawerFn) Draw(ctx Context) {
	fn(ctx)
}
