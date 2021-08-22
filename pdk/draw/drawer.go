package draw

// Drawer define an object that can draw on a Canvas.
type Drawer interface {
	Draw(Canvas)
}

// DrawerFn define a function that implements the Drawer interface.
type DrawerFn func(Canvas)

// Draw implements the Drawable interface.
func (fn DrawerFn) Draw(c Canvas) {
	fn(c)
}
