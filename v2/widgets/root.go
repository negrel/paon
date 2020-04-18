package widgets

import (
	"image"

	"github.com/negrel/ginger/v2/render"
)

var _ Layout = &Root{}

// Root is the root of the widgets tree.
type Root struct {
	*CoreLayout
	Child Widget
}

// ROOT return a new Root object that you can use as
// your widget root tree.
func ROOT(child Widget) *Root {
	r := &Root{
		CoreLayout: NewCoreLayout([]Widget{child}),
		Child:      child,
	}

	r.AdoptChild(child)

	r.Draw = r.draw

	return r
}

/*****************************************************
 ********************* Interface *********************
 *****************************************************/
// ANCHOR Interface

// Widgets

// Attached implements Widget interface.
func (r *Root) Attached() bool {
	return true
}

// Rendable

// Render implements Rendable interface.
func (r *Root) Render(co Constraint) *render.Frame {
	if co == r.cache.C {
		return r.cache.F
	}

	return r.Draw(co)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

func (r *Root) draw(c Constraint) *render.Frame {
	frame := render.NewFrame(image.Pt(0, 0), c.Bounds.Dx(), c.Bounds.Dy())

	frame.Add(r.Child.Render(c))

	return frame
}
