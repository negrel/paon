package tree

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type nodeTest struct {
	name string
	test func(t *testing.T, constructor func() *Node[any])
}

func TestNode(t *testing.T) {
	nodeTests := generateNodeTests()

	for _, methodTest := range nodeTests {
		t.Run(methodTest.name, func(t *testing.T) {
			methodTest.test(t, func() *Node[any] {
				return NewNode[any](time.Now())
			})
		})
	}
}

func generateNodeTests() []nodeTest {
	tests := []nodeTest{
		{
			name: "IsDescendantOf",
			test: func(t *testing.T, constructor func() *Node[any]) {
				t.Run("NilParent", func(t *testing.T) {
					testNodeIsDescendantOfNilParent(t, constructor())
				})
				t.Run("Parent", func(t *testing.T) {
					testNodeIsDescendantOfParent(t, constructor())
				})
				t.Run("NonChildNode", func(t *testing.T) {
					testNodeIsDescendantOfNonChildNode(t, constructor())
				})
				t.Run("PreviousParent", func(t *testing.T) {
					testNodeIsDescendantOfPreviousParent(t, constructor())
				})
				t.Run("GreatParent", func(t *testing.T) {
					testNodeIsDescendantOfGreatParent(t, constructor())
				})
			},
		},
		{
			name: "Root",
			test: func(t *testing.T, constructor func() *Node[any]) {
				t.Run("Nil", func(t *testing.T) {
					testNodeRootNil(t, constructor())
				})
				t.Run("Parent", func(t *testing.T) {
					testNodeRootParent(t, constructor())
				})
				t.Run("GreatParent", func(t *testing.T) {
					testNodeRootGreatParent(t, constructor())
				})
			},
		},
		{
			name: "AppendChild",
			test: func(t *testing.T, constructor func() *Node[any]) {
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
			test: func(t *testing.T, constructor func() *Node[any]) {
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
			test: func(t *testing.T, constructor func() *Node[any]) {
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

	return tests
}

func testNodeIsDescendantOfNilParent(t *testing.T, node *Node[any]) {
	node.parent = nil
	require.False(t, node.IsDescendantOf(nil))
}

func testNodeIsDescendantOfParent(t *testing.T, node *Node[any]) {
	parent := NewNode[any](nil)
	node.parent = parent

	require.True(t, node.IsDescendantOf(parent))
}

func testNodeIsDescendantOfNonChildNode(t *testing.T, node *Node[any]) {
	otherNode := NewNode[any](nil)
	node.parent = nil

	require.False(t, node.IsDescendantOf(otherNode))
}

func testNodeIsDescendantOfPreviousParent(t *testing.T, node *Node[any]) {
	parent := NewNode[any](nil)

	err := parent.AppendChild(node)
	require.NoError(t, err)

	err = parent.RemoveChild(node)
	require.NoError(t, err)

	require.False(t, node.IsDescendantOf(parent))
}

func testNodeIsDescendantOfGreatParent(t *testing.T, node *Node[any]) {
	greatParent := NewNode[any](nil)
	parent := NewNode[any](nil)

	err := parent.AppendChild(node)
	require.NoError(t, err)

	err = greatParent.AppendChild(node)
	require.NoError(t, err)

	require.True(t, node.IsDescendantOf(greatParent))
}

func testNodeRootNil(t *testing.T, node *Node[any]) {
	node.parent = nil
	require.Equal(t, node.Root(), node)
}

func testNodeRootParent(t *testing.T, node *Node[any]) {
	root := NewNode[any](nil)
	err := root.AppendChild(node)
	require.NoError(t, err)
	require.Equal(t, root, node.Root())
}

func testNodeRootGreatParent(t *testing.T, node *Node[any]) {
	root := NewNode[any](nil)
	parent := NewNode[any](nil)

	err := root.AppendChild(parent)
	require.NoError(t, err)

	err = parent.AppendChild(node)
	require.NoError(t, err)

	require.Equal(t, root, node.Root())
}

func testNodeAppendChildToEmptyNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	child := NewNode[any](nil)

	err := node.AppendChild(child)
	require.NoError(t, err)

	require.Equal(t, node, child.Parent())
	require.Equal(t, node.FirstChild(), child)
	require.Equal(t, node.LastChild(), child)
}

func testNodeAppendChildToNonEmptyNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())
	firstChild := NewNode[any](nil)

	err := node.AppendChild(firstChild)
	require.NoError(t, err)

	lastChild := NewNode[any](nil)
	err = node.AppendChild(lastChild)
	require.NoError(t, err)

	require.Equal(t, lastChild.Parent(), node)
	require.Equal(t, lastChild.Parent(), firstChild.Parent())

	require.Equal(t, lastChild.Previous(), firstChild)
	require.Equal(t, firstChild.Next(), lastChild)
}

func testNodeAppendChildNilNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	err := node.AppendChild(nil)
	require.Error(t, err)
}

func testNodeAppendChildParentOfNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode[any](nil)
	err := parent.AppendChild(node)
	require.NoError(t, err)

	err = node.AppendChild(parent)
	require.Error(t, err)
}

func testNodeAppendChildGreatParentOfNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	greatParent := NewNode[any](nil)
	parent := NewNode[any](nil)

	err := greatParent.AppendChild(parent)
	require.NoError(t, err)

	err = parent.AppendChild(node)
	require.NoError(t, err)

	err = node.AppendChild(greatParent)
	require.Error(t, err)
}

