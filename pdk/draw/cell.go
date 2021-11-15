package draw

import "github.com/negrel/paon/styles/property"

// CellStyle define the style of a single terminal Cell.
type CellStyle struct {
	Foreground property.Color
	Background property.Color

	Bold, Blink, Reverse, Underline,
	Dim, Italic, StrikeThrough property.Bool
}

// Cell define a terminal Surface cell.
type Cell struct {
	Style   CellStyle
	Content rune
}
