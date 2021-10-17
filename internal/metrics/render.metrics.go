//go:build metrics
// +build metrics

package metrics

import "time"

var startRenderTime time.Time

// StartRenderTimer starts the render timer.
func StartRenderTimer() {
	startRenderTime = time.Now()
}

// StopRenderTimer stops the render timer.
func StopRenderTimer() {
	store.render.push(time.Now().Sub(startRenderTime))
}
