package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParent_New(t *testing.T) {
	parent := newParent("test_parent")
	assert.NotNil(t, parent)
	assert.Equal(t, "test_parent", parent.Name())
}

func TestParent_AppendChild_Success(t *testing.T) {
	parent := newParent("test_parent")
	child := newNode("test_child")

	// Appending it
	err := parent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.firstChild)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent, child.parent)

	// Appending another child
	child = newNode("second_test_child")
	err = parent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous)
	assert.Equal(t, parent.firstChild.Next(), child)
	assert.Equal(t, parent, child.parent)

	// Re appending the last child
	err = parent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous) // Still only two elements
	assert.Equal(t, parent.firstChild.Next(), child)
	assert.Equal(t, parent, child.parent)
}

func TestParent_AppendChild_AncestorOfParent(t *testing.T) {
	parent := newParent("test_parent")
	fakeChild := newParent("test_fake_child")

	err := fakeChild.AppendChild(parent)
	assert.Nil(t, err)

	// should return an error, fakeChild is the parent
	err = parent.AppendChild(fakeChild)
	assert.NotNil(t, err)

	// Try again with Parent of fake child
	fakeChildParent := newParent("test_fake_child_parent")
	err = fakeChildParent.AppendChild(fakeChild)
	assert.Nil(t, err)

	err = parent.AppendChild(fakeChildParent)
	assert.NotNil(t, err)
}

func TestParent_AppendChild_NilChild(t *testing.T) {
	parent := newParent("test_parent")

	assert.Panics(t, func() {
		_ = parent.AppendChild(nil)
	}, "appending a nil child should panic. (\"assert\" build tag must be enabled)")
}

func TestParent_InsertBefore_Success(t *testing.T) {
	parent := newParent("test_parent")
	reference := newNode("test_reference")
	child := newNode("test_child")

	_ = parent.AppendChild(reference)
	err := parent.InsertBefore(reference, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.firstChild)
	assert.Equal(t, reference, parent.lastChild)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, child, reference.previous)
	assert.Equal(t, reference, child.next)
}

func TestParent_InsertBefore_NilChild(t *testing.T) {
	parent := newParent("test_parent")

	reference := newNode("test_reference")
	err := parent.AppendChild(reference)
	assert.Nil(t, err)

	// Nil child
	assert.Panics(t, func() {
		_ = parent.InsertBefore(reference, nil)
	}, "inserting a nil child should panic. (\"assert\" build tag must be enabled)")
}

func TestParent_InsertBefore_NilReference(t *testing.T) {
	// Inserting before a nil reference is same as appending

	parent := newParent("test_parent")
	child := newNode("test_child")

	// Appending it
	err := parent.InsertBefore(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.firstChild)
	assert.Equal(t, child, parent.lastChild)

	// Appending another child
	child = newNode("second_test_child")
	err = parent.InsertBefore(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous)
	assert.Equal(t, parent.firstChild.Next(), child)

	// Re appending the last child
	err = parent.InsertBefore(nil, child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous) // Still only two elements
	assert.Equal(t, parent.firstChild.Next(), child)
}
