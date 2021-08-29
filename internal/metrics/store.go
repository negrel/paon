// +build !metrics

package metrics

import (
	"io"
)

var store = struct {
	layout *durationSlice
	draw   *durationSlice
}{
	layout: nil,
	draw:   nil,
}

// Report write the metrics report to the given io.Writer.
func Report(w io.Writer) {

}
