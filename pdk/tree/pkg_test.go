package tree

import (
	"testing"
	"time"
)

var nodeData = time.Now()

func TestMain(t *testing.T) {
	TestNodes(t)
}

func TestNodes(t *testing.T) {
	t.Run("LeafNode", TestLeafNode)
	t.Run("Node", TestNode)
	t.Run("RootNode", TestRootNode)
}
