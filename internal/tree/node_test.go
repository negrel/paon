package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	node := newNode()
	assert.NotNil(t, node)
}

func TestNode_isDescendantOf(t *testing.T) {
	parent := NewParent()
	node := newNode()

	err := parent.AppendChildNode(node)
	assert.Nil(t, err)
	assert.True(t, node.IsDescendantOf(parent))
	err = parent.RemoveChildNode(node)
	assert.Nil(t, err)
	assert.False(t, node.IsDescendantOf(parent))

	assert.False(t, newNode().IsDescendantOf(parent))
}

func TestNode_isDescendantOf_NilParent(t *testing.T) {
	node := newNode()
	assert.False(t, node.IsDescendantOf(nil))
}

func TestNode_Root(t *testing.T) {
	parent := newParent()
	assert.Nil(t, parent.RootNode())

	root := NewRoot(parent)
	assert.Equal(t, root, parent.RootNode())

	child := newNode()
	_ = parent.AppendChildNode(child)
	assert.Equal(t, root, child.RootNode())

	_ = parent.RemoveChildNode(child)
	assert.Nil(t, child.RootNode())
}
