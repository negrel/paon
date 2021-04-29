package tree

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestParentNode_New(t *testing.T) {
// 	parent := newParent()
// 	assert.NotNil(t, parent)
// }

// func TestParentNode_AppendChildNode_EmptyParent_Success(t *testing.T) {
// 	parent := newParent()
// 	child := newNode()

// 	// Appending it
// 	err := parent.AppendChildNode(child)
// 	assert.Nil(t, err)
// 	assert.Equal(t, child, parent.FirstChildNode())
// 	assert.Equal(t, child, parent.LastChildNode())
// 	assert.Equal(t, parent, child.parent)
// }

// func TestParentNode_AppendChildNode_NonEmptyParent_Success(t *testing.T) {
// 	parent := newParent()
// 	child := newNode()
// 	newChild := newNode()

// 	// Appending it
// 	_ = parent.AppendChildNode(child)

// 	err := parent.AppendChildNode(newChild)
// 	assert.Nil(t, err)
// 	assert.Equal(t, newChild, parent.LastChildNode())
// 	assert.Equal(t, parent.FirstChildNode(), newChild.previous)
// 	assert.Equal(t, parent.FirstChildNode(), child)
// 	assert.Equal(t, parent.FirstChildNode().NextNode(), newChild)
// 	assert.Equal(t, parent, child.parent)
// }

// func TestParentNode_AppendChildNode_AncestorOfParent(t *testing.T) {
// 	parent := newParent()
// 	ancestor := newParent()

// 	_ = ancestor.AppendChildNode(parent)

// 	// should return an error, ancestor is the parentNode
// 	err := parent.AppendChildNode(ancestor)
// 	assert.NotNil(t, err)
// }

// func TestParentNode_AppendChildNode_NilChild(t *testing.T) {
// 	parent := newParent()

// 	assert.Panics(t, func() {
// 		_ = parent.AppendChildNode(nil)
// 	}, "appending a nil child should panic. (\"assert\" build tag must be enabled)")
// }

// func TestParentNode_AppendChildNode_WithParent(t *testing.T) {
// 	parent := newParent()
// 	child1 := newNode()
// 	child2 := newNode()
// 	_ = parent.AppendChildNode(child1)
// 	_ = parent.AppendChildNode(child2)

// 	otherParent := newParent()
// 	node := newNode()
// 	_ = otherParent.AppendChildNode(node)

// 	err := parent.AppendChildNode(node)
// 	assert.Nil(t, err)
// 	assert.Equal(t, parent, node.parent)
// 	assert.Equal(t, child2, node.previous)
// 	assert.Equal(t, nil, node.next)
// }

// func TestParentNode_AppendChildNode_Itself(t *testing.T) {
// 	parent := NewParent()

// 	err := parent.AppendChildNode(parent)
// 	assert.NotNil(t, err)
// }

// func TestParentNode_InsertBeforeNode_EmptyParent_Success(t *testing.T) {
// 	parent := newParent()
// 	child := newNode()

// 	err := parent.InsertBeforeNode(nil, child)
// 	assert.Nil(t, err)
// 	assert.Equal(t, parent.FirstChildNode(), child)
// 	assert.Equal(t, parent.LastChildNode(), child)
// 	assert.Equal(t, child.parent, parent)
// 	assert.Equal(t, nil, child.next)
// 	assert.Equal(t, nil, child.previous)
// }

// func TestParentNode_InsertBeforeNode_NonEmptyParent_Success(t *testing.T) {
// 	parent := newParent()
// 	prev := newNode()
// 	next := newNode()
// 	child := newNode()

// 	_ = parent.AppendChildNode(prev)
// 	_ = parent.AppendChildNode(next)
// 	err := parent.InsertBeforeNode(next, child)
// 	assert.Nil(t, err)
// 	assert.Equal(t, prev, parent.FirstChildNode())
// 	assert.Equal(t, next, parent.LastChildNode())
// 	assert.Equal(t, parent, child.parent)
// 	assert.Equal(t, child, next.previous)
// 	assert.Equal(t, next, child.next)
// 	assert.Equal(t, child, prev.next)
// 	assert.Equal(t, prev, child.previous)
// }

// func TestParentNode_InsertBeforeNode_NilChild(t *testing.T) {
// 	parent := newParent()

// 	reference := newNode()
// 	_ = parent.AppendChildNode(reference)

// 	// Nil child
// 	assert.Panics(t, func() {
// 		_ = parent.InsertBeforeNode(reference, nil)
// 	}, "inserting a nil child should panic. (\"assert\" build tag must be enabled)")
// }

// func TestParentNode_InsertBeforeNode_EmptyParent_NilReference(t *testing.T) {
// 	// Inserting before a nil reference is same as appending

// 	parent := newParent()
// 	child := newNode()

// 	// Appending it
// 	err := parent.InsertBeforeNode(nil, child)
// 	assert.Nil(t, err)
// 	assert.Equal(t, child, parent.FirstChildNode())
// 	assert.Equal(t, child, parent.LastChildNode())
// }

// func TestParentNode_InsertBeforeNode_NonEmptyParent_NilReference(t *testing.T) {
// 	parent := newParent()
// 	child := newNode()
// 	_ = parent.InsertBeforeNode(nil, child)

// 	// appending another child
// 	otherChild := newNode()
// 	err := parent.InsertBeforeNode(nil, otherChild)
// 	assert.Nil(t, err)
// 	assert.Equal(t, parent.LastChildNode(), otherChild)
// 	assert.Equal(t, child, otherChild.previous)
// 	assert.Equal(t, child.next, otherChild)
// }

// func TestParentNode_InsertBeforeNode_AncestorOfParent(t *testing.T) {
// 	// Setting up the parent node
// 	parent := newParent()
// 	reference := newParent()
// 	_ = parent.AppendChildNode(reference)

// 	// Inserting the parent in the ancestor node
// 	ancestor := newParent()
// 	_ = ancestor.InsertBeforeNode(nil, parent)

// 	// Trying to insert the parent node in the ancestor node
// 	err := parent.InsertBeforeNode(reference, ancestor)
// 	assert.NotNil(t, err)
// }

// func TestParentNode_InsertBeforeNode_ChildWithParent(t *testing.T) {
// 	// Setting up the parent
// 	parent := newParent()
// 	prev := newNode()
// 	next := newNode()
// 	_ = parent.AppendChildNode(next)
// 	_ = parent.InsertBeforeNode(next, prev)

// 	// The initial parent of the child
// 	initialParent := newParent()
// 	reference := newNode()
// 	_ = initialParent.AppendChildNode(reference)

// 	child := newNode()
// 	_ = initialParent.InsertBeforeNode(reference, child)

// 	// Inserting the child to the new parent
// 	err := parent.InsertBeforeNode(next, child)
// 	assert.Nil(t, err)
// 	assert.Equal(t, parent, child.parent)
// 	assert.Equal(t, next, child.next)
// 	assert.Equal(t, prev, child.previous)
// }

// func TestParentNode_InsertBeforeNode_Itself(t *testing.T) {
// 	parent := NewParent()

// 	err := parent.InsertBeforeNode(nil, parent)
// 	assert.NotNil(t, err)
// }

// func TestParentNode_isAncestorOf_NilChild(t *testing.T) {
// 	parent := newParent()
// 	assert.False(t, parent.IsAncestorOf(nil))
// }

// func TestParentNode_RemoveChild_Success(t *testing.T) {
// 	parent := newParent()

// 	node := newNode()
// 	_ = parent.AppendChildNode(node)

// 	child := newNode()
// 	_ = parent.AppendChildNode(child)

// 	err := parent.RemoveChildNode(node)
// 	assert.Nil(t, err)
// 	assert.Equal(t, nil, node.ParentNode())
// 	assert.Equal(t, nil, node.NextNode())
// 	assert.Equal(t, child.PreviousNode(), nil)
// 	assert.Equal(t, parent.FirstChildNode(), child)
// 	assert.Equal(t, parent.LastChildNode(), child)
// }

// func TestParentNode_RemoveChild_NilChild(t *testing.T) {
// 	parent := newParent()

// 	// Remove child from parent with no child
// 	assert.Panics(t, func() {
// 		_ = parent.RemoveChildNode(nil)
// 	})

// 	node := newNode()
// 	_ = parent.AppendChildNode(node)

// 	// Parent with one child
// 	assert.Panics(t, func() {
// 		_ = parent.RemoveChildNode(nil)
// 	})
// }

// func TestParentNode_RemoveChild_ChildOfAnotherParent(t *testing.T) {
// 	node := newParent()
// 	parent := newParent()
// 	child := newNode()

// 	_ = parent.AppendChildNode(child)

// 	err := node.RemoveChildNode(child)
// 	assert.NotNil(t, err)
// }
