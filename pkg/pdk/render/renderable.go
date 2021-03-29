package render

// Renderable define elements that can be rendered to a Surface.
type Renderable interface {
	NeedRendering() bool
	Render() Patch
}
