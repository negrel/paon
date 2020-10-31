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
}
