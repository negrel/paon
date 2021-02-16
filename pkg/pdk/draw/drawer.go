package draw

type Drawer interface {
	Draw(Canvas)
}

// DrawFunc define function that edit a Canvas.
type DrawFunc func(Canvas)

// Draw implements the Drawer interface.
func (df DrawFunc) Draw(c Canvas) {
	df(c)
}
