package tree

import (
	"sync/atomic"
)

type NodeID string

var id int32 = 0

func makeNodeID() NodeID {
	nID := atomic.AddInt32(&id, 1)
	return NodeID(nID)
}
