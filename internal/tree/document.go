package tree

var _ ContainerNode = Document{}

type Document struct {
	*containerNode
}

func NewDocument() *Document {
	var doc *Document
	doc = &Document{
		containerNode: &containerNode{
			Node: &node{
				nType: DocumentNode,
				owner: &doc,
			},
		},
	}

	return doc
}
