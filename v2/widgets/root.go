package widgets

import (
	"image"
	"log"

	"github.com/negrel/ginger/v2/render"
)

var _ Layout = &Root{}

// Root is the root of the widgets tree.
type Root struct {
	*CoreLayout
}

// ROOT return a new Root object that you can use as
// your widget root tree.
func ROOT(child Widget) *Root {
	r := &Root{
		// No NewCoreLayout to get attached child.
		CoreLayout: &CoreLayout{
			Core:     NewCore("root"),
			Children: []Widget{child},
		},
	}

	for _, child := range r.Children {
		r.AdoptChild(child)
		r.AddListener(child)
	}

	r.Rendering = r.rendering

	return r
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Attached implements Widget interface.
func (r *Root) Attached() bool {
	return true
}

// Child return the root widget child.
func (r *Root) Child() Widget {
	return r.Children[0]
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AdoptChild overwrite CoreLayout method.
//
// AdoptChild implements Layout interface.
func (r *Root) AdoptChild(child Widget) {
	// Checking child ready to be adopted
	if child == nil ||
		child.Parent() != nil {
		log.Fatal("can't adopt the child. (child is nil or child parent is not nil)")
	}

	// Adopting the child
	child.setParent(r)
	child.Attach(r)
}

// Rendering implements Widget interface.
func (r *Root) rendering(bounds image.Rectangle) *render.Frame {
	frame := render.NewFrame(bounds.Dx(), bounds.Dy())
	child := r.Child()

	if child == nil {
		return frame
	}

	return child.Render(bounds)
}
