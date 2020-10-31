package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerNode_AppendChild(t *testing.T) {
	cn := makeContainerNode()

	child, err := cn.AppendChild(MakeTextNode("Hello world"))
	assert.Nil(t, err)
	assert.NotNil(t, child)

	assert.Equal(t, cn.LastChild(), cn.FirstChild())
	assert.Equal(t, child, cn.FirstChild())

	child, err = cn.AppendChild(MakeTextNode("Bonjour tout le monde"))
	assert.Nil(t, err)
	assert.NotNil(t, child)

	assert.NotEqual(t, cn.LastChild(), cn.FirstChild())
	assert.Equal(t, child, cn.LastChild())
	assert.Equal(t, cn.FirstChild(), cn.LastChild().Previous())
}

func TestContainerNode_InsertBefore(t *testing.T) {
	cn := makeContainerNode()

	child, err := cn.InsertBefore(MakeTextNode("Hello world"), nil)
	assert.Nil(t, err)
	assert.NotNil(t, child)

	child2, err := cn.InsertBefore(MakeTextNode("Bonjour tout le monde"), child)
	assert.Nil(t, err)
	assert.NotNil(t, child2)

	assert.Equal(t, child2, child.Previous())
	assert.Equal(t, child2, cn.FirstChild())
}
