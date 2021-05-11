package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeafNode(t *testing.T) {
	leafNodeTests := generateLeafNodeTests()

	for _, methodTest := range leafNodeTests {
		for _, test := range methodTest.functions {
			t.Run(methodTest.name, func(t *testing.T) {
				test(t, NewLeafNode(nodeData))
			})
		}
	}
}

func generateLeafNodeTests() []nodeTest {
	return []nodeTest{
		{
			name: "Unwrap",
			functions: []func(t *testing.T, n Node){
				Node_Unwrap,
			},
		},
		{
			name: "FirstChild",
			functions: []func(t *testing.T, n Node){
				LeafNode_FirstChild_Nil,
			},
			leafNodeOnly: true,
		},
		{
			name: "LastChild",
			functions: []func(t *testing.T, n Node){
				LeafNode_LastChild_Nil,
			},
			leafNodeOnly: true,
		},
		{
			name: "IsAncestorOf",
			functions: []func(t *testing.T, n Node){
				LeafNode_IsAncestorOf_Nil,
			},
			leafNodeOnly: true,
		},
		{
			name: "AppendChild",
			functions: []func(t *testing.T, n Node){
				LeafNode_AppendChild_LeafNodeError,
			},
			leafNodeOnly: true,
		},
		{
			name: "InsertBefore",
			functions: []func(t *testing.T, n Node){
				LeafNode_InsertBefore_LeafNodeError,
			},
			leafNodeOnly: true,
		},
		{
			name: "RemoveChild",
			functions: []func(t *testing.T, n Node){
				LeafNode_RemoveChild_LeafNodeError,
			},
			leafNodeOnly: true,
		},
		{
			name: "IsSame",
			functions: []func(t *testing.T, n Node){
				Node_IsSame_true,
				Node_IsSame_false,
			},
		},
		{
			name: "IsDescendantOf",
			functions: []func(t *testing.T, n Node){
				Node_IsDescendantOf_NilParent,
				Node_IsDescendantOf_Parent,
				Node_IsDescendantOf_NonChildNode,
				Node_IsDescendantOf_PreviousParent,
				Node_IsDescendantOf_GreatParent,
			},
		},
		{
			name: "Root",
			functions: []func(t *testing.T, n Node){
				Node_Root_Nil,
				Node_Root_Parent,
				Node_Root_GreatParent,
			},
		},
	}
}

func Node_Unwrap(t *testing.T, node Node) {
	assert.Equal(t, nodeData, node.Unwrap())
}

func LeafNode_FirstChild_Nil(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
}

func LeafNode_LastChild_Nil(t *testing.T, node Node) {
	assert.Nil(t, node.LastChild())
}

func Node_IsSame_true(t *testing.T, node Node) {
	assert.True(t, node.IsSame(node))
}

func Node_IsSame_false(t *testing.T, node Node) {
	assert.False(t, node.IsSame(NewNode(nil)))
}

func Node_IsDescendantOf_NilParent(t *testing.T, node Node) {
	node.SetParent(nil)
	assert.False(t, node.IsDescendantOf(nil))
}

func Node_IsDescendantOf_Parent(t *testing.T, node Node) {
	parent := NewNode(nil)
	node.SetParent(parent)

	assert.True(t, node.IsDescendantOf(parent))
}

func Node_IsDescendantOf_NonChildNode(t *testing.T, node Node) {
	otherNode := NewNode(nil)
	node.SetParent(nil)

	assert.False(t, node.IsDescendantOf(otherNode))
}

func Node_IsDescendantOf_PreviousParent(t *testing.T, node Node) {
	parent := NewNode(nil)

	err := parent.AppendChild(node)
	assert.Nil(t, err)

	err = parent.RemoveChild(node)
	assert.Nil(t, err)

	assert.False(t, node.IsDescendantOf(parent))
}

func Node_IsDescendantOf_GreatParent(t *testing.T, node Node) {
	greatParent := NewNode(nil)
	parent := NewNode(nil)

	err := parent.AppendChild(node)
	assert.Nil(t, err)

	err = greatParent.AppendChild(node)
	assert.Nil(t, err)

	assert.True(t, node.IsDescendantOf(greatParent))
}

func Node_Root_Nil(t *testing.T, node Node) {
	node.SetParent(nil)
	assert.Equal(t, node.Root(), nil)
}

func Node_Root_Parent(t *testing.T, node Node) {
	root := NewRoot(nil)
	err := root.AppendChild(node)
	assert.Nil(t, err)
	assert.Equal(t, root, node.Root())
}

func Node_Root_GreatParent(t *testing.T, node Node) {
	root := NewRoot(nil)
	parent := NewNode(nil)

	err := root.AppendChild(parent)
	assert.Nil(t, err)

	err = parent.AppendChild(node)
	assert.Nil(t, err)

	assert.Equal(t, root, node.Root())
}

func LeafNode_IsAncestorOf_Nil(t *testing.T, node Node) {
	assert.False(t, node.IsAncestorOf(nil))
}

func LeafNode_AppendChild_LeafNodeError(t *testing.T, node Node) {
	err := node.AppendChild(NewNode(nil))
	assert.NotNil(t, err)
}

func LeafNode_InsertBefore_LeafNodeError(t *testing.T, node Node) {
	reference := NewNode(nil)
	_ = node.AppendChild(reference)

	err := node.InsertBefore(reference, NewNode(nil))
	assert.NotNil(t, err)
}

func LeafNode_RemoveChild_LeafNodeError(t *testing.T, node Node) {
	child := NewNode(nil)
	_ = node.AppendChild(child)

	err := node.RemoveChild(child)
	assert.NotNil(t, err)
}
