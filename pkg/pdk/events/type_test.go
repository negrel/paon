package events

import (
	"github.com/negrel/paon/internal/idmap"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func Test_MakeType(t *testing.T) {
	typeCount := 1000
	types := idmap.New(typeCount)

	var wg sync.WaitGroup
	wg.Add(typeCount)

	// Create typeCount Type and store them in types using the
	// channel to avoid concurrent write and read
	for i := 0; i < typeCount; i++ {
		go func(i int) {
			name := strconv.Itoa(i)
			types.Set(MakeType(name))
		}(i)
	}

	// Check that received Type are unique and then write them to the map.
	go func() {
		for _type := range ch {
			assert.NotContains(t, types, _type, "duplicate event type")
			types[_type] = struct{}{}

			wg.Done()
		}
	}()
	wg.Wait()
}

func Test_Type_String(t *testing.T) {
	typeCount := 1000

	var wg sync.WaitGroup
	wg.Add(typeCount)

	for i := 0; i < typeCount; i++ {
		// Create Types concurrently
		go func(i int) {
			name := strconv.Itoa(i)
			_type := MakeType(name)

			// Check the type name concurrently
			go func(_type Type, name string) {
				defer wg.Done()
				assert.Equal(t, _type.name(), name)
			}(_type, name)
		}(i)
	}
	wg.Wait()
}
