package id

import "sync/atomic"

type ID int32

var idCounter int32 = 0

func Make() ID {
	id := atomic.AddInt32(&idCounter, 1)

	return ID(id)
}
