package render

import (
	"github.com/negrel/paon/pkg/pdk/styles/value"
)

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
