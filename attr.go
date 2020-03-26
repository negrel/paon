package gom

// Attr represents one of a DOM element's attributes
// as an object.
// https://developer.mozilla.org/en-US/docs/Web/API/Attr
// https://dom.spec.whatwg.org/#Attr
type Attr struct {
	*Node
	ownerElement *Element
	value        string
}

func createAttribute(name string) *Attr {
	return &Attr{
		Node: &Node{
			nodeType: AttributeNode,
			nodeName: name,
		},
		ownerElement: nil,
		value:        "",
	}
}

/*****************************************************
 **************** Embedded interface *****************
 *****************************************************/
// ANCHOR Embedded interface

/* Node */
/* - Props */

// NodeName return attribute name
func (a *Attr) NodeName() string {
	return a.Name()
}

// NodeType return the "AttributeNode" type.
func (a *Attr) NodeType() NodeType {
	return AttributeNode
}

/* - Methods */

// CloneNode return a clone of the Attr
func (a *Attr) CloneNode(_ bool) *Attr {
	clone := createAttribute(a.nodeName)

	clone.SetValue(a.value)

	return clone
}

// IsEqualNode return wether or not two Attr are equal
func (a *Attr) IsEqualNode(other *Attr) bool {
	if other == nil {
		return false
	}

	// Checking NodeType
	if a.NodeType() != other.NodeType() {
		return false

	}

	// Checking name
	if a.Name() != other.Name() {
		return false
	}

	// Checking value
	if a.value != other.value {
		return false
	}

	return true
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Name return the Attribute name of an element
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/localName
func (a *Attr) Name() string {
	return a.nodeName
}

// OwnerElement return the element holding the Attribute
func (a *Attr) OwnerElement() *Element {
	return a.ownerElement
}

// Value return the Attribute value of an element
func (a *Attr) Value() string {
	return a.value
}

// SetValue set the Attribute value of an element
func (a *Attr) SetValue(value string) {
	a.value = value
}
