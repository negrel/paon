package property

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewID(t *testing.T) {
	idCount := 1000
	ids := make(map[ID]struct{}, idCount)
	ch := make(chan ID)

	var wg sync.WaitGroup
	wg.Add(idCount)

	// Create idCount ID and store them in ids using the
	// channel to avoid concurrent write and read
	for i := 0; i < idCount; i++ {
		go func(i int) {
			name := strconv.Itoa(i)
			ch <- NewID(name)
		}(i)
	}

	// Check that received ID are unique and then write them to the map.
	go func() {
		for _id := range ch {
			assert.NotContains(t, ids, _id, "duplicate event id")
			ids[_id] = struct{}{}

			wg.Done()
		}
	}()
	wg.Wait()
}

func Test_ID_String(t *testing.T) {
	idCount := 1000

	var wg sync.WaitGroup
	wg.Add(idCount)

	for i := 0; i < idCount; i++ {
		// Create IDs concurrently
		go func(i int) {
			name := strconv.Itoa(i)
			_id := NewID(name)

			// Check the id name concurrently
			go func(_id ID, name string) {
				defer wg.Done()
				assert.Equal(t, _id.name(), name)
			}(_id, name)
		}(i)
	}
	wg.Wait()
}
