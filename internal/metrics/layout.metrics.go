// +build metrics

package metrics

import "time"

var startLayoutTime time.Time

// StartLayoutTimer starts the layout timer.
func StartLayoutTimer() {
	startLayoutTime = time.Now()
}

// StopLayoutTimer stops the layout timer.
func StopLayoutTimer() {
	store.layout.push(time.Now().Sub(startLayoutTime))
}
