package draw

// Drawer define an object that can draw on a Canvas.
type Drawer interface {
	Draw(*Context)
}

// DrawerFn define a function that implements the Drawer interface.
type DrawerFn func(*Context)

// Draw implements the Drawable interface.
func (fn DrawerFn) Draw(ctx *Context) {
	fn(ctx)
}
