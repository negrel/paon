package ui

import (
	"github.com/negrel/ginger/internal/utils"
	"github.com/negrel/ginger/v3/render"
)

// Widget define any graphical element on the screen.
type Widget interface {
	/* NOTE
	 * --- GETTERS & SETTERS ---
	 */

	// Name of the widget (button, text, etc).
	Name() string

	// The next sibling widget at the same tree level.
	NextSibling() Widget

	// owner of the tree that contain
	// this widget.
	owner() Manager

	// Layout that contain this widget.
	Parent() Layout

	// The previous sibling widget at the same tree level.
	PreviousSibling() Widget

	/* NOTE
	 * --- METHODS ---
	 */

	active() bool

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
	layout(utils.Point)

	// Render the widget if needed.
	Render(render.Constraints)

	// performRender update the render object of
	// the widget while respecting the given
	// constraint. This function must be overwritten
	// by any widgets.
	PerformRender(render.Constraints)
}

// WidgetCore define the WidgetCore of every widgets.
// it implements the common method for the
// Widget interface.
type WidgetCore struct {
	// widget name
	name string

	// parent widget
	parent Layout
	// position in parent child list.
	slot int

	// owner of the node tree
	// nil if not active state
	_owner Manager

	// position relative to the parent
	pos utils.Point
	// need to be laid out for the
	// next frame.
	needLayout bool
	// need to be rendered for the next
	// frame.
	needRender bool

	renderObj *render.Object
}

var _ Widget = &WidgetCore{}

// NewWidgetCore return a new widget core.
func NewWidgetCore(name string) *WidgetCore {
	return &WidgetCore{
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
// ANCHOR Getters & Setters

// active implements the Widget interface.
func (wc *WidgetCore) active() bool {
	return wc._owner != nil
}

// Name implements the Widget interface.
func (wc *WidgetCore) Name() string {
	return wc.name
}

// NextSibling implements the Widget interface.
func (wc *WidgetCore) NextSibling() Widget {
	if wc.parent == nil {
		return nil
	}

	return wc.parent.ChildAt(wc.slot - 1)
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
func (wc *WidgetCore) PreviousSibling() Widget {
	if wc.parent == nil {
		return nil
	}
	return wc.parent.ChildAt(wc.slot + 1)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// adoptedBy implements the Widget interface.
func (wc *WidgetCore) adoptedBy(parent Layout, slot int) {
	wc.parent = parent
	wc.slot = slot
}

// abandoned implements the Widget interface.
func (wc *WidgetCore) abandoned() {
	wc.parent = nil
	wc.slot = -1
}

// layout implements the Widget interface.
func (wc *WidgetCore) layout(position utils.Point) {
	if !wc.needLayout {
		wc.pos = position

	}
}

// Render implements the Widget interface.
func (wc *WidgetCore) Render(constraints render.Constraints) {
	if constraints.Equal(wc.renderObj.Constraints) {
		if wc.needRender {
			wc.PerformRender(constraints)
		}

		return
	}

	wc.PerformRender(constraints)
}

// PerformRender implements the Widget interface.
//
// Must be overwritten or will paniwc.
func (wc *WidgetCore) PerformRender(_ render.Constraints) {
	panic("The WidgetCore widget doesn't implements PerformRender, you should overwrite it.")
}
