package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type nodeTest struct {
	name           string
	functions      []func(t *testing.T, n Node)
	leafNodeOnly   bool
	concurrentOnly bool
}

func TestNode(t *testing.T) {
	nodeTests := generateNodeTests()

	for _, methodTest := range nodeTests {
		for _, test := range methodTest.functions {
			t.Run(methodTest.name, func(t *testing.T) {
				test(t, NewNode(nodeData))
			})
		}
	}
}

func generateNodeTests() []nodeTest {
	tests := []nodeTest{
		{
			name: "AppendChild",
			functions: []func(t *testing.T, n Node){
				Node_AppendChild_ToEmptyNode,
				Node_AppendChild_ToNonEmptyNode,
				Node_AppendChild_NilNode,
				Node_AppendChild_ParentOfNode,
				Node_AppendChild_GreatParentOfNode,
				Node_AppendChild_NodeWithParent,
				Node_AppendChild_Itself,
			},
		},
		{
			name: "InsertBefore",
			functions: []func(t *testing.T, n Node){
				Node_InsertBeforeNode,
				Node_InsertBeforeNode_NonChildReference,
				Node_InsertBeforeNode_ToEmptyParent_NilReference,
				Node_InsertBeforeNode_ToEmptyParent_NilChild,
				Node_InsertBeforeNode_ParentOfNode,
				Node_InsertBeforeNode_Itself,
				Node_InsertBeforeNode_BetweenTwoNode,
			},
		},
		{
			name: "RemoveChild",
			functions: []func(t *testing.T, n Node){
				Node_RemoveChild,
				Node_RemoveChild_Nil,
				Node_RemoveChild_AnotherParentChild,
				Node_RemoveChild_SecondChild,
			},
		},
	}

	for _, test := range generateLeafNodeTests() {
		if !test.leafNodeOnly && !test.concurrentOnly {
			tests = append(tests, test)
		}
	}

	return tests
}

func Node_AppendChild_ToEmptyNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	child := NewNode(nil)

	err := node.AppendChild(child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(child))
}

func Node_AppendChild_ToNonEmptyNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())
	firstChild := NewNode(nil)

	err := node.AppendChild(firstChild)
	assert.Nil(t, err)

	lastChild := NewNode(nil)
	err = node.AppendChild(lastChild)
	assert.Nil(t, err)

	assert.True(t, lastChild.Parent().IsSame(node))
	assert.True(t, lastChild.Parent().IsSame(firstChild.Parent()))

	assert.True(t, lastChild.Previous().IsSame(firstChild))
	assert.True(t, firstChild.Next().IsSame(lastChild))
}

func Node_AppendChild_NilNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	assert.Panics(t, func() {
		err := node.AppendChild(nil)
		assert.NotNil(t, err)
	})
}

func Node_AppendChild_ParentOfNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	err := parent.AppendChild(node)
	assert.Nil(t, err)

	err = node.AppendChild(parent)
	assert.NotNil(t, err)
}

func Node_AppendChild_GreatParentOfNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	greatParent := NewNode(nil)
	parent := NewNode(nil)

	err := greatParent.AppendChild(parent)
	assert.Nil(t, err)

	err = parent.AppendChild(node)
	assert.Nil(t, err)

	err = node.AppendChild(greatParent)
	assert.NotNil(t, err)
}

func Node_AppendChild_NodeWithParent(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)

	err := parent.AppendChild(child)
	assert.Nil(t, err)

	node.AppendChild(child)
	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(child))
}

func Node_AppendChild_Itself(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	err := node.AppendChild(node)
	assert.NotNil(t, err)
}

func Node_InsertBeforeNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)
	child := NewNode(nil)

	err := node.AppendChild(reference)
	assert.Nil(t, err)

	err = node.InsertBefore(reference, child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(reference))

	assert.True(t, child.IsSame(reference.Previous()))
	assert.True(t, reference.IsSame(child.Next()))
}

func Node_InsertBeforeNode_NonChildReference(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)
	err := NewNode(nil).AppendChild(reference)
	assert.Nil(t, err)

	child := NewNode(nil)

	err = node.InsertBefore(reference, child)
	assert.NotNil(t, err)
}

