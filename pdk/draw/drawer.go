package draw

// Drawer define an object that can draw on a Canvas.
type Drawer interface {
	Draw(Surface)
}

// DrawerFn define a function that implements the Drawer interface.
type DrawerFn func(Surface)

// Draw implements the Drawable interface.
func (fn DrawerFn) Draw(c Surface) {
	fn(c)
}
