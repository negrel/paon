package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParent_New(t *testing.T) {
	parent := newParent()
	assert.NotNil(t, parent)
}

func TestParent_AppendChild_Success(t *testing.T) {
	parent := newParent()
	child := newNode()

	// Appending it
	err := parent.AppendChildNode(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.firstChild)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent, child.parent)

	// Appending another child
	child = newNode()
	err = parent.AppendChildNode(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous)
	assert.Equal(t, parent.firstChild.NextNode(), child)
	assert.Equal(t, parent, child.parent)

	// Re appending the last child
	err = parent.AppendChildNode(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous) // Still only two elements
	assert.Equal(t, parent.firstChild.NextNode(), child)
	assert.Equal(t, parent, child.parent)
}

func TestParent_AppendChild_AncestorOfParent(t *testing.T) {
	parent := newParent()
	fakeChild := newParent()

	err := fakeChild.AppendChildNode(parent)
	assert.Nil(t, err)

	// should return an error, fakeChild is the parentNode
	err = parent.AppendChildNode(fakeChild)
	assert.NotNil(t, err)

	// Try again with ParentNode of fake child
	fakeChildParent := newParent()
	err = fakeChildParent.AppendChildNode(fakeChild)
	assert.Nil(t, err)

	err = parent.AppendChildNode(fakeChildParent)
	assert.NotNil(t, err)
}

func TestParent_AppendChild_NilChild(t *testing.T) {
	parent := newParent()

	assert.Panics(t, func() {
		_ = parent.AppendChildNode(nil)
	}, "appending a nil child should panic. (\"assert\" build tag must be enabled)")
}

func TestParent_AppendChild_WithParent(t *testing.T) {
	parent := newParent()
	prev := newNode()
	next := newNode()
	_ = parent.AppendChildNode(prev)
	_ = parent.AppendChildNode(next)
	child := newNode()
	_ = parent.AppendChildNode(child)

	otherParent := newParent()
	err := otherParent.AppendChildNode(child)
	assert.Nil(t, err)
	assert.Equal(t, otherParent, child.parent)
	assert.Nil(t, child.previous)
	assert.Nil(t, child.next)

	err = parent.AppendChildNode(child)
	assert.Nil(t, err)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, next, child.previous)
}

func TestParent_InsertBefore_Success(t *testing.T) {
	parent := newParent()
	prev := newNode()
	next := newNode()
	child := newNode()

	_ = parent.AppendChildNode(prev)
	_ = parent.AppendChildNode(next)
	err := parent.InsertBeforeNode(next, child)
	assert.Nil(t, err)
	assert.Equal(t, prev, parent.firstChild)
	assert.Equal(t, next, parent.lastChild)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, child, next.previous)
	assert.Equal(t, next, child.next)
	assert.Equal(t, child, prev.next)
	assert.Equal(t, prev, child.previous)
}

func TestParent_InsertBefore_NilChild(t *testing.T) {
	parent := newParent()

	reference := newNode()
	err := parent.AppendChildNode(reference)
	assert.Nil(t, err)

	// Nil child
	assert.Panics(t, func() {
		_ = parent.InsertBeforeNode(reference, nil)
	}, "inserting a nil child should panic. (\"assert\" build tag must be enabled)")
}

func TestParent_InsertBefore_NilReference(t *testing.T) {
	// Inserting before a nil reference is same as appending

	parent := newParent()
	child := newNode()

	// Appending it
	err := parent.InsertBeforeNode(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.firstChild)
	assert.Equal(t, child, parent.lastChild)

	// Appending another child
	child = newNode()
	err = parent.InsertBeforeNode(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous)
	assert.Equal(t, parent.firstChild.NextNode(), child)

	// Re appending the last child
	err = parent.InsertBeforeNode(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous) // Still only two elements
	assert.Equal(t, parent.firstChild.NextNode(), child)
}

func TestParent_InsertBefore_AncestorOfParent(t *testing.T) {
	parent := newParent()
	reference := newParent()
	_ = parent.AppendChildNode(reference)

	fakeChild := newParent()
	_ = fakeChild.InsertBeforeNode(nil, parent)

	err := parent.InsertBeforeNode(reference, fakeChild)
	assert.NotNil(t, err)

	// Try again with ParentNode of fake child
	fakeChildParent := newParent()
	err = fakeChildParent.AppendChildNode(fakeChild)
	assert.Nil(t, err)

	err = parent.InsertBeforeNode(reference, fakeChildParent)
	assert.NotNil(t, err)
}

func TestParent_InsertBefore_ChildWithParent(t *testing.T) {
	parent := newParent()
	prev := newNode()
	next := newNode()
	_ = parent.AppendChildNode(next)
	_ = parent.InsertBeforeNode(next, prev)
	child := newNode()
	_ = parent.InsertBeforeNode(next, child)

	otherParent := newParent()
	reference := newNode()
	_ = otherParent.AppendChildNode(reference)
	err := otherParent.InsertBeforeNode(reference, child)
	assert.Nil(t, err)
	assert.Equal(t, otherParent, child.parent)
	assert.Nil(t, child.previous)
	assert.Equal(t, reference, child.next)

	assert.Equal(t, prev.next, next)
	assert.Equal(t, next.previous, prev)

	err = parent.InsertBeforeNode(next, child)
	assert.Nil(t, err)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, next, child.next)
	assert.Equal(t, prev, child.previous)
}

func TestParent_isAncestorOf_NilChild(t *testing.T) {
	parent := newParent()
	assert.False(t, parent.IsAncestorOf(nil))
}

func TestParentNode_RemoveChild_Success(t *testing.T) {
	parent := newParent()

	node := newNode()
	_ = parent.AppendChildNode(node)

	child := newNode()
	_ = parent.AppendChildNode(child)

	err := parent.RemoveChildNode(node)
	assert.Nil(t, err)
	assert.Equal(t, nil, node.ParentNode())
	assert.Equal(t, nil, node.NextNode())
	assert.Equal(t, child.PreviousNode(), nil)
	assert.Equal(t, parent.FirstChildNode(), child)
	assert.Equal(t, parent.LastChildNode(), child)
}

func TestParentNode_RemoveChild_NilChild(t *testing.T) {
	parent := newParent()

	// Remove child from parent with no child
	assert.Panics(t, func() {
		_ = parent.RemoveChildNode(nil)
	})

	node := newNode()
	_ = parent.AppendChildNode(node)

	// Parent with one child
	assert.Panics(t, func() {
		_ = parent.RemoveChildNode(nil)
	})
}

func TestParentNode_RemoveChild_ChildOfAnotherParent(t *testing.T) {

}
