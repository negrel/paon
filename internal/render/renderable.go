package render

type Renderable interface {
	Render() Patch
}
