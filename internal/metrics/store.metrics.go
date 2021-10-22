//go:build metrics
// +build metrics

package metrics

import (
	"fmt"
	"io"
)

var store = struct {
	render *durationSlice
}{
	render: newDurationSlice(),
}

// Report write the metrics report to the given io.Writer.
func Report(w io.Writer) {
	renderMean := store.render.mean()

	fmt.Fprintf(w, "Render (total: %v) (min/mean/max) %v/%v/%v\n", len(store.render.data), store.render.min, renderMean, store.render.max)
}
