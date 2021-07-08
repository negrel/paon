package id

import (
	"strconv"
	"sync"
	"testing"
)

func Test_ConcurrentReadWrite(t *testing.T) {
	idCount := 1000
	ids := NewMap()

	var wg sync.WaitGroup
	wg.Add(idCount)

	for i := 0; i < idCount; i++ {
		go func(i int32) {
			if ids.Get(ID(i)) != "" {
				t.Fail()
			}
			ids.Set(ID(i), strconv.Itoa(int(i)))

			wg.Done()
		}(int32(i))
	}
	wg.Wait()
}
