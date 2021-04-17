package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWidget_New(t *testing.T) {
	node := newTestWidget()
	assert.NotNil(t, node)
}

func TestWidget_isDescendantOf_Parent(t *testing.T) {
	parent := newTestLayout()
	node := newTestWidget()

	err := parent.AppendChild(node)
	assert.Nil(t, err)
	assert.True(t, node.IsDescendantOf(parent))
}

func TestWidget_IsDescendantOf_NonChild(t *testing.T) {
	parent := newTestLayout()

	_ = parent.AppendChild(newTestWidget())

	assert.False(t, newTestWidget().IsDescendantOf(parent))
}

func TestWidget_IsDescendantOf_RemovedChild(t *testing.T) {
	parent := newTestLayout()
	node := newTestWidget()

	_ = parent.AppendChild(node)

	err := parent.RemoveChild(node)
	assert.Nil(t, err)
	assert.False(t, node.IsDescendantOf(parent))
}

func TestWidget_IsDescendantOf_GreatParent(t *testing.T) {
	greatParent := newTestLayout()
	parent := newTestLayout()
	node := newTestWidget()

	_ = parent.AppendChild(node)
	_ = greatParent.AppendChild(parent)

	assert.True(t, node.IsDescendantOf(greatParent))
	assert.True(t, node.IsDescendantOf(parent))
}

func TestWidget_isDescendantOf_NilParent(t *testing.T) {
	node := newTestWidget()
	assert.False(t, node.IsDescendantOf(nil))
}

func TestWidget_Root(t *testing.T) {
	parent := newTestLayout()
	assert.Nil(t, parent.Root())

	root := newTestRoot()
	_ = root.AppendChild(parent)
	assert.Equal(t, root, parent.Root())

	child := newTestWidget()
	_ = parent.AppendChild(child)
	assert.Equal(t, root, child.Root())

	_ = parent.RemoveChild(child)
	assert.Nil(t, child.Root())
}