func Node_InsertBeforeNode_ToEmptyParent_NilReference(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	child := NewNode(nil)

	err := node.InsertBefore(nil, child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(child))
}

func Node_InsertBeforeNode_ToEmptyParent_NilChild(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)

	assert.Panics(t, func() {
		err := node.InsertBefore(reference, nil)
		assert.NotNil(t, err)
	})
}

func Node_InsertBeforeNode_ParentOfNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	err := parent.AppendChild(node)
	assert.Nil(t, err)

	reference := NewNode(nil)
	err = node.AppendChild(reference)
	assert.Nil(t, err)

	err = node.InsertBefore(reference, parent)
	assert.NotNil(t, err)
}

func Node_InsertBeforeNode_GreatParentOfNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	greatParent := NewNode(nil)
	parent := NewNode(nil)
	err := greatParent.AppendChild(node)
	assert.Nil(t, err)
	err = parent.AppendChild(node)
	assert.Nil(t, err)

	reference := NewNode(nil)
	err = node.AppendChild(reference)
	assert.Nil(t, err)

	err = node.InsertBefore(reference, greatParent)
	assert.NotNil(t, err)
}

func Node_InsertBeforeNode_ChildWithParent(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)
	err := parent.AppendChild(child)
	assert.Nil(t, err)

	reference := NewNode(nil)
	err = node.AppendChild(reference)
	assert.Nil(t, err)

	err = node.InsertBefore(reference, child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(reference))

	assert.True(t, child.IsSame(reference.Previous()))
	assert.True(t, reference.IsSame(child.Next()))
}

func Node_InsertBeforeNode_BetweenTwoNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	previous := NewNode(nil)
	next := NewNode(nil)
	child := NewNode(nil)

	err := node.AppendChild(previous)
	assert.Nil(t, err)

	err = node.AppendChild(next)
	assert.Nil(t, err)

	err = node.InsertBefore(next, child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(previous))
	assert.True(t, node.LastChild().IsSame(next))

	assert.True(t, child.IsSame(next.Previous()))
	assert.True(t, next.IsSame(child.Next()))

	assert.True(t, child.IsSame(previous.Next()))
	assert.True(t, previous.IsSame(child.Previous()))
}

func Node_InsertBeforeNode_Itself(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)
	err := node.AppendChild(reference)
	assert.Nil(t, err)

	err = node.InsertBefore(reference, node)
	assert.NotNil(t, err)
}

func Node_RemoveChild(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	firstChild := NewNode(nil)
	lastChild := NewNode(nil)

	err := node.AppendChild(firstChild)
	assert.Nil(t, err)
	err = node.AppendChild(lastChild)
	assert.Nil(t, err)

	err = node.RemoveChild(firstChild)
	assert.Nil(t, err)

	assert.True(t, lastChild.IsSame(node.FirstChild()))
	assert.True(t, lastChild.IsSame(node.LastChild()))
	assert.Nil(t, lastChild.Previous())
	assert.Nil(t, lastChild.Next())
	assert.Nil(t, firstChild.Previous())
	assert.Nil(t, firstChild.Next())
	assert.Nil(t, firstChild.Parent())
}

func Node_RemoveChild_Nil(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	assert.Panics(t, func() {
		err := node.RemoveChild(nil)
		assert.NotNil(t, err)
	})
}

func Node_RemoveChild_AnotherParentChild(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)
	err := parent.AppendChild(child)
	assert.Nil(t, err)

	err = node.RemoveChild(child)
	assert.NotNil(t, err)
}

func Node_RemoveChild_SecondChild(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	first := NewNode(nil)
	child := NewNode(nil)
	third := NewNode(nil)

	err := node.AppendChild(first)
	assert.Nil(t, err)

	err = node.AppendChild(child)
	assert.Nil(t, err)

	err = node.AppendChild(third)
	assert.Nil(t, err)

	err = node.RemoveChild(child)
	assert.Nil(t, err)

	assert.True(t, third.IsSame(first.Next()))
	assert.True(t, third.Previous().IsSame(first))

	assert.True(t, first.IsSame(third.Previous()))
	assert.True(t, first.Next().IsSame(third))
}
