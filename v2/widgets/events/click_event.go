package events

import (
	"image"
	"time"
)

// ClickHandler are function that can handle a Click events.
type ClickHandler = func(*ClickEvent)

// ClickListener define object that can listen to click events.
type ClickListener interface {
	// IsListening return wether or not the widget
	// listen to the click at the given position.
	IsListening(Position) bool

	OnClick(*ClickEvent)
}

// Position define a position by it's X and Y coordinate.
type Position image.Point

// Button define the mouse button that was clicked
type Button int

// String implements fmt.Stringer interface.
func (b Button) String() string {
	switch b {
	case LeftButton:
		return "left click"
	case MiddleButton:
		return "middle click"
	case RightButton:
		return "right click"
	case Button4:
		return "button 4"
	case Button5:
		return "button 5"
	case Button6:
		return "button 6"
	case Button7:
		return "button 7"
	case Button8:
		return "button 8"

	default:
		return ""
	}
}

// Possible mouse buttons
const (
	LeftButton   Button = 1 << iota // Usually left mouse button.
	MiddleButton                    // Usually the middle mouse button.
	RightButton                     // Usually the right mouse button.
	Button4                         // Often a side button (thumb/next).
	Button5                         // Often a side button (thumb/prev).
	Button6
	Button7
	Button8
)

// ClickEvent is triggered when the user click in
// the terminal window.
type ClickEvent struct {
	*event

	button Button
	pos    Position
}

// NewClickEvent return a new ClickEvent instance.
func NewClickEvent(timeStamp time.Time, button Button, pos Position) *ClickEvent {
	return &ClickEvent{
		event: &event{
			evType:    ScrollEventType,
			timeStamp: timeStamp,
		},
		button: button,
		pos:    pos,
	}
}

/*****************************************************
 ***************** Getters & Setters *****************
 *****************************************************/
// ANCHOR Getters & Setters

// Button return the clicked button.
func (ce *ClickEvent) Button() Button {
	return ce.button
}

// Position return the coordinate of the mouse.
func (ce *ClickEvent) Position() Position {
	return ce.pos
}
