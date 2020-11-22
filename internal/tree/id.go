package tree

import (
	"sync/atomic"
)

type NodeID int32

var id int32 = 0

func nodeID() NodeID {
	nID := atomic.AddInt32(&id, 1)
	return NodeID(nID)
}
