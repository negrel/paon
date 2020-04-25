package events

import "time"

// ResizeHandler are function that can handle a Resize events.
type ResizeHandler = func(*ResizeEvent)

// ResizeListener define object that can listen to resize events.
type ResizeListener interface {
	OnResize(*ResizeEvent)
}

// Size define the size of an element by it's
// height and width.
type Size struct {
	Width  int
	Height int
}

// ResizeEvent is triggered when the user resize
// the terminal window.
type ResizeEvent struct {
	*event

	size           *Size
	wider, greater bool
}

// NewResizeEvent return a new ResizeEvent instance.
func NewResizeEvent(timeStamp time.Time, newSize, oldSize Size) *ResizeEvent {
	return &ResizeEvent{
		event: &event{
			evType:    ScrollEventType,
			timeStamp: timeStamp,
		},
		size:    &newSize,
		wider:   newSize.Width > oldSize.Width,
		greater: newSize.Height > oldSize.Height,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Height return the new height of the window.
func (re *ResizeEvent) Height() int {
	return re.size.Height
}

// Width return the new width of the window.
func (re *ResizeEvent) Width() int {
	return re.size.Width
}

// Size return the new size of the window.
func (re *ResizeEvent) Size() Size {
	return *re.size
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
