package tree

import (
	"log"
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

	// should return an error, fakeChild is the parentNode
	err = parent.AppendChild(fakeChild)
	assert.NotNil(t, err)

	// Try again with ParentNode of fake child
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

func TestParent_AppendChild_WithParent(t *testing.T) {
	parent := newParent("test_parent")
	prev := newNode("test_prev_sibling")
	next := newNode("test_next_sibling")
	_ = parent.AppendChild(prev)
	_ = parent.AppendChild(next)
	child := newNode("test_child")
	_ = parent.AppendChild(child)

	otherParent := newParent("test_other_parent")
	err := otherParent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, otherParent, child.parent)
	assert.Nil(t, child.previous)
	assert.Nil(t, child.next)

	err = parent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, next, child.previous)
}

func TestParent_InsertBefore_Success(t *testing.T) {
	parent := newParent("test_parent")
	prev := newNode("test_prev")
	next := newNode("test_next")
	child := newNode("test_child")

	_ = parent.AppendChild(prev)
	_ = parent.AppendChild(next)
	err := parent.InsertBefore(next, child)
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

func TestParent_InsertBefore_AncestorOfParent(t *testing.T) {
	parent := newParent("test_parent")
	reference := newParent("test_child")
	_ = parent.AppendChild(reference)

	fakeChild := newParent("test_fake_child")
	_ = fakeChild.InsertBefore(nil, parent)

	err := parent.InsertBefore(reference, fakeChild)
	assert.NotNil(t, err)

	// Try again with ParentNode of fake child
	fakeChildParent := newParent("test_fake_child_parent")
	err = fakeChildParent.AppendChild(fakeChild)
	assert.Nil(t, err)

	err = parent.InsertBefore(reference, fakeChildParent)
	assert.NotNil(t, err)
}

func TestParent_InsertBefore_ChildWithParent(t *testing.T) {
	parent := newParent("test_parent")
	prev := newNode("test_prev_sibling")
	next := newNode("test_next_sibling")
	_ = parent.AppendChild(next)
	_ = parent.InsertBefore(next, prev)
	child := newNode("test_child")
	_ = parent.InsertBefore(next, child)

	otherParent := newParent("test_other_parent")
	reference := newNode("test_reference_child")
	_ = otherParent.AppendChild(reference)
	err := otherParent.InsertBefore(reference, child)
	assert.Nil(t, err)
	assert.Equal(t, otherParent, child.parent)
	assert.Nil(t, child.previous)
	assert.Equal(t, reference, child.next)

	log.Println(prev, next.previous)
	log.Println(prev.next, next)
	assert.Equal(t, prev.next, next)
	assert.Equal(t, next.previous, prev)

	err = parent.InsertBefore(next, child)
	assert.Nil(t, err)
	assert.Equal(t, parent, child.parent)
	assert.Equal(t, next, child.next)
	assert.Equal(t, prev, child.previous)
}
