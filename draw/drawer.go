package draw

// Drawer define an object that can draw on a Canvas.
type Drawer interface {
	Draw(Surface)
}

// DrawerFunc define a function that implements the Drawer interface.
type DrawerFunc func(Surface)

// Draw implements the Drawable interface.
func (fn DrawerFunc) Draw(c Surface) {
	fn(c)
}
