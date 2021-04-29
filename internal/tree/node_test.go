package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	t.Run("leafNode", func(t *testing.T) {
		LeafNodeImpl(t, func() Node {
			return newLeafNode(nil)
		})
	})

	t.Run("branchNode", func(t *testing.T) {
		LeafNodeImpl(t, func() Node {
			return newNode(nil)
		})

		BranchNodeImpl(t, func() Node {
			return newNode(nil)
		})
	})
}

func LeafNodeImpl(t *testing.T, nodeConstructor func() Node) {
	t.Run("IsSame", func(t *testing.T) {
		Node_IsSame_true(t, nodeConstructor())
		Node_IsSame_false(t, nodeConstructor())
	})

	t.Run("IsDescendantOf", func(t *testing.T) {
		Node_IsDescendantOf_NilParent(t, nodeConstructor())
		Node_IsDescendantOf_Parent(t, nodeConstructor())
		Node_IsDescendantOf_NonChildNode(t, nodeConstructor())
		Node_IsDescendantOf_PreviousParent(t, nodeConstructor())
		Node_IsDescendantOf_GreatParent(t, nodeConstructor())
	})

	t.Run("Root", func(t *testing.T) {
		Node_Root_Nil(t, nodeConstructor())
		Node_Root_Parent(t, nodeConstructor())
		Node_Root_GreatParent(t, nodeConstructor())
	})
}

func BranchNodeImpl(t *testing.T, nodeConstructor func() Node) {
	t.Run("AppendChild", func(t *testing.T) {
		Node_AppendChild_ToEmptyNode(t, nodeConstructor())
		Node_AppendChild_ToNonEmptyNode(t, nodeConstructor())
		Node_AppendChild_NilNode(t, nodeConstructor())
		Node_AppendChild_ParentOfNode(t, nodeConstructor())
		Node_AppendChild_GreatParentOfNode(t, nodeConstructor())
		Node_AppendChild_NodeWithParent(t, nodeConstructor())
		Node_AppendChild_Itself(t, nodeConstructor())
	})

	t.Run("InsertBefore", func(t *testing.T) {
		Node_InsertBeforeNode(t, nodeConstructor())
		Node_InsertBeforeNode_NonChildReference(t, nodeConstructor())
		Node_InsertBeforeNode_ToEmptyParent_NilReference(t, nodeConstructor())
		Node_InsertBeforeNode_ToEmptyParent_NilChild(t, nodeConstructor())
		Node_InsertBeforeNode_ParentOfNode(t, nodeConstructor())
		Node_InsertBeforeNode_Itself(t, nodeConstructor())
	})

	t.Run("RemoveChild", func(t *testing.T) {
		Node_RemoveChild(t, nodeConstructor())
		Node_RemoveChild_Nil(t, nodeConstructor())
		Node_RemoveChild_AnotherParentChild(t, nodeConstructor())
	})
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
