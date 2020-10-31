package tree

import (
	"fmt"
)

var _ ChildNode = &TextNode{}

// TextNode is a node
type TextNode struct {
	ChildNode

	content string
}

func newTextNode(data string) TextNode {
	return TextNode{
		ChildNode: makeChildNode(&node{
			nodeType: TextNodeType,
		}),
		content: data,
	}
}

func (t TextNode) String() string {
	return fmt.Sprintf("TextNode(%v)", t.content)
}
