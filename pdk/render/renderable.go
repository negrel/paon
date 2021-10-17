package render

// Renderable define any object that can be rendered.
type Renderable interface {
	// Render performs the rendering of this object in the given rendering context.
	Render(Context)

	// Layer returns the render layer of the renderable object.
	Layer() *Layer
}
