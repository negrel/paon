package events

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewType(t *testing.T) {
	typeCount := 1000
	typeSet := sync.Map{}

	var wg sync.WaitGroup
	wg.Add(typeCount)

	// Create some event Type and store them in a set.
	for i := 0; i < typeCount; i++ {
		name := strconv.Itoa(i)
		eventType := NewType(name)

		_, ok := typeSet.Load(eventType)
		require.False(t, ok)
		typeSet.Store(eventType, nil)
	}
}

func TestTypeString(t *testing.T) {
	typeCount := 1000

	for i := 0; i < typeCount; i++ {
		// Create Types concurrently

		name := strconv.Itoa(i)
		eventType := NewType(name)

		// Check the type name concurrently
		require.Equal(t, eventType.name(), name)
	}
}
