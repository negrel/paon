package events

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTarget(t *testing.T) {
	testEventType := NewType("test-type")
	target := NewTarget()

	counter := 0
	listener1 := ListenerFunc(func(event Event) { counter++ })
	listener2 := ListenerFunc(func(event Event) { counter += 2 })

	target.AddEventListener(testEventType, listener1)
	target.AddEventListener(testEventType, listener2)

	target.DispatchEvent(NewEvent(testEventType))
	assert.Equal(t, 3, counter)

	counter = 0
	target.RemoveEventListener(testEventType, listener2)
	target.DispatchEvent(NewEvent(testEventType))
	assert.Equal(t, 1, counter)
}

func BenchmarkTargetDispatchEvent(b *testing.B) {
	for nbEvent := 1; nbEvent <= 1024; nbEvent = nbEvent << 1 {
		for nbListener := 1; nbListener <= 1024; nbListener = nbListener << 1 {
			testName := fmt.Sprintf("%d/With-%d-Listeners", nbEvent, nbListener)
			b.Run(testName, func(b *testing.B) {
				benchEventType := NewType("bench-type")
				target := NewTarget()

				for i := 0; i < nbListener; i++ {
					target.AddEventListener(benchEventType, ListenerFunc(func(event Event) {}))
				}

				events := make([]Event, nbEvent+1)
				for i := 0; i < nbEvent; i++ {
					events[i] = NewEvent(benchEventType)
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
