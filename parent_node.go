package gom

// ParentNode mixin contains methods and properties
// that are common to all types of Node objects that can have children
// https://developer.mozilla.org/en-US/docs/Web/API/ParentNode
// https://dom.spec.whatwg.org/#parentnode
type ParentNode struct {
	children GOMLCollection
}

func newParentNode() *ParentNode {
	return &ParentNode{
		children: newGOMLCollection(),
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// FirstElementChild returns the object's first child
// Element, or null if there are no child elements.
// https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/firstElementChild
func (pn *ParentNode) FirstElementChild() Element {
	return pn.children.Item(0)
}

// LastElementChild returns the object's last child
// Element, or null if there are no child elements.
// https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/lastElementChild
func (pn *ParentNode) LastElementChild() Element {
	var lastIndex = pn.children.Length() - 1

	return pn.children.Item(lastIndex)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Append inserts a set of Node objects or DOMString
// objects after the last child of the ParentNode
// https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/append
func (pn *ParentNode) Append() {
	// TODO func (pn *ParentNode) Append(nodes ...Node)
	// https://dom.spec.whatwg.org/#dom-parentnode-append
}

// Prepend inserts a set of Node objects or DOMString
// objects before the first child of the ParentNode
// https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/prepend
func (pn *ParentNode) Prepend(nodes ...Node) {
	// TODO func (pn *ParentNode) Prepend(nodes ...Node)
	// https://dom.spec.whatwg.org/#dom-parentnode-prepend
}

// QuerySelector returns the first Element within
// the document that matches the specified selector,
// or group of selectors. If no matches are found,
// null is returned.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (pn *ParentNode) QuerySelector(selector string) {
	// TODO func (d *Document) QuerySelector(selector string)
}

// QuerySelectorAll returns a static (not live) NodeList
// representing a list of the document's elements that
// match the specified group of selectors.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
func (pn *ParentNode) QuerySelectorAll(selector string) {
	// TODO func (d *Document) QuerySelectorAll(selector string)
}
