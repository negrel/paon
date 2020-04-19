package widgets

import (
	"log"
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

	for _, child := range cl.Children {
		cl.AdoptChild(child)
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

	// Cache not valid anymore, need a new render frame.
	cl.cache.valid = false

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

	// Cache not valid anymore, need a total new render frame.
	cl.cache.valid = false

	child.setParent(nil)
	if child.Attached() {
		child.Detach()
	}
}
