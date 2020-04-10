package layout

import "github.com/negrel/ginger/v1/widget"

var _ Layout = &BaseMultipleChild{}

// BaseMultipleChild is use as base for every layout that contain
// multiple children.
type BaseMultipleChild struct {
	*widget.Base

	Children []widget.Widget
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Layout

// AppendChild implements Layout interface.
func (bmc *BaseMultipleChild) AppendChild(child widget.Widget) error {
	bmc.Children = append(bmc.Children, child)

	return nil
}

// Reflow implements Layout interface.
func (bmc *BaseMultipleChild) Reflow() {

}
