package tree

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/negrel/debuggo/pkg/log"

	"github.com/negrel/paon/internal/events"
	"github.com/negrel/paon/internal/style"
)

type Element interface {
	ParentNode
	events.EventTarget

	Name() string
	ID() uuid.UUID

	FirstElementChild() Element
	LastElementChild() Element

	Style() map[style.PropertyType]style.Property
}

var _ Element = &element{}

type element struct {
	ParentNode
	events.Target

	id   uuid.UUID
	name string
}

func newElement(name string) Element {
	log.Infoln("creating", name, "element")

	return element{
		ParentNode: newParentNode(ElementNodeType),
		id:         uuid.New(),
		name:       name,
	}
}

func (e element) Name() string {
	return e.name
}

func (e element) ID() uuid.UUID {
	return e.id
}

func (e element) String() string {
	return fmt.Sprintf("%v-%v", e.name, e.id)
}

func (e element) FirstElementChild() Element {
	var node Node = e.FirstChild()
	for node != nil {
		if el, isElement := node.(Element); isElement {
			return el
		}
		node = node.Previous()
	}

	return nil
}

func (e element) LastElementChild() Element {
	var node Node = e.LastChild()
	for node != nil {
		if el, isElement := node.(Element); isElement {
			return el
		}
		node = node.Next()
	}

	return nil
}

func (e element) Style() map[style.PropertyType]style.Property {
	return map[style.PropertyType]style.Property{}
}

func (e element) InsertBefore(reference, newChild Node) (cn ChildNode, err error) {
	log.Infoln("inserting", newChild, "before", reference, "in element", e)

	cn, err = e.ParentNode.InsertBefore(reference, newChild)

	e.registerChildToDoc(cn)

	return
}

func (e element) registerChildToDoc(child Node) {
	if doc := e.Owner(); doc != nil {
		if el, isElement := child.(Element); isElement {
			doc.registerElement(el)
		}
	}
}

func (e element) AppendChild(newChild Node) (cn ChildNode, err error) {
	log.Infoln("appending", newChild, "to ", e)

	cn, err = e.ParentNode.AppendChild(newChild)

	e.registerChildToDoc(cn)

	return
}
