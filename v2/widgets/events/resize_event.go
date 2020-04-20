package events

import "time"

// ResizeListener define object that can listen to resize events.
type ResizeListener interface {
	OnResize(*ResizeEvent)
}

// Size define the size of an element by it's
// height and width.
type size struct {
	width  int
	height int
}

// ResizeEvent is triggered when the user resize
// the terminal window.
type ResizeEvent struct {
	*event

	size           *size
	wider, greater bool
}

// NewResizeEvent return a new ResizeEvent instance.
func NewResizeEvent(timeStamp time.Time, newSize, oldSize size) *ResizeEvent {
	return &ResizeEvent{
		event: &event{
			evType:    ScrollEventType,
			timeStamp: timeStamp,
		},
		size:    &newSize,
		wider:   newSize.width > oldSize.width,
		greater: newSize.height > oldSize.height,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Height return the new height of the window.
func (re *ResizeEvent) Height() int {
	return re.size.height
}

// Width return the new width of the window.
func (re *ResizeEvent) Width() int {
	return re.size.width
}

// IsWider return wether or not the window width is larger
// since last resize event.
func (re *ResizeEvent) IsWider() bool {
	return re.wider
}

// IsGreater return wether or not the window height is larger
// since last resize event.
func (re *ResizeEvent) IsGreater() bool {
	return re.greater
}
