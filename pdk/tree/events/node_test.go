package events

import (
	"testing"

	"github.com/negrel/paon/pdk/tree"
	"github.com/stretchr/testify/require"
)

func newRootNode(data interface{}) tree.Node {
	return tree.NewRoot(data)
}

// wal walks recursively the given node subtree and call fn
// on each non-nil node.
func walk(node *BaseNode, fn func(*BaseNode)) {
	if node == nil {
		return
	}

	fn(node)
	for child := node.FirstChild(); child != nil; child = child.Next() {
		walk(child.(*BaseNode), fn)
	}
}

func TestNodeLifeCycleStage(t *testing.T) {
	var root, greatParent, parent, child1, child2 *BaseNode
	root = NewBaseNode(NodeConstructor(newRootNode))
	greatParent = NewBaseNode(NodeConstructor(tree.NewNode))
	parent = NewBaseNode(NodeConstructor(tree.NewNode))
	child1 = NewBaseNode(NodeConstructor(tree.NewLeafNode))
	child2 = NewBaseNode(NodeConstructor(tree.NewLeafNode))

	// Let's add some counter for beforeXXX lifecycle stages
	beforeMountCounters := make(map[*BaseNode]int)
	beforeUnmountCounters := make(map[*BaseNode]int)
	lifecycleListener := func(lce LifeCycleEvent) {
		if lce.Stage == LCSBeforeMount {
			beforeMountCounters[lce.Node.(*BaseNode)]++
		} else if lce.Stage == LCSBeforeUnmount {
			beforeUnmountCounters[lce.Node.(*BaseNode)]++
		}
	}

	root.AddEventListener(LifeCycleEventListener(lifecycleListener))
	greatParent.AddEventListener(LifeCycleEventListener(lifecycleListener))
	parent.AddEventListener(LifeCycleEventListener(lifecycleListener))
	child1.AddEventListener(LifeCycleEventListener(lifecycleListener))
	child2.AddEventListener(LifeCycleEventListener(lifecycleListener))

	// Check initial states
	require.Equal(t, LCSMounted, root.LifeCycleStage())
	require.Equal(t, LCSInitial, greatParent.LifeCycleStage())
	require.Equal(t, LCSInitial, parent.LifeCycleStage())
	require.Equal(t, LCSInitial, child1.LifeCycleStage())
	require.Equal(t, LCSInitial, child2.LifeCycleStage())

	// Adding a child to a non mounted node should'n
	// change the state of both nodes
	err := greatParent.AppendChild(parent)
	require.Nil(t, err)
	walk(greatParent, func(n *BaseNode) {
		require.Equal(t, LCSInitial, n.LifeCycleStage())
		require.Equal(t, 0, beforeMountCounters[n])
	})

	// Same for child1 and child2
	err = parent.AppendChild(child2)
	require.Nil(t, err)
	walk(parent, func(n *BaseNode) {
		require.Equal(t, LCSInitial, n.LifeCycleStage())
		require.Equal(t, 0, beforeMountCounters[n])
	})
	err = parent.InsertBefore(child2, child1)
	require.Nil(t, err)
	walk(parent, func(n *BaseNode) {
		require.Equal(t, LCSInitial, n.LifeCycleStage())
		require.Equal(t, 0, beforeMountCounters[n])
	})

	// Now let's mount the entire subtree
	err = root.AppendChild(greatParent)
	require.Nil(t, err)
	walk(greatParent, func(n *BaseNode) {
		require.Equal(t, LCSMounted, n.LifeCycleStage())
		require.Equal(t, 1, beforeMountCounters[n])
	})

	// Let's remove the entire subtree
	err = root.RemoveChild(greatParent)
	require.Nil(t, err)
	walk(greatParent, func(n *BaseNode) {
		require.Equal(t, LCSUnmounted, n.LifeCycleStage())
		require.Equal(t, 1, beforeUnmountCounters[n])
	})
}

func TestNodeWrapOption(t *testing.T) {
	node := NewBaseNode(Wrap(t))
	require.Equal(t, t, node.Unwrap())

	node = NewBaseNode()
	require.Equal(t, node, node.Unwrap())
}
