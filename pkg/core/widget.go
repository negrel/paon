package core

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/negrel/debugo"
	"github.com/negrel/ginger/internal/helpers"
)

// Widget define any graphical element on the screen.
type Widget interface {
	/*
	 * --- GETTERS & SETTERS ---
	 */

	active() bool

	// Unique identifier of the widget.
	UID() uuid.UUID

	// Name of the widget (button, text, etc).
	Name() string

	// The next sibling widget at the same tree level.
	NextSibling() (Widget, bool)

	// owner of the tree that contain
	// this widget.
	owner() Manager

	// Layout that contain this widget.
	Parent() Layout

	// The previous sibling widget at the same tree level.
	PreviousSibling() (Widget, bool)

	// The widget position in parent childrens
	slot() int
	setSlot(int)

	/*
	 * --- METHODS ---
	 */

	// set the parent of the widget and the
	// slot position.
	adoptedBy(Layout, int)

	// remove the parent of the widget and
	// the slot.
	abandoned()

	// lay out the widget at the given relative
	// position (relative to parent).
	// this method is called by layout when
	// adopt adopt a widget child.
	layout(helpers.Point)

	// Render the widget if needed.
	Render()

	// performRender update the render object of
	// the widget while respecting the given
	// constraint. This function must be overwritten
	// by any widgets.
	PerformRender()
}

// WidgetCore is the core of every widget.
// It implements the common method for the
// Widget interface.
type WidgetCore struct {
	uid uuid.UUID

	// widget name
	name string

	// parent widget
	parent Layout
	// child position in parent children list
	_slot int

	// owner of the node tree
	// nil if not active state
	_owner Manager

	// position relative to the parent
	pos helpers.Point
	// need to be laid out for the
	// next frame.
	needLayout bool
	// need to be rendered for the next
	// frame.
	needRender bool

	// renderObj *render.Object
}

var _ Widget = &WidgetCore{}

// NewWidgetCore return a new widget core.
func NewWidgetCore(name string) *WidgetCore {
	return &WidgetCore{
		uid:        uuid.New(),
		name:       name,
		parent:     nil,
		_owner:     nil,
		needLayout: true,
		needRender: true,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// SECTION Getters & Setters

// UID return the widget unique ID.
func (wc *WidgetCore) UID() uuid.UUID {
	return wc.uid
}

// active implements the Widget interface.
func (wc *WidgetCore) active() bool {
	return wc._owner != nil
}

// Name implements the Widget interface.
func (wc *WidgetCore) Name() string {
	return wc.name
}

// NextSibling implements the Widget interface.
func (wc *WidgetCore) NextSibling() (Widget, bool) {
	if wc.parent == nil {
		return nil, false
	}

	return wc.parent.ChildAt(wc._slot + 1)
}

// owner implements the Widget interface.
func (wc *WidgetCore) owner() Manager {
	return wc._owner
}

// Parent implements the Widget interface.
func (wc *WidgetCore) Parent() Layout {
	return wc.parent
}

// PreviousSibling implements the Widget interface.
func (wc *WidgetCore) PreviousSibling() (Widget, bool) {
	if wc.parent == nil {
		return nil, false
	}

	return wc.parent.ChildAt(wc._slot - 1)
}

// slot implements the Widget interface.
func (wc *WidgetCore) slot() int {
	return wc._slot
}

// setSlot implements the Widget interface.
func (wc *WidgetCore) setSlot(n int) {
	if next, ok := wc.NextSibling(); ok {
		next.setSlot(n - 1)
	}

	wc._slot = n
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// SECTION Methods

// adoptedBy implements the Widget interface.
func (wc *WidgetCore) adoptedBy(parent Layout, slot int) {
	defer debugo.Printf("%s was adopted by %s", wc, parent)

	debugo.Assert(
		parent == nil,
		fmt.Sprintf("%s can't be adopted, the given parent is nil", wc),
	)

	wc.parent = parent
	wc._slot = slot
}

// abandoned implements the Widget interface.
func (wc *WidgetCore) abandoned() {
	wc.parent = nil
}

// layout implements the Widget interface.
func (wc *WidgetCore) layout(position helpers.Point) {

}

// Render implements the Widget interface.
func (wc *WidgetCore) Render() {

}

// PerformRender implements the Widget interface.
//
// Must be overwritten or will paniwc.
func (wc *WidgetCore) PerformRender() {
	debugo.Assert(
		false,
		fmt.Sprintf("%s widget doesn't implements PerformRender, you should overwrite it.", wc),
	)
}

// String implements the fmt.Stringer interface.
func (wc *WidgetCore) String() string {
	return fmt.Sprintf("%v-%v", wc.name, wc.uid)
}
