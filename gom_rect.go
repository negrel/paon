package gom

// GOMRect represents a rectangle.
type GOMRect interface {
	/* GETTERS & SETTERS (props) */
	X() int
	Y() int
	Width() int
	Height() int
	Top() int
	Right() int
	Bottom() int
	Left() int
}

var _ GOMRect = &gomRect{}

type gomRect struct {
	x      int
	y      int
	width  int
	height int
	top    int
	right  int
	bottom int
	left   int
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// X coordinate of the DOMRect's origin.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/x
func (r *gomRect) X() int {
	return r.x
}

// Y coordinate of the DOMRect's origin.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/y
func (r *gomRect) Y() int {
	return r.y
}

// Width of the DOMRect.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/width
func (r *gomRect) Width() int {
	return r.width
}

// Height of the DOMRect.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/height
func (r *gomRect) Height() int {
	return r.height
}

// Tops returns the top coordinate value of the DOMRect.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/top
func (r *gomRect) Top() int {
	return r.top
}

// Right returns the right coordinate value of the DOMRect.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/right
func (r *gomRect) Right() int {
	return r.right
}

// Bottom returns the bottom coordinate value of the DOMRect.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/bottom
func (r *gomRect) Bottom() int {
	return r.bottom
}

// Left returns the left coordinate value of the DOMRect.
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly/left
func (r *gomRect) Left() int {
	return r.left
}
