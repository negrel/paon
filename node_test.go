package gom

import (
	"math/rand"
	"testing"
)

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

func TestChildNodes(t *testing.T) {
	// Creating node & child
	node := newNode()
	child := newNode()

	// Appending child
	child = node.AppendChild(child)

	// Getting children node list of node.
	childNodes := node.ChildNodes()

	if same := childNodes.Item(0).IsSameNode(child); !same {
		t.Log("Node first child node must be equal to child. (pointer must be the same)")
		t.Logf("child  pointer address                : %p", child)
		t.Logf("Node first child node pointer address : %p", childNodes)
		t.Fail()
	}
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

func TestAppendChild(t *testing.T) {
	// Creating node & child
	node := newNode()
	child := node.CloneNode(false)

	// Appending the child
	child = node.AppendChild(child)

	// Check that node contains the child
	if contain := node.Contains(child); !contain {
		t.Log("Node must contain the appended child.")
		t.Logf("Node number of child : %v", node.ChildNodes().Length())
		t.Fail()
	}

	// Check that child parent is node
	if isParent := child.ParentNode().IsSameNode(node); !isParent {
		t.Log("Node must be the parent of the appended child. (pointer must be the same)")
		t.Logf("Child parent node pointer address : %p", child.ParentNode())
		t.Logf("Node pointer address              : %p", node)
		t.Fail()
	}
}

func TestCloneNode(t *testing.T) {
	node := newNode()
	child := newNode()
	child2 := newNode()

	node.AppendChild(child2)
	child = node.AppendChild(child)

	// Clone the node but not his childs
	clone := node.CloneNode(false)

	// Clone must not be equal (different childs)
	if equal := clone.IsEqualNode(node); equal {
		t.Log("Clone must not be equal to node. (different child)")
		t.Logf("Clone child node : %v", clone.ChildNodes())
		t.Logf("Node child node  : %v", node.ChildNodes())
		t.Fail()
	}

	clone = node.CloneNode(true)

	// Checking that clone is equal to node
	if equal := clone.IsEqualNode(node); !equal {
		t.Log("Clone must be equal to node. (deep clone)")
		t.Logf("Clone child node : %v", clone.ChildNodes().Values())
		t.Logf("Node child node  : %v", node.ChildNodes().Values())
		t.Fail()
	}

	// Checking that clone doesn't point to node
	if same := clone.IsSameNode(node); same {
		t.Log("Clone must not point to the same reference than node.")
		t.Logf("Clone pointers address : %p", clone)
		t.Logf("Node pointers address  : %p", node)
		t.Fail()
	}

	// Checking that clone child parent is
	// same as clone
	if same := clone.ChildNodes().Item(0).ParentNode().IsSameNode(clone); !same {
		t.Log("Clone child parent node must be clone.")
		t.Logf("Clone child parent node pointer address : %p", clone.ChildNodes().Item(0).ParentNode())
		t.Logf("Clone pointer address                   : %p", clone)
		t.Fail()
	}
}

func TestCompareDocumentPosition(t *testing.T) {
	// TODO func TestCompareDocumentPosition(t *testing.T)
}

func TestContains(t *testing.T) {
	node := newNode()
	child1 := newNode()
	child2 := child1.CloneNode(false)

	child1 = node.AppendChild(child1)

	// Checking that child1 is equal child2
	if equal := child1.IsEqualNode(child2); !equal {
		t.Log("Child2 must be a clone of child1.")
		t.Fail()
	}

	// Checking that node doesn't contain child2
	if contain := node.Contains(child2); contain {
		t.Log("Node must not contain child2. (pointer must be different)")
		t.Logf("Node childrens pointer address : %v", node.ChildNodes().Values())
		t.Logf("Child2 pointer address         : %p", child2)
		t.Fail()
	}

	child2 = child1.AppendChild(child2)

	// Checking that child1 contain child2
	if contain := child1.Contains(child2); !contain {
		t.Log("Child1 must contain child2 (direct child). (pointer must be the same)")
		t.Logf("Child1 childrens pointer address : %v", child1.ChildNodes().Values())
		t.Logf("Child2 pointer address           : %p", child2)
		t.Fail()
	}

	// Checking that node contain child2
	// (child of child1)
	if contain := node.Contains(child2); !contain {
		t.Log("Node must contain child2 (direct child of child1). (pointer must be the same)")
		t.Logf("Child1 childrens pointer address : %v", child1.ChildNodes().Values())
		t.Logf("Child2 pointer address           : %p", child2)
		t.Fail()
	}
}

func TestGetRootNode(t *testing.T) {
	node := newNode()
	child1 := newNode()
	child2 := newNode()

	child1 = node.AppendChild(child1)

	// Checking that node root is node
	if same := node.GetRootNode().IsSameNode(node); !same {
		t.Log("Node root node must be node itself. (pointer must be the same)")
		t.Logf("Node root pointer address : %p", node.GetRootNode())
		t.Logf("Node pointer address      : %p", node)
		t.Fail()
	}

	// Checking that child root is node
	if same := child1.GetRootNode().IsSameNode(node); !same {
		t.Log("Node must be the root node of child1. (pointer must be the same)")
		t.Logf("Child1 root pointer address : %p", child1.GetRootNode())
		t.Logf("Node pointer address        : %p", node)
		t.Fail()
	}

	// Checking that child2 root is not node
	if same := child2.GetRootNode().IsSameNode(node); same {
		t.Log("Node must not be the root node of child2. (pointer must not be the same)")
		t.Logf("Child2 root pointer address : %p", child2.GetRootNode())
		t.Logf("Node pointer address        : %p", node)
		t.Fail()
	}
}

func TestHasChildNodes(t *testing.T) {
	node := newNode()
	child := newNode()
	clone := node.CloneNode(false)

	node.AppendChild(child)

	// Checking that node has child.
	if hasChild := node.HasChildNodes(); !hasChild {
		t.Log("Node have child. (node.HasChildNodes must return true)")
		t.Fail()
	}

	// Checking that clone has not child.
	if hasChild := clone.HasChildNodes(); hasChild {
		t.Log("Clone haven't any child. (clone.HasChildNodes must return false)")
		t.Fail()
	}
}

func TestInsertBefore(t *testing.T) {
	node := newNode()

	childNodesCount := 1000

	// Appending 100 of child to node
	for i := 0; i < childNodesCount; i++ {
		child := newNode()
		node.AppendChild(child)
	}

	// Pick a random reference child node
	index := rand.Intn(childNodesCount)
	reference := node.ChildNodes().Item(index)

	// The node to insert
	new := newNode()
	new = node.InsertBefore(new, reference)

	// Check if the node to insert is at the
	// good index
	if insertedIndex := node.ChildNodes().IndexOf(new); insertedIndex != index {
		t.Logf("Inserted node index is %v and should be %v.", insertedIndex, index)
		t.Fail()
	}

	// Checking that node at inserted index
	// is the same that the inserted one
	if same := node.ChildNodes().Item(index).IsSameNode(new); !same {
		t.Log("Child at inserted index must be the same node than the inserted one.")
		t.Logf("Child at inserted index pointer address : %p", node.ChildNodes().Item(index))
		t.Logf("Inserted node pointer address           : %p", new)
		t.Fail()
	}
}

func TestIsEqualNode(t *testing.T) {
	node := newNode()
	clone := node.CloneNode(false)

	// Checking that clone is equal node
	if equal := node.IsEqualNode(clone); !equal {
		t.Log("Clone must be equal to node. (same node type)")
		t.Fail()
	}

	node.AppendChild(newNode())

	// Recloning node (with child)
	clone = node.CloneNode(true)

	// Checking that clone is equal node
	if equal := node.IsEqualNode(clone); !equal {
		t.Log("Clone must be equal to node. (same childrens)")
		t.Fail()
	}

	// Recloning node (without child)
	clone = node.CloneNode(false)

	// Checking that clone is not equal node
	if equal := node.IsEqualNode(clone); equal {
		t.Log("Clone must not be equal to node. (different childrens)")
		t.Fail()
	}
}

func TestNormalize(t *testing.T) {
	// TODO func TestNormalize(t *testing.T)
}

func TestRemoveChild(t *testing.T) {
	node := newNode()
	child := newNode()

	child = node.AppendChild(child)

	// Checking that node contains the child
	if contain := node.Contains(child); !contain {
		t.Log("Node must contain the child.")
		t.FailNow()
	}

	// Removing the child
	child, err := node.RemoveChild(child)

	// Checking that no error occur while
	// removing the child
	if err != nil {
		t.Logf("Error while removing the child : %v", err)
		t.Fail()
	}

	// Checking that child parent is
	// not the same as node
	if same := node.IsSameNode(child.ParentNode()); same {
		t.Log("Child parent must not be node anymore.")
		t.Logf("Child parent node pointer address : %p", child.ParentNode())
		t.Logf("Node pointer address              : %p", node)
		t.Fail()
	}

	// Checking that the node doesn't
	// contain the child anymore
	if contain := node.Contains(child); !contain {
		t.Log("Node must not contain the child.")
		t.Fail()
	}

	/*
	 * Testing error
	 */

	child = node.AppendChild(child)

	_, err = node.RemoveChild(nil)

	if err == nil {
		t.Log("Removing a nil node child pointer must return an error.")
		t.Fail()
	}
}

func TestReplaceChild(t *testing.T) {
	node := newNode()
	child := newNode()

	child = node.AppendChild(child)

	child2 := node.CloneNode(true)

	node.ReplaceChild(child2, child)

	// Checking that child2 is same as
	// node first direct child
	if same := node.ChildNodes().Item(0).IsSameNode(child2); !same {
		t.Fail()
	}

	/*
	 * Testing error
	 */

	err := node.ReplaceChild(child2, nil)

	if err == nil {
		t.Log("Replacing a nil node child pointer must return an error.")
		t.Fail()
	}
}
