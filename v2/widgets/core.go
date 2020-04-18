package widgets

import (
	"log"

	"github.com/negrel/ginger/v2/render"
)

// ANCHOR CORE

var _ Widget = &Core{}

// Core is the core element of all widgets.
// Core is intended to be embed in more advanced
// widget.
type Core struct {
	parent Layout
	owner  Layout
	cache  Cache

	Draw func(Constraint) *render.Frame
}

// NewCore return a new core layout.
func NewCore() *Core {
	return &Core{}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Attached implements Widget interface.
func (c *Core) Attached() bool {
	return c.owner != nil
}

// Attach implements Widget interface.
func (c *Core) Attach(owner Layout) {
	c.owner = owner
}

// Detach implements Widget interface.
func (c *Core) Detach() {
	c.owner = nil
}

// Owner implements Widget interface.
func (c *Core) Owner() Layout {
	return c.owner
}

// Parent implements Widget interface.
func (c *Core) Parent() Layout {
	return c.parent
}

func (c *Core) setParent(parent Layout) {
	c.parent = parent
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Render implements Rendable interface.
func (c *Core) Render(co Constraint) *render.Frame {
	if c.cache.F != nil &&
		co.Bounds.Dx() >= c.cache.F.Patch.Width() &&
		co.Bounds.Dy() >= c.cache.F.Patch.Height() {

		return c.cache.F
	}

	return c.Draw(co)
}

/*****************************************************
 *****************************************************
 *****************************************************/
// ANCHOR CORE LAYOUT

var _ Widget = &CoreLayout{}
var _ Layout = &CoreLayout{}

// CoreLayout is the core element of layout
// widgets. CoreLayout is intended to be embed in
// more advanced layout.
type CoreLayout struct {
	*Core

	Children []Widget
	Draw     func(Constraint) *render.Frame
}

// NewCoreLayout return a new core layout.
func NewCoreLayout(children []Widget) *CoreLayout {
	return &CoreLayout{
		Core:     NewCore(),
		Children: children,
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
		log.Fatal("can't adopt the child. (child is nil or child parent is not nil)")
	}

	// Checking that child is not parent this node.
	var node Widget = cl
	for node.Parent() != nil {
		if node == child {
			log.Fatal("can't adopt child, child is an ancestor of node")
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
		log.Fatal("can't drop the child. (child is nil or child parent is nil or child attach state is different)")
	}

	child.setParent(nil)
	if child.Attached() {
		child.Detach()
	}
}

// Render implements Rendable interface.
func (cl *CoreLayout) Render(co Constraint) *render.Frame {
	if co == cl.cache.C {
		return cl.cache.F
	}

	frame := cl.Draw(co)

	// Updatin cache
	cl.cache.C = co
	cl.cache.F = frame

	return frame
}
