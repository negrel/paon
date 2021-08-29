// +build !metrics

package metrics

import "time"

var startDrawTime time.Time

// StartDrawTimer starts the draw timer.
func StartDrawTimer() {
	startDrawTime = time.Now()
}

// StopDrawTimer stops the draw timer.
func StopDrawTimer() {
	store.draw.push(startDrawTime.Sub(time.Now()))
}
