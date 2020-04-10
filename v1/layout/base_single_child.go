package layout

import (
	"errors"

	"github.com/negrel/ginger/v1/widget"
)

var _ Layout = &BaseSingleChild{}

// BaseSingleChild is use as base for every layout that contain
// only one child.
type BaseSingleChild struct {
	*widget.Base

	Child widget.Widget
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Layout

// AppendChild implements Layout interface.
func (bsc *BaseSingleChild) AppendChild(child widget.Widget) error {
	if child == nil {
		bsc.Child = child
		return nil
	}

	return errors.New("can't append to a full single child layout")
}

// Reflow implements Layout interface.
func (bsc *BaseSingleChild) Reflow() {

}
