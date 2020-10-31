package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParent_AppendChild_Success(t *testing.T) {
	parent := newParent("test_parent")
	assert.NotNil(t, parent)

	// Creating first child
	child := newNode("test_child")
	assert.NotNil(t, child)

	// Appending it
	err := parent.AppendChild(child)
	assert.Nil(t, err)
	assert.NotNil(t, parent.firstChild)
	assert.NotNil(t, parent.lastChild)
	assert.Equal(t, child, parent.firstChild)
	assert.Equal(t, child, parent.lastChild)

	// Appending another child
	child = newNode("second_test_child")
	err = parent.AppendChild(child)
	assert.Nil(t, err)
	assert.NotNil(t, parent.firstChild)
	assert.NotNil(t, parent.lastChild)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous)
	assert.Equal(t, parent.firstChild.Next(), child)

	// Re appending the last child
	err = parent.AppendChild(child)
	assert.Nil(t, err)
	assert.Equal(t, child, parent.lastChild)
	assert.Equal(t, parent.firstChild, child.previous) // Still only two elements
	assert.Equal(t, parent.firstChild.Next(), child)
}

func TestParent_AppendChild_AncestorOfParent(t *testing.T) {
	parent := newParent("test_parent")
	fakeChild := newParent("test_fake_child")

	err := fakeChild.AppendChild(parent)
	assert.Nil(t, err)

	err = parent.AppendChild(fakeChild)
	assert.NotNil(t, err)

	// Try again with Parent of fake child
	fakeChildParent := newParent("test_fake_child_parent")
	err = fakeChildParent.AppendChild(fakeChild)
	assert.Nil(t, err)

	err = parent.AppendChild(fakeChildParent)
	assert.NotNil(t, err)
}
