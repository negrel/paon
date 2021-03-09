package draw

type Drawable interface {
	Draw(Context)
}

type Script func(Context)

func (s Script) Draw(ctx Context) {
	s(ctx)
}
