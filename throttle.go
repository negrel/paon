package paon

import "time"

// throttle executes the given function at most once every d duration.
// Function will be executed in its own goroutine and after d duration has passed.
func throttle(d time.Duration, fn func()) func() {
	scheduled := false

	return func() {
		if !scheduled {
			time.AfterFunc(d, func() {
				scheduled = false
				fn()
			})
			scheduled = true
		}
	}
}
