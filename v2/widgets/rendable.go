package widgets

import "github.com/negrel/ginger/v2/render"

// Rendable is implementeb by widget that have a Render
// method that return the rendered Frame.
type Rendable interface {
	Render(Constraint) *render.Frame
}
