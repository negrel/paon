package render

import (
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

// CellStyle define the style of a single terminal Cell.
type CellStyle struct {
	Foreground value.Color
	Background value.Color

	Bold, Blink, Reverse, Underline,
	Dim, Italic, StrikeThrough bool
}

// Cell define a terminal Surface cell.
type Cell struct {
	Style   CellStyle
	Content rune
}
