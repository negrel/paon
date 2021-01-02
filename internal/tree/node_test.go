package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	node := NewNode()
	assert.NotNil(t, node)
}

func TestNode_isDescendantOf(t *testing.T) {
	parent := NewParent()
	node := NewNode()

	err := parent.AppendChildNode(node)
	assert.Nil(t, err)
	assert.True(t, node.IsDescendantOf(parent))
	err = parent.RemoveChildNode(node)
	assert.Nil(t, err)
	assert.False(t, node.IsDescendantOf(parent))

	assert.False(t, NewNode().IsDescendantOf(parent))
}

func TestNode_isDescendantOf_NilParent(t *testing.T) {
	node := NewNode()
	assert.False(t, node.IsDescendantOf(nil))
}
