package render

// Renderable define any object that can produce a render Patch.
type Renderable interface {
	Render() Patch
}
