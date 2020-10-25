package widgets

import (
	"github.com/negrel/debuggo/pkg/assert"
)

const NotFound = -1

// widgetList is a list of Widget.
type widgetList struct {
	list   []Widget
	length int
}

// NewWidgetList return a new empty widgetList
func NewWidgetList() *widgetList {
	return &widgetList{
		list: make([]Widget, 0, 8),
	}
}

// Append the given child to the end of the list.
func (wl *widgetList) Append(child Widget) {
	wl.list = append(wl.list, child)
	wl.length++
}

func (wl *widgetList) InsertBefore(child, reference Widget) {
	index := wl.IndexOf(reference)
	assert.NotEqualf(index, NotFound, "%v reference child is not in the list", reference)

	wl.Insert(child, index)
}

// Insert the given child node at the given index.
func (wl *widgetList) Insert(child Widget, index int) {
	assert.NotNil(child, "child must be non-nil")
	assert.GreaterOrEqualf(index, 0, "%v is an invalid index (should be >= 0)")

	// Inserting the element and shifting the other
	last := wl.length - 1
	wl.list = append(wl.list, wl.list[last])
	copy(wl.list[index+1:], wl.list[index:last])
	wl.list[index] = child

	wl.length++
}

// IndexOf return the index of the given child node. If the child is not in the
// list -1 is returned.
func (wl *widgetList) IndexOf(child Widget) int {
	assert.NotNil(child, "child must be non-nil")

	for i, widget := range wl.list {
		if widget == child {
			return i
		}
	}

	return NotFound
}

// Get return the widget present at the given index in the list.
func (wl *widgetList) Get(index int) Widget {
	assert.GreaterOrEqualf(index, 0, "%v is an invalid index (should be >= 0)")
	assert.Lessf(index, wl.length, "%v index overflow the widget list (length: %v)", index, wl.length)

	return wl.list[index]
}

// Length return the length of the widget list.
func (wl *widgetList) Length() int {
	return wl.length
}

// Remove remove the given child from list
func (wl *widgetList) Remove(child Widget) {
	assert.NotNil(child, "child must be non-nil")

	index := wl.IndexOf(child)
	assert.GreaterOrEqualf(index, 0, "%v is not in the list")

	wl.list = append(wl.list[:index], wl.list[index+1:]...)
	wl.length--
}
