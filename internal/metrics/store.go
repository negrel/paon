//go:build !metrics
// +build !metrics

package metrics

import (
	"io"
)

var store = struct {
	render *durationSlice
}{}

// Report write the metrics report to the given io.Writer.
func Report(w io.Writer) {

}
