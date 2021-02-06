package render

type Renderer interface {
	Layout(ctx Context)
	Draw(ctx Context)
}
