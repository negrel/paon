package tree

import "golang.org/x/text/encoding"

/* NOTE Document missing props & methods (OFFICIAL DOM) :
 * ** Props **
 * characterSet
 * compatMode (experimental api)
 * contentType (experimental api)
 * documentURI
 * embeds
 * fonts
 * forms
 * images
 * implementation
 * lastStyleSheetSet
 * links
 * plugins
 * featurePolicy (experimental api)
 * preferredStyleSheetSet
 * scripts
 * scrollingElement
 * SelectedStyleSheetSet
 * styleSheetSets
 * timeline
 * all obsolete or non-standardized props
 *
 * ** Methods **
 * caretRangeFromPoint
 * createAttributeNS
 * createCDATASection
 * createElementNS
 * createEvent
 * createNodeIterator
 * createProcessingInstruction
 * createRange
 * createTouchList
 * createTreeWalker
 * enableStyleSheetsForSet
 * hasStorageAccess
 * requestStorageAccess
 * createExpression
 * createNSResolver
 * all obsolete or non-standardized methods
 */

// Document interface represents any page loaded
// and serves as an entry point into the page's
// content
// https://developer.mozilla.org/en-US/docs/Web/API/Document
// https://dom.spec.whatwg.org/#document
type Document struct {
	*Node
	Body            *Element
	characterSet    encoding.Encoding
	documentElement *Element
	head            *Element
	hidden          bool
	visibilityState string
}

// NewDocument return a new document object serving
// as an entry point into the page's content.
func NewDocument(name string) *Document {
	return &Document{
		Node: &Node{
			nodeType: DocumentNode,
			nodeName: "#document",
		},
		Body:            createElement("body"),
		documentElement: nil,
		head:            nil,
		hidden:          false,
		visibilityState: "visible",
	}
}

/*****************************************************
 **************** Embedded interface *****************
 *****************************************************/
// ANCHOR Embedded interface

/* Node */
/* - Props */

// NodeName return the GOML-uppercased name
func (d *Document) NodeName() string {
	return "#document"
}

// NodeType return the "ElementNode" type.
func (d *Document) NodeType() NodeType {
	return DocumentNode
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// CharacterSet return the current character set used by
// the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (d *Document) CharacterSet() encoding.Encoding {
	return d.characterSet
}

// DocumentElement returns the Element that is the root
// element of the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (d *Document) DocumentElement() *Element {
	return d.documentElement
}

// Head return the <head> element of the current document
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (d *Document) Head() *Element {
	return d.head
}

// Hidden returns a Boolean value indicating if the page
// is considered hidden or not.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (d *Document) Hidden() bool {
	return d.hidden
}

// SetBody set the body node of the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (d *Document) SetBody(body *Element) {
	d.Body = body
}

// SetCharacterSet method set the document character
// set (UTF-8).
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (d *Document) SetCharacterSet(charSet encoding.Encoding) {
	d.characterSet = charSet
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// AdoptNode transfers a node from another document
// into the document on which the method was called.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/adoptNode
func (d *Document) AdoptNode(external *Node) {
	// If external node have a parent
	if extParent := external.ParentNode(); extParent != nil {
		// Removing the child from the parent
		external, _ = extParent.RemoveChild(external)
	}

	// Adopting the node
	d.AppendChild(external)
	// Changing ownerDocument of the child and subchild...
	d.apply(func(node *Node) {
		node.setOwnerDocument(d)
	})
}

// CreateAttribute method creates a new attribute node,
// and returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createAttribute
func (d *Document) CreateAttribute(name string) *Attr {
	return createAttribute(name)
}

// CreateComment creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createComment
func (d *Document) CreateComment(data string) *Comment {
	// TODO func (d *Document) CreateComment() Node
	return createComment(data)
}

// CreateElement creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (d *Document) CreateElement(tagName string) *Element {
	return createElement(tagName)
}

// CreateTextNode creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode
func (d *Document) CreateTextNode(content string) *Text {
	return createTextNode(content)
}

// GetElementsByClassName method of Document interface
// returns an array-like object of all child elements
// which have all of the given class names.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
func (d *Document) GetElementsByClassName(className string) *Element {
	// TODO func (d *Document) GetElementsByClassName(className string) Element
	return createElement("")
}

// GetElementsByTagName method of Document interface
// returns an array-like object of all child elements
// which have all of the given tag names.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (d *Document) GetElementsByTagName(tagName string) *Element {
	// TODO func (d *Document) GetElementsByTagName(tagName string) Element
	return createElement(tagName)
}

// ImportNode method creates a copy of a Node or
// DocumentFragment from another document, to be
// inserted into the current document later.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/importNode
func (d *Document) ImportNode(node *Node, deep bool) *Node {
	// TODO func (d *Document) ImportNode(node Node, deep bool) Node
	return node
}

// // GetElementById returns an Element object representing
// // the element whose id property matches the specified string
// // https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
// func (d *Document) GetElementById(id string) *Element {
// 	// TODO func (d *Document) GetElementById(id string) Element
// 	return createElement("")
// }

// // QuerySelector returns the first Element within the document
// // that matches the specified selector, or group of selectors.
// // https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
// func (d *Document) QuerySelector(selector string) Element {
// 	// TODO func (d *Document) QuerySelector(selector string) Element
// }

// // QuerySelectorAll returns a static (not live) NodeList
// // representing a list of the document's elements that match
// // the specified group of selectors.
// // https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
// func (d *Document) QuerySelectorAll(selector string) NodeList {
// 	// TODO func (d *Document) QuerySelectorAll(selector string) NodeList
// }
