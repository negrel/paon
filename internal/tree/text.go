package tree

var _ ChildNode = &Text{}

// Text is a node
type Text struct {
	ChildNode

	content string
}

func MakeTextNode(data string) Text {
	return Text{
		ChildNode: makeChildNode(&node{
			nType: TextNode,
		}),
		content: data,
	}
}
