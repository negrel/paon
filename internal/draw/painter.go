package draw

type Painter interface {
	Layout(ctx Context)
	Draw(ctx Context)
}
