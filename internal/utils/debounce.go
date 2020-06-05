package utils

import "time"

// Debounce merge multiple call of the given function until
// the wait time has elapsed.
func Debounce(fn func(), wait time.Duration) func() {
	timer := time.NewTimer(wait)

	return func() {
		select {
		case <-timer.C:
			fn()

		default:
			timer.Reset(wait)
		}
	}
}
