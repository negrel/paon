package events

import (
	"testing"

	"github.com/negrel/paon/tree"
	"github.com/stretchr/testify/require"
)

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

func TestNode(t *testing.T) {
	t.Run("NewBaseNode", func(t *testing.T) {
		t.Run("LeafNode/StageIsInitialStage", func(t *testing.T) {
			node := NewBaseNode(NodeConstructor(tree.NewLeafNode))
			require.Equal(t, LCSInitial.String(), node.LifeCycleStage().String())
		})

		t.Run("Node/StageIsInitialStage", func(t *testing.T) {
			node := NewBaseNode(NodeConstructor(tree.NewNode))
			require.Equal(t, LCSInitial.String(), node.LifeCycleStage().String())
		})

		t.Run("RootNode/StageIsMounted", func(t *testing.T) {
			node := NewBaseNode(RootNode)
			require.Equal(t, LCSMounted.String(), node.LifeCycleStage().String())
		})
	})

	t.Run("LifeCycleStage", func(t *testing.T) {
		var root, greatParent, parent, child1, child2 *BaseNode
		root = NewBaseNode(RootNode, NodeConstructor(tree.NewNode))
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

		// Adding a child to a non mounted node should not
		// change the state of both nodes
		err := greatParent.AppendChild(parent)
		require.NoError(t, err)
		walk(greatParent, func(n *BaseNode) {
			require.Equal(t, LCSInitial, n.LifeCycleStage())
			require.Equal(t, 0, beforeMountCounters[n])
		})

		// Same for child1 and child2
		err = parent.AppendChild(child2)
		require.NoError(t, err)
		walk(parent, func(n *BaseNode) {
			require.Equal(t, LCSInitial, n.LifeCycleStage())
			require.Equal(t, 0, beforeMountCounters[n])
		})
		err = parent.InsertBefore(child1, child2)
		require.NoError(t, err)
		walk(parent, func(n *BaseNode) {
			require.Equal(t, LCSInitial, n.LifeCycleStage())
			require.Equal(t, 0, beforeMountCounters[n])
		})

		// Now let's mount the entire subtree
		err = root.AppendChild(greatParent)
		require.NoError(t, err)
		walk(greatParent, func(n *BaseNode) {
			require.Equal(t, LCSMounted, n.LifeCycleStage())
			require.Equal(t, 1, beforeMountCounters[n])
		})

		// Let's remove the entire subtree
		err = root.RemoveChild(greatParent)
		require.NoError(t, err)
		walk(greatParent, func(n *BaseNode) {
			require.Equal(t, LCSUnmounted, n.LifeCycleStage())
			require.Equal(t, 1, beforeUnmountCounters[n])
		})
	})
}

func TestNodeWrapOption(t *testing.T) {
	node := NewBaseNode(Wrap(t))
	require.Equal(t, t, node.Unwrap())

	node = NewBaseNode()
	require.Equal(t, node, node.Unwrap())
}
