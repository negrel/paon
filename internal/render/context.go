package render

// Context is the rendering context passed to Object for the layout step and the draw step.
type Context struct {
	Object Object
	Layer  *Layer
}

func NewContext(obj Object) *Context {
	return &Context{
		Object: obj,
		Layer:  NewLayer(),
	}
}