func testNodeAppendChildNodeWithParent(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode[any](nil)
	child := NewNode[any](nil)

	err := parent.AppendChild(child)
	require.NoError(t, err)

	err = node.AppendChild(child)
	require.NoError(t, err)
	require.Equal(t, node, child.Parent())
	require.Equal(t, node.FirstChild(), child)
	require.Equal(t, node.LastChild(), child)
}

func testNodeAppendChildItself(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	err := node.AppendChild(node)
	require.Error(t, err)
}

func testNodeInsertBeforeNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode[any](nil)
	child := NewNode[any](nil)

	err := node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(child, reference)
	require.NoError(t, err)

	require.Equal(t, node, child.Parent())
	require.Equal(t, node.FirstChild(), child)
	require.Equal(t, node.LastChild(), reference)

	require.Equal(t, child, reference.Previous())
	require.Equal(t, reference, child.Next())
}

func testNodeInsertBeforeNodeNonChildReference(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode[any](nil)
	err := NewNode[any](nil).AppendChild(reference)
	require.NoError(t, err)

	child := NewNode[any](nil)

	err = node.InsertBefore(reference, child)
	require.Error(t, err)
}

func testNodeInsertBeforeNilReference(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	child := NewNode[any](nil)

	err := node.InsertBefore(child, nil)
	require.NoError(t, err)

	require.Equal(t, node, child.Parent())
	require.Equal(t, node.FirstChild(), child)
	require.Equal(t, node.LastChild(), child)
}

func testNodeInsertBeforeNilChild(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode[any](nil)

	err := node.InsertBefore(nil, reference)
	require.Error(t, err)
}

func testNodeInsertBeforeParentOfNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode[any](nil)
	err := parent.AppendChild(node)
	require.NoError(t, err)

	reference := NewNode[any](nil)
	err = node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(parent, reference)
	require.Error(t, err)
}

func NodeInsertBeforeNodeGreatParentOfNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	greatParent := NewNode[any](nil)
	parent := NewNode[any](nil)
	err := greatParent.AppendChild(node)
	require.NoError(t, err)
	err = parent.AppendChild(node)
	require.NoError(t, err)

	reference := NewNode[any](nil)
	err = node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(greatParent, reference)
	require.Error(t, err)
}

func NodeInsertBeforeNodeChildWithParent(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode[any](nil)
	child := NewNode[any](nil)
	err := parent.AppendChild(child)
	require.NoError(t, err)

	reference := NewNode[any](nil)
	err = node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(child, reference)
	require.NoError(t, err)

	require.Equal(t, node, child.Parent())
	require.Equal(t, node.FirstChild(), child)
	require.Equal(t, node.LastChild(), reference)

	require.Equal(t, child, reference.Previous())
	require.Equal(t, reference, child.Next())
}

func testNodeInsertBeforeBetweenTwoNode(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	previous := NewNode[any](nil)
	next := NewNode[any](nil)
	child := NewNode[any](nil)

	err := node.AppendChild(previous)
	require.NoError(t, err)

	err = node.AppendChild(next)
	require.NoError(t, err)

	err = node.InsertBefore(child, next)
	require.NoError(t, err)

	require.Equal(t, node, child.Parent())
	require.Equal(t, node.FirstChild(), previous)
	require.Equal(t, node.LastChild(), next)

	require.Equal(t, child, next.Previous())
	require.Equal(t, next, child.Next())

	require.Equal(t, child, previous.Next())
	require.Equal(t, previous, child.Previous())
}

func testNodeInsertBeforeItself(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	reference := NewNode[any](nil)
	err := node.AppendChild(reference)
	require.NoError(t, err)

	err = node.InsertBefore(node, reference)
	require.Error(t, err)
}

func testNodeRemoveChild(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	firstChild := NewNode[any](nil)
	lastChild := NewNode[any](nil)

	err := node.AppendChild(firstChild)
	require.NoError(t, err)
	err = node.AppendChild(lastChild)
	require.NoError(t, err)

	err = node.RemoveChild(firstChild)
	require.NoError(t, err)

	require.Equal(t, lastChild, node.FirstChild())
	require.Equal(t, lastChild, node.LastChild())
	require.Nil(t, lastChild.Previous())
	require.Nil(t, lastChild.Next())
	require.Nil(t, firstChild.Previous())
	require.Nil(t, firstChild.Next())
	require.Nil(t, firstChild.Parent())
}

func testNodeRemoveChildNil(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	require.Panics(t, func() {
		err := node.RemoveChild(nil)
		require.Error(t, err)
	})
}

func testNodeRemoveChildAnotherParentChild(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	parent := NewNode[any](nil)
	child := NewNode[any](nil)
	err := parent.AppendChild(child)
	require.NoError(t, err)

	err = node.RemoveChild(child)
	require.Error(t, err)
}

func testNodeRemoveChildSecondChild(t *testing.T, node *Node[any]) {
	require.Nil(t, node.FirstChild())
	require.Nil(t, node.LastChild())

	first := NewNode[any](nil)
	child := NewNode[any](nil)
	third := NewNode[any](nil)

	err := node.AppendChild(first)
	require.NoError(t, err)

	err = node.AppendChild(child)
	require.NoError(t, err)

	err = node.AppendChild(third)
	require.NoError(t, err)

	err = node.RemoveChild(child)
	require.NoError(t, err)

	require.Equal(t, third, first.Next())
	require.Equal(t, third.Previous(), first)

	require.Equal(t, first, third.Previous())
	require.Equal(t, first.Next(), third)
}
