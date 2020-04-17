package widgets

import "github.com/negrel/ginger/v2/render"

// Rendable is implementeb by widget that have a Render
// method that return a render Frame.
type Rendable interface {
	Render(Constraint) *render.Frame
}
