package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	node := NewNode("test_node")
	assert.NotNil(t, node)
}

func TestNode_isDescendantOf(t *testing.T) {
	parent := newParent("test_parent")
	node := NewNode("test_child")

	err := parent.AppendChildNode(node)
	assert.Nil(t, err)
	assert.True(t, node.IsDescendantOf(parent))
	err = parent.RemoveChildNode(node)
	assert.Nil(t, err)
	assert.False(t, node.IsDescendantOf(parent))

	assert.False(t, NewNode("test_node").IsDescendantOf(parent))
}

func TestNode_isDescendantOf_NilParent(t *testing.T) {
	node := NewNode("test_child")
	assert.False(t, node.IsDescendantOf(nil))
}
