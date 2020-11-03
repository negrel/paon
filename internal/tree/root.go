package tree

import (
	"github.com/google/uuid"
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

type Root interface {
	Parent

	register(Node)
	unregister(Node)
}

var _ Root = &root{}

type root struct {
	*parent

	widgets map[uuid.UUID]Node
}

func newRoot() *root {
	return &root{
		parent:  newParent("root"),
		widgets: make(map[uuid.UUID]Node),
	}
}

func (r *root) register(child Node) {
	assert.NotNil(child, "can't register a non-nil child")

	id := child.ID()
	_, isAlreadyRegistered := r.widgets[id]
	if isAlreadyRegistered {
		log.Infoln(child, "is already registered")
	}

	log.Infoln("registering", child)
	r.widgets[id] = child
}

func (r *root) unregister(child Node) {
	assert.NotNil(child, "can't unregister a non-nil child")

	id := child.ID()
	_, isRegistered := r.widgets[id]
	if !isRegistered {
		log.Infoln(child, "is not registered")
	}

	delete(r.widgets, id)
}
