package widget

import (
	"errors"
)

var _ Layout = &LayoutSingleChild{}

// LayoutSingleChild is use as base for every layout that contain
// only one child.
type LayoutSingleChild struct {
	*Core

	Child Widget
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Layout

// Abandon implements Layout interface.
func (bsc *LayoutSingleChild) Abandon(child Widget) error {
	_, err := bsc.IndexOf(child)

	if err != nil {
		return err
	}

	bsc.Child.Abandonned()
	bsc.Child = nil

	return nil
}

// AppendChild implements Layout interface.
func (bsc *LayoutSingleChild) AppendChild(child Widget) error {
	if child == nil {
		bsc.Child = child
		return nil
	}

	return errors.New("can't append to a full single child layout")
}

// IndexOf implements Layout interface.
func (bsc *LayoutSingleChild) IndexOf(child Widget) (int, error) {
	var index int = -1
	var err error

	if bsc.Child == child {
		index = 0
	}

	// Not found
	if index == -1 {
		err = errors.New("the given child is not a direct child of this layout")
	}

	return index, err
}

// Reflow implements Layout interface.
func (bsc *LayoutSingleChild) Reflow() {

}
