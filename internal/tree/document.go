package tree

import (
	"github.com/google/uuid"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

var _ ParentNode = Document{}

type Document struct {
	*parentNode

	elements map[uuid.UUID]Element
}

func NewDocument() *Document {
	var doc *Document
	doc = &Document{
		parentNode: &parentNode{
			Node: &node{
				nodeType: DocumentNodeType,
				// Document own itself so when it adopt node
				// it sets node owner to Document
				owner: &doc,
			},
		},
		elements: make(map[uuid.UUID]Element),
	}

	return doc
}

func (d *Document) registerElement(element Element) {
	id := element.ID()
	_, alreadyExist := d.elements[id]

	assert.False(alreadyExist, "can't register element", element, ", element is already registered")
	log.Infoln("registering element", element)

	d.elements[id] = element
}
