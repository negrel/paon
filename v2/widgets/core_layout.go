package widgets

import (
	"image"
	"log"

	"github.com/negrel/ginger/v2/render"
	"github.com/negrel/ginger/v2/widgets/events"
)

var _ Widget = &CoreLayout{}
var _ Layout = &CoreLayout{}

// CoreLayout is the core element of layout
// widgets. CoreLayout is intended to be embedded in
// more advanced layout.
//
// CoreLayout focus on the widget tree structure
// and optimization.
type CoreLayout struct {
	*Core

	Children []Widget
}

// NewCoreLayout return a new core layout.
func NewCoreLayout(name string, children []Widget) *CoreLayout {
	cl := &CoreLayout{
		Core:     NewCore(name),
		Children: children,
	}

	// No cache for layouts.
	cl.cache = nil

	for _, child := range cl.Children {
		cl.AdoptChild(child)
		cl.AddListener(child)
	}

	return cl
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Attach overwrite Core attached method.
//
// Attach implements the Widgets interface.
func (cl *CoreLayout) Attach(owner Layout) {
	cl.owner = owner

	for _, child := range cl.Children {
		child.Attach(owner)
	}
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AdoptChild implements Branch interface.
func (cl *CoreLayout) AdoptChild(child Widget) {
	// Checking child ready to be adopted
	if child == nil ||
		child.Parent() != nil {

		log.Println("can't adopt the child. (child is nil or child parent is not nil)")
		log.Printf(" ├─ Child %v: %+v", child.Name(), child)
		log.Fatalf(" └─ Child parent %v: %+v", child.Parent().Name(), child.Parent())
	}

	// Checking that child is not parent this node.
	var node Widget = cl
	for node.Parent() != nil {
		if node == child {
			log.Printf("can't adopt child, child is an ancestor of node")
			log.Fatalf(" └─ Child %v: %+v", child.Name(), child)
		}

		node = node.Parent()
	}

	// Adopting the child
	child.setParent(cl)
	if cl.Attached() {
		child.Attach(cl.owner)
	}
}

// DropChild implements Branch interface.
func (cl *CoreLayout) DropChild(child Widget) {
	if child == nil ||
		child.Parent() == nil ||
		child.Attached() != cl.Attached() {

		log.Println("can't drop the child. (child is nil or child parent is nil or child attach state is different)")
		log.Printf(" ├─ Child %v: %+v", child.Name(), child)
		log.Fatalf(" └─ Child parent %v: %+v %v", child.Parent().Name(), child.Parent(), child.Parent().Attached())
	}

	child.setParent(nil)
	if child.Attached() {
		child.Detach()
	}
}

// AddListener method add the given child to events.Emitter listener
// if it implements one of the events.Listener interface.
func (cl *CoreLayout) AddListener(child Widget) {
	// ResizeListener
	if listener, ok := child.(events.ResizeListener); ok {
		events.Emitter.AddResizeListener(listener)
	}

	// ScrollListener
	if listener, ok := child.(events.ScrollListener); ok {
		events.Emitter.AddScrollListener(listener)
	}

	// ClickListener
	if listener, ok := child.(events.ClickListener); ok {
		events.Emitter.AddClickListener(listener)
	}
}

// Render implements Rendable interface.
func (cl *CoreLayout) Render(bounds image.Rectangle) *render.Frame {
	if cl.Attached() {
		// children bounds are relative to parent.
		childBounds := image.Rectangle{
			Min: image.Pt(0, 0),
			Max: bounds.Max.Sub(bounds.Min),
		}

		frame := cl.Rendering(childBounds)
		frame.Position = bounds.Min

		return frame
	}

	return nil
}
