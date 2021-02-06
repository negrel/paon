package render

import "github.com/negrel/paon/internal/draw"

// Context is the rendering context passed to Object for the layout step and the draw step.
type Context struct {
	Object       Object
	ParentObject Object
	Canvas       *draw.Canvas
	Constraint   Constraint
}

func MakeContext(parent Object, constraint Constraint) Context {
	return Context{
		ParentObject: parent,
		Canvas:       draw.NewCanvas(constraint.Max),
		Constraint:   constraint,
	}
}
