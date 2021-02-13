package draw

// Context is the rendering context passed to Object for the layout step and the draw step.
type Context struct {
	Object     Object
	Canvas     Canvas
	Constraint Constraint
}
