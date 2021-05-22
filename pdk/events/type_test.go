package events

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewType(t *testing.T) {
	typeCount := 1000
	typeSet := sync.Map{}

	var wg sync.WaitGroup
	wg.Add(typeCount)

	// Create some event Type and store them in a set.
	for i := 0; i < typeCount; i++ {
		go func(i int) {
			defer wg.Done()

			name := strconv.Itoa(i)
			eventType := NewType(name)

			_, ok := typeSet.Load(eventType)
			assert.False(t, ok)
			typeSet.Store(eventType, nil)
		}(i)
	}

	wg.Wait()
}

func Test_Type_String(t *testing.T) {
	typeCount := 1000

	var wg sync.WaitGroup
	wg.Add(typeCount)

	for i := 0; i < typeCount; i++ {
		// Create Types concurrently
		go func(i int) {
			defer wg.Done()

			name := strconv.Itoa(i)
			_type := NewType(name)

			// Check the type name concurrently
			assert.Equal(t, _type.name(), name)
		}(i)
	}
	wg.Wait()
}
