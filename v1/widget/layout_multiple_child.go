package widget

import (
	"errors"
)

// ChildList is a list of widget
type ChildList []Widget

// ForEach apply the given function for each child of the list.
func (cl ChildList) ForEach(fn func(int, Widget)) {
	for i, v := range cl {
		fn(i, v)
	}
}

var _ Layout = &LayoutMultipleChild{}

// LayoutMultipleChild is use as base for every layout that contain
// multiple children.
type LayoutMultipleChild struct {
	*Core

	Children ChildList
}

/*****************************************************
 ********************* Private ***********************
 *****************************************************/
// ANCHOR Private

func (bmc *LayoutMultipleChild) remove(s int) {
	bmc.Children = append(bmc.Children[:s], bmc.Children[s+1:]...)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Layout

// Abandon implements Layout interface.
func (bmc *LayoutMultipleChild) Abandon(child Widget) error {
	index, err := bmc.IndexOf(child)

	if err != nil {
		return err
	}

	bmc.Children[index].Abandonned()
	bmc.remove(index)

	return nil
}

// AppendChild implements Layout interface.
func (bmc *LayoutMultipleChild) AppendChild(child Widget) error {
	bmc.Children = append(bmc.Children, child)

	return nil
}

// ForEach implements Layout interface.
func (bmc *LayoutMultipleChild) ForEach(fn func(int, Widget)) {
	bmc.Children.ForEach(fn)
}

// IndexOf implements Layout interface.
func (bmc *LayoutMultipleChild) IndexOf(child Widget) (int, error) {
	var index int = -1
	var err error

	for i, widget := range bmc.Children {
		if widget == child {
			index = i
			break
		}
	}

	// Not found
	if index == -1 {
		err = errors.New("the given child is not a direct child of this layout")
	}

	return index, err
}

// Reflow implements Layout interface.
func (bmc *LayoutMultipleChild) Reflow() {

}
