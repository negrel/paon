package gom

// GOMLElement interface represents any GOML element
type GOMLElement interface {
	/* EMBEDDED INTERFACE */
	Element
	/* GETTERS & SETTERS (props) */
	Hidden() bool
	InnerText() string
	SetInnerText(string)
	Style() interface{}
	/* METHODS */
	Click()
}

var _ GOMLElement = &gomlElement{}
var _ Element = &gomlElement{}
var _ Node = &gomlElement{}

type gomlElement struct {
	*element
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Hidden return true if the element is hidden or not.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/hidden
func (e *gomlElement) Hidden() bool {
	// TODO (e *gomlElement) Hidden() bool
	return false
}

// InnerText represents the "rendered" text content of
// a node and its descendants.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (e *gomlElement) InnerText() string {
	// TODO (e *gomlElement) InnerText() string
	return ""
}

// SetInnerText set the inner text of a GOML element.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (e *gomlElement) SetInnerText(string) {
	// TODO (e *gomlElement) SetInnerText(string)
}

// Style return
func (e *gomlElement) Style() interface{} {
	// TODO (e *gomlElement) Style() interface{}
	return nil
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Click handle click event a GOML element.
func (e *gomlElement) Click() {}
