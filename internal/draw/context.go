package draw

import "github.com/negrel/paon/internal/render"

// Context is the rendering context passed to Object for the layout step and the draw step.
type Context struct {
	Object     Object
	Canvas     *render.Buffer
	Constraint Constraint
}

func MakeContext(constraint Constraint) Context {
	return Context{
		Canvas:     render.NewCanvas(constraint.Max),
		Constraint: constraint,
	}
}
