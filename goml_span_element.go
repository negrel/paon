package gom

// GOMLSpanElement define a <span> element
// and embbed the GOMLElement.
type GOMLSpanElement struct {
	gomlElement
}

var _ GOMLElement = &GOMLSpanElement{}
var _ Element = &GOMLSpanElement{}
var _ Node = &GOMLSpanElement{}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// ScrollHeight return always 0 for a span element
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollHeight
func (s *GOMLSpanElement) ScrollHeight() int {
	return 0
}

// ScrollLeft return always 0 for a span element
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (s *GOMLSpanElement) ScrollLeft() int {
	return 0
}

// ScrollTop return always 0 for a span element
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
func (s *GOMLSpanElement) ScrollTop() int {
	return 0
}

// ScrollWidth return always 0 for a span element
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (s *GOMLSpanElement) ScrollWidth() int {
	return 0
}

// TagName returns the tag name of the element on which
// it's called.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (s *GOMLSpanElement) TagName() string {
	return "span"
}
