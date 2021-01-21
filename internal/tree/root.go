package tree

import (
	"github.com/negrel/debuggo/pkg/assert"
)

type Root interface {
	ParentNode
}

var _ Root = &root{}

type root struct {
	*parentNode
}

func NewRoot(children Node) Root {
	assert.NotNil(children, "node must be non-nil to be the root")

	r := newRoot()
	r.parentNode.setParentNode(r)
	r.appendChildNode(children)

	return r
}

func newRoot() *root {
	return &root{
		parentNode: newParent(),
	}
}

func (r *root) RootNode() Root {
	return r
}
