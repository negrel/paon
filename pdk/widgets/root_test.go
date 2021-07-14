package widgets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot_New(t *testing.T) {
	root := NewRoot()
	assert.Equal(t, root, root.Node().Unwrap())
	assert.Equal(t, root, root.Root())
}

func TestRoot_SetChild(t *testing.T) {
	root := NewRoot()
	child := newWidget()
	root.SetChild(child)

	assert.Equal(t, root, child.Root())
	assert.True(t, root.Node().IsSame(child.Node().Root()))

	assert.Equal(t, child.Parent(), child.Root())
}
