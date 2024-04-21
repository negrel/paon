package events

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTarget(t *testing.T) {
	testEventType := NewType("TestType")
	target := NewTarget()

	counter := 0
	listener1 := NewListenerFunc(func(event Event) { counter++ })
	listener2 := NewListenerFunc(func(event Event) { counter += 2 })

	target.AddEventListener(testEventType, listener1)
	target.AddEventListener(testEventType, listener2)

	target.DispatchEvent(NewEvent(testEventType, nil))
	assert.Equal(t, 3, counter)

	counter = 0
	target.RemoveEventListener(testEventType, listener2)
	target.DispatchEvent(NewEvent(testEventType, nil))
	assert.Equal(t, 1, counter)
}

func TestTargetModifyAndForwardEvent(t *testing.T) {
	testEventType := NewType("TestType")
	target := NewTarget()
	// target2 := NewTarget()
	otherTarget := NewTarget()

	listenerCall := 0
	listener1 := NewListenerFunc(func(event Event) {
		require.Equal(t, 0, listenerCall)
		require.Equal(t, target, event.Target())
		event.WithTarget(otherTarget) // Try to overwrite.
		listenerCall++
	})
	listener2 := NewListenerFunc(func(event Event) {
		require.Equal(t, target, event.Target())
		listenerCall++
	})

	target.AddEventListener(testEventType, listener1)
	target.AddEventListener(testEventType, listener2)

	ev := NewEvent(testEventType, nil).WithTarget(target)
	target.DispatchEvent(ev)
}

func BenchmarkTargetDispatchEvent(b *testing.B) {
	for nbEvent := 1; nbEvent <= 1024; nbEvent = nbEvent << 1 {
		for nbListener := 1; nbListener <= 1024; nbListener = nbListener << 1 {
			testName := fmt.Sprintf("%d/With-%d-Listeners", nbEvent, nbListener)
			b.Run(testName, func(b *testing.B) {
				benchEventType := NewType("BenchType")
				target := NewTarget()

				for i := 0; i < nbListener; i++ {
					target.AddEventListener(benchEventType, NewListenerFunc(func(event Event) {}))
				}

				events := make([]Event, nbEvent+1)
				for i := 0; i < nbEvent; i++ {
					events[i] = NewEvent(benchEventType, i)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for j := 0; j < nbEvent; j++ {
						target.DispatchEvent(events[j])
					}
				}
			})
		}
	}
}
