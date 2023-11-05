package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type nodeTest struct {
	name         string
	test         func(t *testing.T, constructor func() Node)
	leafNodeOnly bool
}

func TestNode(t *testing.T) {
	nodeTests := generateNodeTests()

	for _, methodTest := range nodeTests {
		t.Run(methodTest.name, func(t *testing.T) {
			methodTest.test(t, func() Node {
				return NewNode(nodeData)
			})
		})
	}
}

func generateNodeTests() []nodeTest {
	tests := []nodeTest{
		{
			name: "AppendChild",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("ToEmptyNode", func(t *testing.T) {
					testNodeAppendChildToEmptyNode(t, constructor())
				})
				t.Run("ToNonEmptyNode", func(t *testing.T) {
					testNodeAppendChildToNonEmptyNode(t, constructor())
				})
				t.Run("NilChildNode", func(t *testing.T) {
					testNodeAppendChildNilNode(t, constructor())
				})
				t.Run("ParentOfNode", func(t *testing.T) {
					testNodeAppendChildParentOfNode(t, constructor())
				})
				t.Run("GreatParentOfNode", func(t *testing.T) {
					testNodeAppendChildGreatParentOfNode(t, constructor())
				})
				t.Run("NodeWithParent", func(t *testing.T) {
					testNodeAppendChildNodeWithParent(t, constructor())
				})
				t.Run("Itself", func(t *testing.T) {
					testNodeAppendChildItself(t, constructor())
				})
			},
		},
		{
			name: "InsertBefore",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("Node", func(t *testing.T) {
					testNodeInsertBeforeNode(t, constructor())
				})
				t.Run("NonChildReference", func(t *testing.T) {
					testNodeInsertBeforeNodeNonChildReference(t, constructor())
				})
				t.Run("NilReference", func(t *testing.T) {
					testNodeInsertBeforeNilReference(t, constructor())
				})
				t.Run("NilChild", func(t *testing.T) {
					testNodeInsertBeforeNilChild(t, constructor())
				})
				t.Run("ParentOfNode", func(t *testing.T) {
					testNodeInsertBeforeParentOfNode(t, constructor())
				})
				t.Run("Itself", func(t *testing.T) {
					testNodeInsertBeforeItself(t, constructor())
				})
				t.Run("BetweenTwoNode", func(t *testing.T) {
					testNodeInsertBeforeBetweenTwoNode(t, constructor())
				})
			},
		},
		{
			name: "RemoveChild",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("", func(t *testing.T) {
					testNodeRemoveChild(t, constructor())
				})
				t.Run("Nil", func(t *testing.T) {
					testNodeRemoveChildNil(t, constructor())
				})
				t.Run("AnotherParentChild", func(t *testing.T) {
					testNodeRemoveChildAnotherParentChild(t, constructor())
				})
				t.Run("SecondChild", func(t *testing.T) {
					testNodeRemoveChildSecondChild(t, constructor())
				})
			},
		},
	}

	for _, test := range generateLeafNodeTests() {
		if !test.leafNodeOnly {
			tests = append(tests, test)
		}
	}

	return tests
}

func testNodeAppendChildToEmptyNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	child := NewNode(nil)

	err := node.AppendChild(child)
	require.NoError(t, err)

	require.True(t, node.IsSame(child.Parent()))
	require.True(t, node.FirstChild().IsSame(child))
	require.True(t, node.LastChild().IsSame(child))
}

func testNodeAppendChildToNonEmptyNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())
	firstChild := NewNode(nil)

	err := node.AppendChild(firstChild)
	require.NoError(t, err)

	lastChild := NewNode(nil)
	err = node.AppendChild(lastChild)
	require.NoError(t, err)

	require.True(t, lastChild.Parent().IsSame(node))
	require.True(t, lastChild.Parent().IsSame(firstChild.Parent()))

	require.True(t, lastChild.Previous().IsSame(firstChild))
	require.True(t, firstChild.Next().IsSame(lastChild))
}

func testNodeAppendChildNilNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	err := node.AppendChild(nil)
	require.Error(t, err)
}

func testNodeAppendChildParentOfNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode(nil)
	err := parent.AppendChild(node)
	require.NoError(t, err)

	err = node.AppendChild(parent)
	require.Error(t, err)
}

func testNodeAppendChildGreatParentOfNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	greatParent := NewNode(nil)
	parent := NewNode(nil)

	err := greatParent.AppendChild(parent)
	require.NoError(t, err)

	err = parent.AppendChild(node)
	require.NoError(t, err)

	err = node.AppendChild(greatParent)
	require.Error(t, err)
}

func testNodeAppendChildNodeWithParent(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)

	err := parent.AppendChild(child)
	require.NoError(t, err)

	err = node.AppendChild(child)
	require.NoError(t, err)
	require.True(t, node.IsSame(child.Parent()))
	require.True(t, node.FirstChild().IsSame(child))
	require.True(t, node.LastChild().IsSame(child))
}

