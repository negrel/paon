package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerNode_AppendChild(t *testing.T) {
	cn := newContainerNode()

	child, err := cn.AppendChild(MakeTextNode("Hello world"))
	assert.Nil(t, err)
	assert.NotNil(t, child)
	assert.Equal(t, child.Parent(), cn)
	assert.Nil(t, child.Owner())
	assert.Equal(t, child.Owner(), cn.Owner())

	assert.Equal(t, cn.LastChild(), cn.FirstChild())
	assert.Equal(t, child, cn.FirstChild())

	child, err = cn.AppendChild(MakeTextNode("Bonjour tout le monde"))
	assert.Nil(t, err)
	assert.NotNil(t, child)
	assert.Equal(t, child.Parent(), cn)
	assert.Nil(t, child.Owner())
	assert.Equal(t, child.Owner(), cn.Owner())

	assert.NotEqual(t, cn.LastChild(), cn.FirstChild())
	assert.Equal(t, child, cn.LastChild())
	assert.Equal(t, cn.FirstChild(), cn.LastChild().Previous())
}

func TestContainerNode_FailAdoptingDocument(t *testing.T) {
	cn := newContainerNode()

	_, err := cn.AppendChild(NewDocument())
	assert.NotNil(t, err, "appending a DocumentNode should fail")

	child, _ := cn.AppendChild(MakeTextNode("Hello world"))
	_, err = cn.InsertBefore(child, NewDocument())
	assert.NotNil(t, err, "inserting a DocumentNode should fail")
}

func TestContainerNode_InsertBefore(t *testing.T) {
	cn := newContainerNode()

	child, err := cn.InsertBefore(nil, MakeTextNode("Hello world"))
	assert.Nil(t, err)
	assert.NotNil(t, child)
	assert.Equal(t, child.Parent(), cn)
	assert.Nil(t, child.Owner())
	assert.Equal(t, child.Owner(), cn.Owner())

	assert.Equal(t, child, cn.LastChild())
	assert.Equal(t, child, cn.FirstChild())

	child, err = cn.InsertBefore(child, MakeTextNode("Bonjour tout le monde"))
	assert.Nil(t, err)
	assert.NotNil(t, child)
	assert.Equal(t, child.Parent(), cn)
	assert.Nil(t, child.Owner())
	assert.Equal(t, child.Owner(), cn.Owner())

	assert.NotEqual(t, cn.FirstChild(), cn.LastChild())
	assert.Equal(t, child, cn.LastChild().Previous())
	assert.Equal(t, child, cn.FirstChild())

	child, err = cn.InsertBefore(cn.LastChild(), MakeTextNode("Hola mundo"))
	assert.Nil(t, err)
	assert.NotNil(t, child)
	assert.Equal(t, child.Parent(), cn)
	assert.Nil(t, child.Owner())
	assert.Equal(t, child.Owner(), cn.Owner())

	assert.Equal(t, child, cn.FirstChild().Next())
	assert.Equal(t, child, cn.LastChild().Previous())
}

func TestContainerNode_Fail_InsertBefore_NonChildReference(t *testing.T) {
	cn := newContainerNode()

	_, err := cn.InsertBefore(MakeTextNode("Hello world"), MakeTextNode("non child reference"))
	assert.NotNil(t, err, "inserting a child before a reference that is not contained by the parent should fail")
}
