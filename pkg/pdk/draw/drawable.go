package draw

type Drawable interface {
	Draw(Context)
}

// Script define a function that implements the Drawable interface.
type Script func(Context)

// Draw implements the Drawable interface.
func (s Script) Draw(ctx Context) {
	s(ctx)
}
