package tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeafNode(t *testing.T) {
	leafNodeTests := generateLeafNodeTests()

	for _, methodTest := range leafNodeTests {
		t.Run(methodTest.name, func(t *testing.T) {
			methodTest.test(t, func() Node {
				return NewLeafNode(nodeData)
			})
		})
	}
}

func generateLeafNodeTests() []nodeTest {
	return []nodeTest{
		{
			name: "Unwrap",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("EqualToTheNode", func(t *testing.T) {
					testNodeUnwrap(t, constructor())
				})
			},
		},
		{
			name: "FirstChild",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("Nil", func(t *testing.T) {
					testLeafNodeFirstChildNil(t, constructor())
				})
			},
			leafNodeOnly: true,
		},
		{
			name: "LastChild",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("Nil", func(t *testing.T) {
					testLeafNodeLastChildNil(t, constructor())
				})
			},
			leafNodeOnly: true,
		},
		{
			name: "IsAncestorOf",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("Nil", func(t *testing.T) {
					testLeafNodeIsAncestorOfNil(t, constructor())
				})
			},
			leafNodeOnly: true,
		},
		{
			name: "AppendChild",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("LeafNodeError", func(t *testing.T) {
					testLeafNodeAppendChildLeafNodeError(t, constructor())
				})
			},
			leafNodeOnly: true,
		},
		{
			name: "InsertBefore",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("LeafNodeError", func(t *testing.T) {
					testLeafNodeInsertBeforeLeafNodeError(t, constructor())
				})
			},
			leafNodeOnly: true,
		},
		{
			name: "RemoveChild",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("LeafNodeError", func(t *testing.T) {
					testLeafNodeRemoveChildLeafNodeError(t, constructor())
				})
			},
			leafNodeOnly: true,
		},
		{
			name: "IsSame",
			test: func(t *testing.T, constructor func() Node) {
				t.Run("True", func(t *testing.T) {
					testNodeIsSameTrue(t, constructor())
				})
				t.Run("False", func(t *testing.T) {
					testNodeIsSameFalse(t, constructor())
				})
			},
		},
		{
			name: "IsDescendantOf",
			test: func(t *testing.T, constructor func() Node) {
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
			test: func(t *testing.T, constructor func() Node) {
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
	}
}

func testNodeUnwrap(t *testing.T, node Node) {
	require.Equal(t, nodeData, node.Unwrap())
}

func testLeafNodeFirstChildNil(t *testing.T, node Node) {
	require.Nil(t, node.FirstChild())
}

func testLeafNodeLastChildNil(t *testing.T, node Node) {
	require.Nil(t, node.LastChild())
}

func testNodeIsSameTrue(t *testing.T, node Node) {
	require.True(t, node.IsSame(node))
}

func testNodeIsSameFalse(t *testing.T, node Node) {
	require.False(t, node.IsSame(NewNode(nil)))
}

func testNodeIsDescendantOfNilParent(t *testing.T, node Node) {
	node.SetParent(nil)
	require.False(t, node.IsDescendantOf(nil))
}

func testNodeIsDescendantOfParent(t *testing.T, node Node) {
	parent := NewNode(nil)
	node.SetParent(parent)

	require.True(t, node.IsDescendantOf(parent))
}

func testNodeIsDescendantOfNonChildNode(t *testing.T, node Node) {
	otherNode := NewNode(nil)
	node.SetParent(nil)

	require.False(t, node.IsDescendantOf(otherNode))
}

func testNodeIsDescendantOfPreviousParent(t *testing.T, node Node) {
	parent := NewNode(nil)

	err := parent.AppendChild(node)
	require.NoError(t, err)

	err = parent.RemoveChild(node)
	require.NoError(t, err)

	require.False(t, node.IsDescendantOf(parent))
}

func testNodeIsDescendantOfGreatParent(t *testing.T, node Node) {
	greatParent := NewNode(nil)
	parent := NewNode(nil)

	err := parent.AppendChild(node)
	require.NoError(t, err)

	err = greatParent.AppendChild(node)
	require.NoError(t, err)

	require.True(t, node.IsDescendantOf(greatParent))
}

func testNodeRootNil(t *testing.T, node Node) {
	node.SetParent(nil)
	require.Nil(t, node.Root())
}

func testNodeRootParent(t *testing.T, node Node) {
	root := NewRoot(nil)
	err := root.AppendChild(node)
	require.NoError(t, err)
	require.Equal(t, root, node.Root())
}

func testNodeRootGreatParent(t *testing.T, node Node) {
	root := NewRoot(nil)
	parent := NewNode(nil)

	err := root.AppendChild(parent)
	require.NoError(t, err)

	err = parent.AppendChild(node)
	require.NoError(t, err)

	require.Equal(t, root, node.Root())
}

func testLeafNodeIsAncestorOfNil(t *testing.T, node Node) {
	require.False(t, node.IsAncestorOf(nil))
}

func testLeafNodeAppendChildLeafNodeError(t *testing.T, node Node) {
	err := node.AppendChild(NewNode(nil))
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrLeafNode()))
}

func testLeafNodeInsertBeforeLeafNodeError(t *testing.T, node Node) {
	reference := NewNode(nil)
	_ = node.AppendChild(reference)

	err := node.InsertBefore(reference, NewNode(nil))
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrLeafNode()))
}

func testLeafNodeRemoveChildLeafNodeError(t *testing.T, node Node) {
	child := NewNode(nil)
	_ = node.AppendChild(child)

	err := node.RemoveChild(child)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrLeafNode()))
}
