package draw

import "github.com/negrel/paon/styles/value"

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

func (c Cell) Merge(other Cell) Cell {
	if c.Content == '\000' {
		return other
	} else if other.Content == '\000' {
		return c
	}

	return Cell{}
}

// ZeroCell returns the zero value of a Cell.
func ZeroCell() Cell {
	return Cell{
		Content: '\000',
	}
}