func testNodeAppendChildItself(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	err := node.AppendChild(node)
	require.Error(t, err)
}

func testNodeInsertBeforeNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode(nil)
	child := NewNode(nil)

	err := node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(child, reference)
	require.NoError(t, err)

	require.True(t, node.IsSame(child.Parent()))
	require.True(t, node.FirstChild().IsSame(child))
	require.True(t, node.LastChild().IsSame(reference))

	require.True(t, child.IsSame(reference.Previous()))
	require.True(t, reference.IsSame(child.Next()))
}

func testNodeInsertBeforeNodeNonChildReference(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode(nil)
	err := NewNode(nil).AppendChild(reference)
	require.NoError(t, err)

	child := NewNode(nil)

	err = node.InsertBefore(reference, child)
	require.Error(t, err)
}

func testNodeInsertBeforeNilReference(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	child := NewNode(nil)

	err := node.InsertBefore(child, nil)
	require.NoError(t, err)

	require.True(t, node.IsSame(child.Parent()))
	require.True(t, node.FirstChild().IsSame(child))
	require.True(t, node.LastChild().IsSame(child))
}

func testNodeInsertBeforeNilChild(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode(nil)

	err := node.InsertBefore(nil, reference)
	require.Error(t, err)
}

func testNodeInsertBeforeParentOfNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode(nil)
	err := parent.AppendChild(node)
	require.NoError(t, err)

	reference := NewNode(nil)
	err = node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(parent, reference)
	require.Error(t, err)
}

func NodeInsertBeforeNodeGreatParentOfNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	greatParent := NewNode(nil)
	parent := NewNode(nil)
	err := greatParent.AppendChild(node)
	require.NoError(t, err)
	err = parent.AppendChild(node)
	require.NoError(t, err)

	reference := NewNode(nil)
	err = node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(greatParent, reference)
	require.Error(t, err)
}

func NodeInsertBeforeNodeChildWithParent(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)
	err := parent.AppendChild(child)
	require.NoError(t, err)

	reference := NewNode(nil)
	err = node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(child, reference)
	require.NoError(t, err)

	require.True(t, node.IsSame(child.Parent()))
	require.True(t, node.FirstChild().IsSame(child))
	require.True(t, node.LastChild().IsSame(reference))

	require.True(t, child.IsSame(reference.Previous()))
	require.True(t, reference.IsSame(child.Next()))
}

func testNodeInsertBeforeBetweenTwoNode(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	previous := NewNode(nil)
	next := NewNode(nil)
	child := NewNode(nil)

	err := node.AppendChild(previous)
	require.NoError(t, err)

	err = node.AppendChild(next)
	require.NoError(t, err)

	err = node.InsertBefore(child, next)
	require.NoError(t, err)

	require.True(t, node.IsSame(child.Parent()))
	require.True(t, node.FirstChild().IsSame(previous))
	require.True(t, node.LastChild().IsSame(next))

	require.True(t, child.IsSame(next.Previous()))
	require.True(t, next.IsSame(child.Next()))

	require.True(t, child.IsSame(previous.Next()))
	require.True(t, previous.IsSame(child.Previous()))
}

func testNodeInsertBeforeItself(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode(nil)
	err := node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(node, reference)
	require.Error(t, err)
}

func testNodeRemoveChild(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	firstChild := NewNode(nil)
	lastChild := NewNode(nil)

	err := node.AppendChild(firstChild)
	require.NoError(t, err)
	err = node.AppendChild(lastChild)
	require.NoError(t, err)

	err = node.RemoveChild(firstChild)
	require.NoError(t, err)

	require.True(t, lastChild.IsSame(node.FirstChild()))
	require.True(t, lastChild.IsSame(node.LastChild()))
	require.Nil(t, lastChild.Previous())
	require.Nil(t, lastChild.Next())
	require.Nil(t, firstChild.Previous())
	require.Nil(t, firstChild.Next())
	require.Nil(t, firstChild.Parent())
}

func testNodeRemoveChildNil(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	require.Panics(t, func() {
		err := node.RemoveChild(nil)
		require.Error(t, err)
	})
}

func testNodeRemoveChildAnotherParentChild(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)
	err := parent.AppendChild(child)
	require.NoError(t, err)

	err = node.RemoveChild(child)
	require.Error(t, err)
}

func testNodeRemoveChildSecondChild(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	first := NewNode(nil)
	child := NewNode(nil)
	third := NewNode(nil)

	err := node.AppendChild(first)
	require.NoError(t, err)

	err = node.AppendChild(child)
	require.NoError(t, err)

	err = node.AppendChild(third)
	require.NoError(t, err)

	err = node.RemoveChild(child)
	require.NoError(t, err)

	require.True(t, third.IsSame(first.Next()))
	require.True(t, third.Previous().IsSame(first))

	require.True(t, first.IsSame(third.Previous()))
	require.True(t, first.Next().IsSame(third))
}
