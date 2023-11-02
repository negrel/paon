package draw

import "github.com/negrel/paon/colors"

// CellStyle define the style of a single terminal Cell.
type CellStyle struct {
	Foreground colors.Color
	Background colors.Color

	Bold, Blink, Reverse, Underline,
	Dim, Italic, StrikeThrough bool
}

// Cell define a terminal Surface cell.
type Cell struct {
	Style   CellStyle
	Content rune
}
