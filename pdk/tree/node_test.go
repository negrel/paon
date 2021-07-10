package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	child := NewNode(nil)

	err := node.AppendChild(child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(child))
}

func testNodeAppendChildToNonEmptyNode(t *testing.T, node Node) {
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

func testNodeAppendChildNilNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	assert.Panics(t, func() {
		err := node.AppendChild(nil)
		assert.NotNil(t, err)
	})
}

func testNodeAppendChildParentOfNode(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	err := parent.AppendChild(node)
	assert.Nil(t, err)

	err = node.AppendChild(parent)
	assert.NotNil(t, err)
}

func testNodeAppendChildGreatParentOfNode(t *testing.T, node Node) {
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

func testNodeAppendChildNodeWithParent(t *testing.T, node Node) {
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

func testNodeAppendChildItself(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	err := node.AppendChild(node)
	assert.NotNil(t, err)
}

func testNodeInsertBeforeNode(t *testing.T, node Node) {
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

func testNodeInsertBeforeNodeNonChildReference(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)
	err := NewNode(nil).AppendChild(reference)
	assert.Nil(t, err)

	child := NewNode(nil)

	err = node.InsertBefore(reference, child)
	assert.NotNil(t, err)
}

func testNodeInsertBeforeNilReference(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	child := NewNode(nil)

	err := node.InsertBefore(nil, child)
	assert.Nil(t, err)

	assert.True(t, node.IsSame(child.Parent()))
	assert.True(t, node.FirstChild().IsSame(child))
	assert.True(t, node.LastChild().IsSame(child))
}

func testNodeInsertBeforeNilChild(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)

	assert.Panics(t, func() {
		err := node.InsertBefore(reference, nil)
		assert.NotNil(t, err)
	})
}

func testNodeInsertBeforeParentOfNode(t *testing.T, node Node) {
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

func NodeInsertBeforeNodeGreatParentOfNode(t *testing.T, node Node) {
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

func NodeInsertBeforeNodeChildWithParent(t *testing.T, node Node) {
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

func testNodeInsertBeforeBetweenTwoNode(t *testing.T, node Node) {
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

func testNodeInsertBeforeItself(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	reference := NewNode(nil)
	err := node.AppendChild(reference)
	assert.Nil(t, err)

	err = node.InsertBefore(reference, node)
	assert.NotNil(t, err)
}

func testNodeRemoveChild(t *testing.T, node Node) {
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

func testNodeRemoveChildNil(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	assert.Panics(t, func() {
		err := node.RemoveChild(nil)
		assert.NotNil(t, err)
	})
}

func testNodeRemoveChildAnotherParentChild(t *testing.T, node Node) {
	assert.Nil(t, node.FirstChild())
	assert.Nil(t, node.LastChild())

	parent := NewNode(nil)
	child := NewNode(nil)
	err := parent.AppendChild(child)
	assert.Nil(t, err)

	err = node.RemoveChild(child)
	assert.NotNil(t, err)
}

func testNodeRemoveChildSecondChild(t *testing.T, node Node) {
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
