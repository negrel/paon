package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayout_New(t *testing.T) {
	parent := newTestLayout()
	assert.NotNil(t, parent)
}

func TestLayout_AppendChild_EmptyParent_Success(t *testing.T) {
	parent := newTestLayout()
	child := newTestWidget()

	// Appending it
	err := parent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.FirstChild())
	assert.Equal(t, child, parent.LastChild())
	assert.Equal(t, parent, child.parent)
}

func TestLayout_AppendChild_NonEmptyParent_Success(t *testing.T) {
	parent := newTestLayout()
	child := newTestWidget()
	newChild := newTestWidget()

	// Appending it
	_ = parent.AppendChild(child)

	err := parent.AppendChild(newChild)
	assert.Nil(t, err)
	assert.Equal(t, newChild, parent.LastChild())
	assert.Equal(t, parent.FirstChild(), newChild.prevWidget)
	assert.Equal(t, parent.FirstChild(), child)
	assert.Equal(t, parent.FirstChild().Next(), newChild)
	assert.Equal(t, parent, child.parent)
}

func TestLayout_AppendChild_AncestorOfParent(t *testing.T) {
	parent := newTestLayout()
	ancestor := newTestLayout()

	_ = ancestor.AppendChild(parent)

	// should return an error, ancestor is the parentNode
	err := parent.AppendChild(ancestor)
	assert.NotNil(t, err)
}

func TestLayout_AppendChild_NilChild(t *testing.T) {
	parent := newTestLayout()

	assert.Panics(t, func() {
		_ = parent.AppendChild(nil)
	}, "appending a nil child should panic. (\"assert\" build tag must be enabled)")
}

func TestLayout_AppendChild_WithParent(t *testing.T) {
	parent := newTestLayout()
	child1 := newTestWidget()
	child2 := newTestWidget()
	_ = parent.AppendChild(child1)
	_ = parent.AppendChild(child2)

	otherParent := newTestLayout()
	node := newTestWidget()
	_ = otherParent.AppendChild(node)

	err := parent.AppendChild(node)
	assert.Nil(t, err)
	assert.Equal(t, parent, node.parent)
	assert.Equal(t, child2, node.prevWidget)
	assert.Equal(t, nil, node.nextWidget)
}

func TestLayout_AppendChild_Itself(t *testing.T) {
	parent := newTestLayout()

	err := parent.AppendChild(parent)
	assert.NotNil(t, err)
}

func TestLayout_InsertBefore_EmptyParent_Success(t *testing.T) {
	parent := newTestLayout()
	child := newTestWidget()

	err := parent.InsertBefore(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, parent.FirstChild(), child)
	assert.Equal(t, parent.LastChild(), child)
	assert.Equal(t, child.parent, parent)
	assert.Equal(t, nil, child.nextWidget)
	assert.Equal(t, nil, child.prevWidget)
}

func TestLayout_InsertBefore_NonEmptyParent_Success(t *testing.T) {
	parent := newTestLayout()
	prev := newTestWidget()
	next := newTestWidget()
	child := newTestWidget()

	_ = parent.AppendChild(prev)
	_ = parent.AppendChild(next)
	err := parent.InsertBefore(next, child)
	assert.Nil(t, err)
	assert.Equal(t, prev, parent.FirstChild())
	assert.Equal(t, next, parent.LastChild())
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, child, next.prevWidget)
	assert.Equal(t, next, child.nextWidget)
	assert.Equal(t, child, prev.nextWidget)
	assert.Equal(t, prev, child.prevWidget)
}

func TestLayout_InsertBefore_NilChild(t *testing.T) {
	parent := newTestLayout()

	reference := newTestWidget()
	_ = parent.AppendChild(reference)

	// Nil child
	assert.Panics(t, func() {
		_ = parent.InsertBefore(reference, nil)
	}, "inserting a nil child should panic. (\"assert\" build tag must be enabled)")
}

func TestLayout_InsertBefore_EmptyParent_NilReference(t *testing.T) {
	// Inserting before a nil reference is same as appending

	parent := newTestLayout()
	child := newTestWidget()

	// Appending it
	err := parent.InsertBefore(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.FirstChild())
	assert.Equal(t, child, parent.LastChild())
}

func TestLayout_InsertBefore_NonEmptyParent_NilReference(t *testing.T) {
	parent := newTestLayout()
	child := newTestWidget()
	_ = parent.InsertBefore(nil, child)

	// appending another child
	otherChild := newTestWidget()
	err := parent.InsertBefore(nil, otherChild)
	assert.Nil(t, err)
	assert.Equal(t, parent.LastChild(), otherChild)
	assert.Equal(t, child, otherChild.prevWidget)
	assert.Equal(t, child.nextWidget, otherChild)
}

func TestLayout_InsertBefore_AncestorOfParent(t *testing.T) {
	// Setting up the parent node
	parent := newTestLayout()
	reference := newTestLayout()
	_ = parent.AppendChild(reference)

	// Inserting the parent in the ancestor node
	ancestor := newTestLayout()
	_ = ancestor.InsertBefore(nil, parent)

	// Trying to insert the parent node in the ancestor node
	err := parent.InsertBefore(reference, ancestor)
	assert.NotNil(t, err)
}

func TestLayout_InsertBefore_ChildWithParent(t *testing.T) {
	// Setting up the parent
	parent := newTestLayout()
	prev := newTestWidget()
	next := newTestWidget()
	_ = parent.AppendChild(next)
	_ = parent.InsertBefore(next, prev)

	// The initial parent of the child
	initialParent := newTestLayout()
	reference := newTestWidget()
	_ = initialParent.AppendChild(reference)

	child := newTestWidget()
	_ = initialParent.InsertBefore(reference, child)

	// Inserting the child to the new parent
	err := parent.InsertBefore(next, child)
	assert.Nil(t, err)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, next, child.nextWidget)
	assert.Equal(t, prev, child.prevWidget)
}

func TestLayout_InsertBefore_Itself(t *testing.T) {
	parent := newTestLayout()

	err := parent.InsertBefore(nil, parent)
	assert.NotNil(t, err)
}

func TestLayout_isAncestorOf_NilChild(t *testing.T) {
	parent := newTestLayout()
	assert.False(t, parent.IsAncestorOf(nil))
}

func TestLayout_RemoveChild_Success(t *testing.T) {
	parent := newTestLayout()

	node := newTestWidget()
	_ = parent.AppendChild(node)

	child := newTestWidget()
	_ = parent.AppendChild(child)

	err := parent.RemoveChild(node)
	assert.Nil(t, err)
	assert.Equal(t, nil, node.Parent())
	assert.Equal(t, nil, node.Next())
	assert.Equal(t, child.Previous(), nil)
	assert.Equal(t, parent.FirstChild(), child)
	assert.Equal(t, parent.LastChild(), child)
}

func TestLayout_RemoveChild_NilChild(t *testing.T) {
	parent := newTestLayout()

	// Remove child from parent with no child
	assert.Panics(t, func() {
		_ = parent.RemoveChild(nil)
	})

	node := newTestWidget()
	_ = parent.AppendChild(node)

	// Parent with one child
	assert.Panics(t, func() {
		_ = parent.RemoveChild(nil)
	})
}

func TestLayout_RemoveChild_ChildOfAnotherParent(t *testing.T) {
	node := newTestLayout()
	parent := newTestLayout()
	child := newTestWidget()

	_ = parent.AppendChild(child)

	err := node.RemoveChild(child)
	assert.NotNil(t, err)
}
