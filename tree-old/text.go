package tree

import (
	e "github.com/negrel/gom/exception"
)

/* NOTE Element missing props & methods (OFFICIAL DOM) :
 * ** Props **
 * assignedSlot
 * all obsolete or non-standardized props
 *
 * ** Methods **
 * all obsolete or non-standardized methods
 */

// Text interface represents the textual content
// of Element or Attr.
// https://developer.mozilla.org/en-US/docs/Web/API/Text
// https://dom.spec.whatwg.org/#interface-text
type Text struct {
	*CharacterData
}

// createTextNode return a new Text node.
func createTextNode(content string) *Text {
	return &Text{
		&CharacterData{
			data: []rune(content),
		},
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// WholeText return the text of the node.
func (t *Text) WholeText() string {
	return t.Data()
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// SplitText method breaks the Text node into two nodes
// at the specified offset, keeping both nodes in the
// tree as siblings.
// https://developer.mozilla.org/en-US/docs/Web/API/Text/splitText
// https://dom.spec.whatwg.org/#dom-text-splittext
func (t *Text) SplitText(offset int) (*Text, e.Exception) {
	var length int = t.Length()
	var count int = length - offset

	// If offset is greater than length
	if count < 0 {
		return nil, e.RangeError("The offset %v is larger than the Text node's length.", string(offset))
	}

	// Substring is data of the new node
	newTextData, err := t.SubstringData(offset, count)

	if err != nil {
		return nil, err
	}

	// Creating new node
	newText := createTextNode(newTextData)
	newText.setOwnerDocument(t.OwnerDocument())

	if parent := t.ParentNode(); parent != nil {
		parent.InsertBefore(newText.Node, t.NextSibling())

	}

	// Deleting data of the current Text node
	t.DeleteData(offset, count)

	return newText, nil
}
