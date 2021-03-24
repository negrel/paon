package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	node := NewNode()
	assert.NotNil(t, node)
}

func TestNode_isDescendantOf_Parent(t *testing.T) {
	parent := newTestParent()
	node := NewNode()

	err := parent.AppendChildNode(node)
	assert.Nil(t, err)
	assert.True(t, node.IsDescendantOf(parent))
}

func TestNode_IsDescendantOf_NonChildNode(t *testing.T) {
	parent := newTestParent()

	_ = parent.AppendChildNode(newNode())

	assert.False(t, newNode().IsDescendantOf(parent))
}

func TestNode_IsDescendantOf_RemovedChild(t *testing.T) {
	parent := newTestParent()
	node := NewNode()

	_ = parent.AppendChildNode(node)

	err := parent.RemoveChildNode(node)
	assert.Nil(t, err)
	assert.False(t, node.IsDescendantOf(parent))
}

func TestNode_IsDescendantOf_GreatParent(t *testing.T) {
	greatParent := newTestParent()
	parent := newTestParent()
	node := NewNode()

	_ = parent.AppendChildNode(node)
	_ = greatParent.AppendChildNode(parent)

	assert.True(t, node.IsDescendantOf(greatParent))
	assert.True(t, node.IsDescendantOf(parent))
}

func TestNode_isDescendantOf_NilParent(t *testing.T) {
	node := newNode()
	assert.False(t, node.IsDescendantOf(nil))
}

type testRoot testParent

func newTestRoot() *testRoot {
	return &testRoot{
		parentNode: newParent(),
	}
}

func (tr *testRoot) RootNode() ParentNode {
	return tr
}

func (tr *testRoot) AppendChildNode(newChild Node) (err error) {
	err = tr.parentNode.AppendChildNode(newChild)
	if err == nil {
		SetParentOf(newChild, tr)
	}

	return err
}

func TestNode_Root(t *testing.T) {
	parent := newTestParent()
	assert.Nil(t, parent.RootNode())

	root := newTestRoot()
	_ = root.AppendChildNode(parent)
	assert.Equal(t, root, parent.RootNode())

	child := newNode()
	_ = parent.AppendChildNode(child)
	assert.Equal(t, root, child.RootNode())

	_ = parent.RemoveChildNode(child)
	assert.Nil(t, child.RootNode())
}
