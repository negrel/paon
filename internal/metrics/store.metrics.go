// +build metrics

package metrics

import (
	"fmt"
	"io"
)

var store = struct {
	layout *durationSlice
	draw   *durationSlice
}{
	layout: newDurationSlice(),
	draw:   newDurationSlice(),
}

// Report write the metrics report to the given io.Writer.
func Report(w io.Writer) {
	layoutMean := store.layout.mean()
	drawMean := store.draw.mean()

	fmt.Fprintf(w, "Layout (min/mean/max) %v/%v/%v\n", store.layout.min, layoutMean, store.layout.max)
	fmt.Fprintf(w, "Draw (min/mean/max) %v/%v/%v\n", store.draw.min, drawMean, store.draw.max)
	fmt.Fprintf(w, "Layout+Draw (min/mean/max) %v/%v/%v\n",
		store.layout.min+store.draw.min,
		layoutMean+drawMean,
		store.layout.max+store.draw.max,
	)
}
